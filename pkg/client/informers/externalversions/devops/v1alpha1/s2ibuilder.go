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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	devopsv1alpha1 "imaginekube.com/api/devops/v1alpha1"
	versioned "imaginekube.com/imaginekube/pkg/client/clientset/versioned"
	internalinterfaces "imaginekube.com/imaginekube/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "imaginekube.com/imaginekube/pkg/client/listers/devops/v1alpha1"
)

// S2iBuilderInformer provides access to a shared informer and lister for
// S2iBuilders.
type S2iBuilderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.S2iBuilderLister
}

type s2iBuilderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewS2iBuilderInformer constructs a new informer for S2iBuilder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewS2iBuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredS2iBuilderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredS2iBuilderInformer constructs a new informer for S2iBuilder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredS2iBuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha1().S2iBuilders(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha1().S2iBuilders(namespace).Watch(context.TODO(), options)
			},
		},
		&devopsv1alpha1.S2iBuilder{},
		resyncPeriod,
		indexers,
	)
}

func (f *s2iBuilderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredS2iBuilderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *s2iBuilderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&devopsv1alpha1.S2iBuilder{}, f.defaultInformer)
}

func (f *s2iBuilderInformer) Lister() v1alpha1.S2iBuilderLister {
	return v1alpha1.NewS2iBuilderLister(f.Informer().GetIndexer())
}