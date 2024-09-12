/*
Copyright The Karpor Authors.

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

package v1beta1

import (
	"context"
	time "time"

	searchv1beta1 "github.com/KusionStack/karpor/pkg/kubernetes/apis/search/v1beta1"
	versioned "github.com/KusionStack/karpor/pkg/kubernetes/generated/clientset/versioned"
	internalinterfaces "github.com/KusionStack/karpor/pkg/kubernetes/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/KusionStack/karpor/pkg/kubernetes/generated/listers/search/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrimRuleInformer provides access to a shared informer and lister for
// TrimRules.
type TrimRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.TrimRuleLister
}

type trimRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTrimRuleInformer constructs a new informer for TrimRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrimRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrimRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTrimRuleInformer constructs a new informer for TrimRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrimRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SearchV1beta1().TrimRules().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SearchV1beta1().TrimRules().Watch(context.TODO(), options)
			},
		},
		&searchv1beta1.TrimRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *trimRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrimRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trimRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&searchv1beta1.TrimRule{}, f.defaultInformer)
}

func (f *trimRuleInformer) Lister() v1beta1.TrimRuleLister {
	return v1beta1.NewTrimRuleLister(f.Informer().GetIndexer())
}
