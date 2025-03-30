/*
Copyright The Kubernetes Authors.

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
	context "context"
	time "time"

	apisapplesv1alpha1 "github.com/example/custom-operator/pkg/apis/apples/v1alpha1"
	versioned "github.com/example/custom-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/example/custom-operator/pkg/generated/informers/externalversions/internalinterfaces"
	applesv1alpha1 "github.com/example/custom-operator/pkg/generated/listers/apples/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppleInformer provides access to a shared informer and lister for
// Apples.
type AppleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() applesv1alpha1.AppleLister
}

type appleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppleInformer constructs a new informer for Apple type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppleInformer constructs a new informer for Apple type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().Apples(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().Apples(namespace).Watch(context.Background(), options)
			},
		},
		&apisapplesv1alpha1.Apple{},
		resyncPeriod,
		indexers,
	)
}

func (f *appleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisapplesv1alpha1.Apple{}, f.defaultInformer)
}

func (f *appleInformer) Lister() applesv1alpha1.AppleLister {
	return applesv1alpha1.NewAppleLister(f.Informer().GetIndexer())
}
