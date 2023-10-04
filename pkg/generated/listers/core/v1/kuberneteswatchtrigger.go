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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jwebb1334/fission/pkg/apis/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubernetesWatchTriggerLister helps list KubernetesWatchTriggers.
// All objects returned here must be treated as read-only.
type KubernetesWatchTriggerLister interface {
	// List lists all KubernetesWatchTriggers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KubernetesWatchTrigger, err error)
	// KubernetesWatchTriggers returns an object that can list and get KubernetesWatchTriggers.
	KubernetesWatchTriggers(namespace string) KubernetesWatchTriggerNamespaceLister
	KubernetesWatchTriggerListerExpansion
}

// kubernetesWatchTriggerLister implements the KubernetesWatchTriggerLister interface.
type kubernetesWatchTriggerLister struct {
	indexer cache.Indexer
}

// NewKubernetesWatchTriggerLister returns a new KubernetesWatchTriggerLister.
func NewKubernetesWatchTriggerLister(indexer cache.Indexer) KubernetesWatchTriggerLister {
	return &kubernetesWatchTriggerLister{indexer: indexer}
}

// List lists all KubernetesWatchTriggers in the indexer.
func (s *kubernetesWatchTriggerLister) List(selector labels.Selector) (ret []*v1.KubernetesWatchTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubernetesWatchTrigger))
	})
	return ret, err
}

// KubernetesWatchTriggers returns an object that can list and get KubernetesWatchTriggers.
func (s *kubernetesWatchTriggerLister) KubernetesWatchTriggers(namespace string) KubernetesWatchTriggerNamespaceLister {
	return kubernetesWatchTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubernetesWatchTriggerNamespaceLister helps list and get KubernetesWatchTriggers.
// All objects returned here must be treated as read-only.
type KubernetesWatchTriggerNamespaceLister interface {
	// List lists all KubernetesWatchTriggers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KubernetesWatchTrigger, err error)
	// Get retrieves the KubernetesWatchTrigger from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.KubernetesWatchTrigger, error)
	KubernetesWatchTriggerNamespaceListerExpansion
}

// kubernetesWatchTriggerNamespaceLister implements the KubernetesWatchTriggerNamespaceLister
// interface.
type kubernetesWatchTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubernetesWatchTriggers in the indexer for a given namespace.
func (s kubernetesWatchTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1.KubernetesWatchTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubernetesWatchTrigger))
	})
	return ret, err
}

// Get retrieves the KubernetesWatchTrigger from the indexer for a given namespace and name.
func (s kubernetesWatchTriggerNamespaceLister) Get(name string) (*v1.KubernetesWatchTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kuberneteswatchtrigger"), name)
	}
	return obj.(*v1.KubernetesWatchTrigger), nil
}
