/*
Copyright 2019 The ImagineKube Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package workspace

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	tenantv1alpha1 "imaginekube.com/api/tenant/v1alpha1"

	"imaginekube.com/imaginekube/pkg/constants"
	controllerutils "imaginekube.com/imaginekube/pkg/controller/utils/controller"
	"imaginekube.com/imaginekube/pkg/utils/k8sutil"
	"imaginekube.com/imaginekube/pkg/utils/sliceutil"
)

const (
	controllerName = "workspace-controller"
)

// Reconciler reconciles a Workspace object
type Reconciler struct {
	client.Client
	Logger                  logr.Logger
	Recorder                record.EventRecorder
	MaxConcurrentReconciles int
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.Client == nil {
		r.Client = mgr.GetClient()
	}
	if r.Logger.GetSink() == nil {
		r.Logger = ctrl.Log.WithName("controllers").WithName(controllerName)
	}
	if r.Recorder == nil {
		r.Recorder = mgr.GetEventRecorderFor(controllerName)
	}
	if r.MaxConcurrentReconciles <= 0 {
		r.MaxConcurrentReconciles = 1
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: r.MaxConcurrentReconciles,
		}).
		For(&tenantv1alpha1.Workspace{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=tenant.imaginekube.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tenant.imaginekube.com,resources=workspaces/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=iam.imaginekube.com,resources=users,verbs=get;list;watch
// +kubebuilder:rbac:groups=iam.imaginekube.com,resources=rolebases,verbs=get;list;watch
// +kubebuilder:rbac:groups=iam.imaginekube.com,resources=workspaceroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=iam.imaginekube.com,resources=workspacerolebindings,verbs=get;list;watch;create;update;patch;delete
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Logger.WithValues("workspace", req.NamespacedName)
	rootCtx := context.Background()
	workspace := &tenantv1alpha1.Workspace{}
	if err := r.Get(rootCtx, req.NamespacedName, workspace); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// name of your custom finalizer
	finalizer := "finalizers.tenant.imaginekube.com"

	if workspace.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object.
		if !sliceutil.HasString(workspace.ObjectMeta.Finalizers, finalizer) {
			workspace.ObjectMeta.Finalizers = append(workspace.ObjectMeta.Finalizers, finalizer)
			if err := r.Update(rootCtx, workspace); err != nil {
				return ctrl.Result{}, err
			}
			workspaceOperation.WithLabelValues("create", workspace.Name).Inc()
		}
	} else {
		// The object is being deleted
		if sliceutil.HasString(workspace.ObjectMeta.Finalizers, finalizer) {
			// remove our finalizer from the list and update it.
			workspace.ObjectMeta.Finalizers = sliceutil.RemoveString(workspace.ObjectMeta.Finalizers, func(item string) bool {
				return item == finalizer
			})
			logger.V(4).Info("update workspace")
			if err := r.Update(rootCtx, workspace); err != nil {
				logger.Error(err, "update workspace failed")
				return ctrl.Result{}, err
			}
			workspaceOperation.WithLabelValues("delete", workspace.Name).Inc()
		}
		// Our finalizer has finished, so the reconciler can do nothing.
		return ctrl.Result{}, nil
	}

	var namespaces corev1.NamespaceList
	if err := r.List(rootCtx, &namespaces, client.MatchingLabels{tenantv1alpha1.WorkspaceLabel: req.Name}); err != nil {
		logger.Error(err, "list namespaces failed")
		return ctrl.Result{}, err
	} else {
		for _, namespace := range namespaces.Items {
			// managed by kubefed-controller-manager
			kubefedManaged := namespace.Labels[constants.KubefedManagedLabel] == "true"
			if kubefedManaged {
				continue
			}
			// managed by workspace
			if err := r.bindWorkspace(rootCtx, logger, &namespace, workspace); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	r.Recorder.Event(workspace, corev1.EventTypeNormal, controllerutils.SuccessSynced, controllerutils.MessageResourceSynced)
	return ctrl.Result{}, nil
}

func (r *Reconciler) bindWorkspace(ctx context.Context, logger logr.Logger, namespace *corev1.Namespace, workspace *tenantv1alpha1.Workspace) error {
	// owner reference not match workspace label
	if !metav1.IsControlledBy(namespace, workspace) {
		namespace := namespace.DeepCopy()
		namespace.OwnerReferences = k8sutil.RemoveWorkspaceOwnerReference(namespace.OwnerReferences)
		if err := controllerutil.SetControllerReference(workspace, namespace, scheme.Scheme); err != nil {
			logger.Error(err, "set controller reference failed")
			return err
		}
		logger.V(4).Info("update namespace owner reference", "workspace", workspace.Name)
		if err := r.Update(ctx, namespace); err != nil {
			logger.Error(err, "update namespace failed")
			return err
		}
	}
	return nil
}
