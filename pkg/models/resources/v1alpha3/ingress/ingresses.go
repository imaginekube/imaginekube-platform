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

package ingress

import (
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha3"
)

type ingressGetter struct {
	sharedInformers informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) v1alpha3.Interface {
	return &ingressGetter{sharedInformers: sharedInformers}
}

func (g *ingressGetter) Get(namespace, name string) (runtime.Object, error) {
	return g.sharedInformers.Networking().V1().Ingresses().Lister().Ingresses(namespace).Get(name)
}

func (g *ingressGetter) List(namespace string, query *query.Query) (*api.ListResult, error) {
	// first retrieves all deployments within given namespace
	ingresses, err := g.sharedInformers.Networking().V1().Ingresses().Lister().Ingresses(namespace).List(query.Selector())
	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, ingress := range ingresses {
		result = append(result, ingress)
	}

	return v1alpha3.DefaultList(result, query, g.compare, g.filter), nil
}

func (g *ingressGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {

	leftIngress, ok := left.(*v1.Ingress)
	if !ok {
		return false
	}

	rightIngress, ok := right.(*v1.Ingress)
	if !ok {
		return false
	}

	switch field {
	case query.FieldUpdateTime:
		fallthrough
	default:
		return v1alpha3.DefaultObjectMetaCompare(leftIngress.ObjectMeta, rightIngress.ObjectMeta, field)
	}
}

func (g *ingressGetter) filter(object runtime.Object, filter query.Filter) bool {
	deployment, ok := object.(*v1.Ingress)
	if !ok {
		return false
	}
	return v1alpha3.DefaultObjectMetaFilter(deployment.ObjectMeta, filter)
}
