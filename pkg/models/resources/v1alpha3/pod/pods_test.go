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

package pod

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/models/resources/v1alpha3"
)

func TestListPods(t *testing.T) {
	tests := []struct {
		description string
		namespace   string
		query       *query.Query
		expected    *api.ListResult
		expectedErr error
	}{
		{
			"test name filter",
			"default",
			&query.Query{
				Pagination: &query.Pagination{
					Limit:  10,
					Offset: 0,
				},
				SortBy:    query.FieldName,
				Ascending: false,
				Filters:   map[query.Field]query.Value{query.FieldNamespace: query.Value("default")},
			},
			&api.ListResult{
				Items:      []interface{}{foo5, foo4, foo3, foo2, foo1},
				TotalItems: len(pods),
			},
			nil,
		},
		{
			"test pvcName filter",
			"default",
			&query.Query{
				Pagination: &query.Pagination{
					Limit:  10,
					Offset: 0,
				},
				SortBy:    query.FieldName,
				Ascending: false,
				Filters: map[query.Field]query.Value{
					query.FieldNamespace: query.Value("default"),
					fieldPVCName:         query.Value(foo4.Spec.Volumes[0].PersistentVolumeClaim.ClaimName),
				},
			},
			&api.ListResult{
				Items:      []interface{}{foo4},
				TotalItems: 1,
			},
			nil,
		},
		{
			"test phase filter",
			"default",
			&query.Query{
				Pagination: &query.Pagination{
					Limit:  10,
					Offset: 0,
				},
				SortBy:    query.FieldName,
				Ascending: false,
				Filters: map[query.Field]query.Value{
					query.FieldNamespace: query.Value("default"),
					fieldPhase:           query.Value(corev1.PodRunning),
				},
			},
			&api.ListResult{
				Items:      []interface{}{foo5},
				TotalItems: 1,
			},
			nil,
		},
	}

	getter := prepare()

	for _, test := range tests {
		got, err := getter.List(test.namespace, test.query)
		if test.expectedErr != nil && err != test.expectedErr {
			t.Errorf("expected error, got nothing")
		} else if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, test.expected); diff != "" {
			t.Errorf("%T differ (-got, +want): %s", test.expected, diff)
		}
	}
}

var (
	foo1 = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo1",
			Namespace: "default",
		},
	}
	foo2 = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo2",
			Namespace: "default",
		},
	}
	foo3 = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo3",
			Namespace: "default",
		},
	}
	foo4 = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo4",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{
				{
					Name: "data",
					VolumeSource: corev1.VolumeSource{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: "pvc-1",
							ReadOnly:  false,
						},
					},
				},
			},
		},
	}
	foo5 = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo5",
			Namespace: "default",
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
		},
	}

	pods = []interface{}{foo1, foo2, foo3, foo4, foo5}
)

func prepare() v1alpha3.Interface {

	client := fake.NewSimpleClientset()
	informer := informers.NewSharedInformerFactory(client, 0)

	for _, pod := range pods {
		_ = informer.Core().V1().Pods().Informer().GetIndexer().Add(pod)
	}

	return New(informer)
}
