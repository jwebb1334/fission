/*
Copyright 2016 The Fission Authors.

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

package fake

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	fv1 "github.com/jwebb1334/fission/pkg/apis/core/v1"
	v1 "github.com/jwebb1334/fission/pkg/controller/client/v1"
)

type (
	FakeKubeWatcher struct{}
)

func newKubeWatcher(c *v1.V1) v1.KubeWatcherInterface {
	return &FakeKubeWatcher{}
}

func (c *FakeKubeWatcher) Create(w *fv1.KubernetesWatchTrigger) (*metav1.ObjectMeta, error) {
	return nil, nil
}

func (c *FakeKubeWatcher) Get(m *metav1.ObjectMeta) (*fv1.KubernetesWatchTrigger, error) {
	return nil, nil
}

func (c *FakeKubeWatcher) Update(w *fv1.KubernetesWatchTrigger) (*metav1.ObjectMeta, error) {
	return nil, nil
}

func (c *FakeKubeWatcher) Delete(m *metav1.ObjectMeta) error {
	return nil
}

func (c *FakeKubeWatcher) List(ns string) ([]fv1.KubernetesWatchTrigger, error) {
	return nil, nil
}
