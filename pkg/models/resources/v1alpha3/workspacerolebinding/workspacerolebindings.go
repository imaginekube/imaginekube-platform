/*
Copyright 2023 The ImagineKube Authors.

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

package workspacerolebinding

import (
	"k8s.io/apimachinery/pkg/runtime"

	iamv1alpha2 "imaginekube.com/api/iam/v1alpha2"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	informers "imaginekube.com/imaginekube/pkg/client/informers/externalversions"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha3"
)

const RoleName = "rolename"

type workspacerolebindingsGetter struct {
	sharedInformers informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) v1alpha3.Interface {
	return &workspacerolebindingsGetter{sharedInformers: sharedInformers}
}

func (d *workspacerolebindingsGetter) Get(_, name string) (runtime.Object, error) {
	return d.sharedInformers.Iam().V1alpha2().GlobalRoleBindings().Lister().Get(name)
}

func (d *workspacerolebindingsGetter) List(_ string, query *query.Query) (*api.ListResult, error) {

	globalRoleBindings, err := d.sharedInformers.Iam().V1alpha2().WorkspaceRoleBindings().Lister().List(query.Selector())

	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, globalRoleBinding := range globalRoleBindings {
		result = append(result, globalRoleBinding)
	}

	return v1alpha3.DefaultList(result, query, d.compare, d.filter), nil
}

func (d *workspacerolebindingsGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {

	leftRoleBinding, ok := left.(*iamv1alpha2.WorkspaceRoleBinding)
	if !ok {
		return false
	}

	rightRoleBinding, ok := right.(*iamv1alpha2.WorkspaceRoleBinding)
	if !ok {
		return false
	}

	return v1alpha3.DefaultObjectMetaCompare(leftRoleBinding.ObjectMeta, rightRoleBinding.ObjectMeta, field)
}

func (d *workspacerolebindingsGetter) filter(object runtime.Object, filter query.Filter) bool {
	role, ok := object.(*iamv1alpha2.WorkspaceRoleBinding)

	if !ok {
		return false
	}
	switch filter.Field {
	case RoleName:
		return role.RoleRef.Name == string(filter.Value)
	default:
		return v1alpha3.DefaultObjectMetaFilter(role.ObjectMeta, filter)
	}
}
