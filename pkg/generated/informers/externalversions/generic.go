/*
Copyright The Fission Authors.

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

package externalversions

import (
	"fmt"

	v1 "github.com/jwebb1334/fission/pkg/apis/core/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=fission.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("canaryconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().CanaryConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("environments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Environments().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("functions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Functions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("httptriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().HTTPTriggers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kuberneteswatchtriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().KubernetesWatchTriggers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("messagequeuetriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().MessageQueueTriggers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("packages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Packages().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("timetriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().TimeTriggers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
