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

package volumesnapshot

import (
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	"github.com/kubernetes-csi/external-snapshotter/client/v4/informers/externalversions"
	"k8s.io/apimachinery/pkg/runtime"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha3"
)

const (
	statusCreating = "creating"
	statusReady    = "ready"
	statusDeleting = "deleting"

	volumeSnapshotClassName   = "volumeSnapshotClassName"
	persistentVolumeClaimName = "persistentVolumeClaimName"
)

type volumeSnapshotGetter struct {
	informers externalversions.SharedInformerFactory
}

func New(informer externalversions.SharedInformerFactory) v1alpha3.Interface {
	return &volumeSnapshotGetter{informers: informer}
}

func (v *volumeSnapshotGetter) Get(namespace, name string) (runtime.Object, error) {
	return v.informers.Snapshot().V1().VolumeSnapshots().Lister().VolumeSnapshots(namespace).Get(name)
}

func (v *volumeSnapshotGetter) List(namespace string, query *query.Query) (*api.ListResult, error) {
	all, err := v.informers.Snapshot().V1().VolumeSnapshots().Lister().VolumeSnapshots(namespace).List(query.Selector())
	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, snapshot := range all {
		result = append(result, snapshot)
	}

	return v1alpha3.DefaultList(result, query, v.compare, v.filter), nil
}

func (v *volumeSnapshotGetter) compare(left, right runtime.Object, field query.Field) bool {
	leftSnapshot, ok := left.(*v1.VolumeSnapshot)
	if !ok {
		return false
	}
	rightSnapshot, ok := right.(*v1.VolumeSnapshot)
	if !ok {
		return false
	}
	return v1alpha3.DefaultObjectMetaCompare(leftSnapshot.ObjectMeta, rightSnapshot.ObjectMeta, field)
}

func (v *volumeSnapshotGetter) filter(object runtime.Object, filter query.Filter) bool {
	snapshot, ok := object.(*v1.VolumeSnapshot)
	if !ok {
		return false
	}

	switch filter.Field {
	case query.FieldStatus:
		return snapshotStatus(snapshot) == string(filter.Value)
	case volumeSnapshotClassName:
		name := snapshot.Spec.VolumeSnapshotClassName
		return name != nil && *name == string(filter.Value)
	case persistentVolumeClaimName:
		name := snapshot.Spec.Source.PersistentVolumeClaimName
		return name != nil && *name == string(filter.Value)
	default:
		return v1alpha3.DefaultObjectMetaFilter(snapshot.ObjectMeta, filter)
	}
}

func snapshotStatus(item *v1.VolumeSnapshot) string {
	status := statusCreating
	if item != nil && item.Status != nil && item.Status.ReadyToUse != nil && *item.Status.ReadyToUse {
		status = statusReady
	}
	if item != nil && item.DeletionTimestamp != nil {
		status = statusDeleting
	}
	return status
}
