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

package s2ibuilder

import (
	"k8s.io/apimachinery/pkg/labels"

	"imaginekube.com/api/devops/v1alpha1"

	"imaginekube.com/imaginekube/pkg/client/informers/externalversions"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha2"

	"sort"

	"imaginekube.com/imaginekube/pkg/server/params"
)

type s2iBuilderSearcher struct {
	informers externalversions.SharedInformerFactory
}

func NewS2iBuilderSearcher(informers externalversions.SharedInformerFactory) v1alpha2.Interface {
	return &s2iBuilderSearcher{informers: informers}
}

func (s *s2iBuilderSearcher) Get(namespace, name string) (interface{}, error) {
	return s.informers.Devops().V1alpha1().S2iBuilders().Lister().S2iBuilders(namespace).Get(name)
}

func (*s2iBuilderSearcher) match(match map[string]string, item *v1alpha1.S2iBuilder) bool {
	for k, v := range match {
		if !v1alpha2.ObjectMetaExactlyMath(k, v, item.ObjectMeta) {
			return false
		}
	}
	return true
}

func (*s2iBuilderSearcher) fuzzy(fuzzy map[string]string, item *v1alpha1.S2iBuilder) bool {
	for k, v := range fuzzy {
		if !v1alpha2.ObjectMetaFuzzyMath(k, v, item.ObjectMeta) {
			return false
		}
	}
	return true
}

func (s *s2iBuilderSearcher) Search(namespace string, conditions *params.Conditions, orderBy string, reverse bool) ([]interface{}, error) {
	s2iBuilders, err := s.informers.Devops().V1alpha1().S2iBuilders().Lister().S2iBuilders(namespace).List(labels.Everything())

	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.S2iBuilder, 0)

	if len(conditions.Match) == 0 && len(conditions.Fuzzy) == 0 {
		result = s2iBuilders
	} else {
		for _, item := range s2iBuilders {
			if s.match(conditions.Match, item) && s.fuzzy(conditions.Fuzzy, item) {
				result = append(result, item)
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if reverse {
			i, j = j, i
		}
		return v1alpha2.ObjectMetaCompare(result[i].ObjectMeta, result[j].ObjectMeta, orderBy)
	})

	r := make([]interface{}, 0)
	for _, i := range result {
		r = append(r, i)
	}
	return r, nil
}
