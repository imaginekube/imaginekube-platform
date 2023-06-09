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

package notification

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"imaginekube.com/api/notification/v2beta2"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/client/clientset/versioned/fake"
	ksinformers "imaginekube.com/imaginekube/pkg/client/informers/externalversions"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha3"
	"imaginekube.com/imaginekube/pkg/server/errors"
)

const (
	Prefix    = "foo"
	LengthMin = 3
	LengthMax = 10
)

func TestListObjects(t *testing.T) {
	tests := []struct {
		description string
		key         string
	}{
		{
			"test name filter",
			v2beta2.ResourcesPluralConfig,
		},
		{
			"test name filter",
			v2beta2.ResourcesPluralReceiver,
		},
		{
			"test name filter",
			v2beta2.ResourcesPluralRouter,
		},
		{
			"test name filter",
			v2beta2.ResourcesPluralSilence,
		},
	}

	q := &query.Query{
		Pagination: &query.Pagination{
			Limit:  10,
			Offset: 0,
		},
		SortBy:    query.FieldName,
		Ascending: true,
		Filters:   map[query.Field]query.Value{query.FieldName: query.Value(Prefix)},
	}

	for _, test := range tests {

		getter, objects, err := prepare(test.key)
		if err != nil {
			t.Fatal(err)
		}

		got, err := getter.List("", q)
		if err != nil {
			t.Fatal(err)
		}

		expected := &api.ListResult{
			Items:      objects,
			TotalItems: len(objects),
		}

		if diff := cmp.Diff(got, expected); diff != "" {
			t.Errorf("[%s] %T differ (-got, +want): %s", test.description, expected, diff)
		}
	}
}

func prepare(key string) (v1alpha3.Interface, []interface{}, error) {
	client := fake.NewSimpleClientset()
	informer := ksinformers.NewSharedInformerFactory(client, 0)

	var obj runtime.Object
	var indexer cache.Indexer
	var getter func(informer ksinformers.SharedInformerFactory) v1alpha3.Interface
	switch key {
	case v2beta2.ResourcesPluralConfig:
		indexer = informer.Notification().V2beta2().Configs().Informer().GetIndexer()
		getter = NewNotificationConfigGetter
		obj = &v2beta2.Config{}
	case v2beta2.ResourcesPluralReceiver:
		indexer = informer.Notification().V2beta2().Receivers().Informer().GetIndexer()
		getter = NewNotificationReceiverGetter
		obj = &v2beta2.Receiver{}
	case v2beta2.ResourcesPluralRouter:
		indexer = informer.Notification().V2beta2().Routers().Informer().GetIndexer()
		getter = NewNotificationRouterGetter
		obj = &v2beta2.Router{}
	case v2beta2.ResourcesPluralSilence:
		indexer = informer.Notification().V2beta2().Silences().Informer().GetIndexer()
		getter = NewNotificationSilenceGetter
		obj = &v2beta2.Silence{}
	default:
		return nil, nil, errors.New("unowned type %s", key)
	}

	num := rand.Intn(LengthMax)
	if num < LengthMin {
		num = LengthMin
	}

	var suffix []string
	for i := 0; i < num; i++ {
		s := uuid.New().String()
		suffix = append(suffix, s)
	}
	sort.Strings(suffix)

	var objects []interface{}
	for i := 0; i < num; i++ {
		val := obj.DeepCopyObject()
		accessor, err := meta.Accessor(val)
		if err != nil {
			return nil, nil, err
		}

		accessor.SetName(Prefix + "-" + suffix[i])
		err = indexer.Add(accessor)
		if err != nil {
			return nil, nil, err
		}
		objects = append(objects, val)
	}

	return getter(informer), objects, nil
}
