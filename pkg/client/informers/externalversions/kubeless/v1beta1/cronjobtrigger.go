/*
Copyright (c) 2016-2017 Bitnami

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

// This file was automatically generated by informer-gen

package v1beta1

import (
	kubeless_v1beta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	versioned "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeless/kubeless/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubeless/kubeless/pkg/client/listers/kubeless/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CronJobTriggerInformer provides access to a shared informer and lister for
// CronJobTriggers.
type CronJobTriggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CronJobTriggerLister
}

type cronJobTriggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCronJobTriggerInformer constructs a new informer for CronJobTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronJobTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronJobTriggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCronJobTriggerInformer constructs a new informer for CronJobTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronJobTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubelessV1beta1().CronJobTriggers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubelessV1beta1().CronJobTriggers(namespace).Watch(options)
			},
		},
		&kubeless_v1beta1.CronJobTrigger{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronJobTriggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronJobTriggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronJobTriggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeless_v1beta1.CronJobTrigger{}, f.defaultInformer)
}

func (f *cronJobTriggerInformer) Lister() v1beta1.CronJobTriggerLister {
	return v1beta1.NewCronJobTriggerLister(f.Informer().GetIndexer())
}
