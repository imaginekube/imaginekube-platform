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

package cluster

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	clusterv1alpha1 "imaginekube.com/api/cluster/v1alpha1"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/client/clientset/versioned/fake"
	"imaginekube.com/imaginekube/pkg/client/informers/externalversions"
)

var clusters = []*clusterv1alpha1.Cluster{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo",
			Labels: map[string]string{
				"cluster.imaginekube.com/region": "beijing",
				"cluster.imaginekube.com/group":  "development",
			},
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bar",
			Labels: map[string]string{
				"cluster.imaginekube.com/region": "beijing",
				"cluster.imaginekube.com/group":  "production",
			},
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "whatever",
			Labels: map[string]string{
				"cluster.imaginekube.com/region": "shanghai",
				"cluster.imaginekube.com/group":  "testing",
			},
		},
	},
}

func clustersToInterface(clusters ...*clusterv1alpha1.Cluster) []interface{} {
	items := make([]interface{}, 0)

	for _, cluster := range clusters {
		items = append(items, cluster)
	}

	return items
}

func clustersToRuntimeObject(clusters ...*clusterv1alpha1.Cluster) []runtime.Object {
	items := make([]runtime.Object, 0)

	for _, cluster := range clusters {
		items = append(items, cluster)
	}

	return items
}

func TestClustersGetter(t *testing.T) {
	var testCases = []struct {
		description string
		query       *query.Query
		expected    *api.ListResult
	}{
		{
			description: "Test normal case",
			query: &query.Query{
				LabelSelector: "cluster.imaginekube.com/region=beijing",
				Ascending:     false,
			},
			expected: &api.ListResult{
				TotalItems: 2,
				Items:      clustersToInterface(clusters[0], clusters[1]),
			},
		},
	}

	client := fake.NewSimpleClientset(clustersToRuntimeObject(clusters...)...)
	informer := externalversions.NewSharedInformerFactory(client, 0)

	for _, cluster := range clusters {
		informer.Cluster().V1alpha1().Clusters().Informer().GetIndexer().Add(cluster)
	}

	for _, testCase := range testCases {

		clusterGetter := New(informer)
		t.Run(testCase.description, func(t *testing.T) {
			result, err := clusterGetter.List("", testCase.query)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(result, testCase.expected); len(diff) != 0 {
				t.Errorf("%T, got+ expected-, %s", testCase.expected, diff)
			}
		})
	}
}
