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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	corev1 "github.com/jwebb1334/fission/pkg/apis/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTimeTriggers implements TimeTriggerInterface
type FakeTimeTriggers struct {
	Fake *FakeCoreV1
	ns   string
}

var timetriggersResource = schema.GroupVersionResource{Group: "fission.io", Version: "v1", Resource: "timetriggers"}

var timetriggersKind = schema.GroupVersionKind{Group: "fission.io", Version: "v1", Kind: "TimeTrigger"}

// Get takes name of the _timeTrigger, and returns the corresponding timeTrigger object, and an error if there is any.
func (c *FakeTimeTriggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.TimeTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(timetriggersResource, c.ns, name), &corev1.TimeTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.TimeTrigger), err
}

// List takes label and field selectors, and returns the list of TimeTriggers that match those selectors.
func (c *FakeTimeTriggers) List(ctx context.Context, opts v1.ListOptions) (result *corev1.TimeTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(timetriggersResource, timetriggersKind, c.ns, opts), &corev1.TimeTriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.TimeTriggerList{ListMeta: obj.(*corev1.TimeTriggerList).ListMeta}
	for _, item := range obj.(*corev1.TimeTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested timeTriggers.
func (c *FakeTimeTriggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(timetriggersResource, c.ns, opts))

}

// Create takes the representation of a _timeTrigger and creates it.  Returns the server's representation of the timeTrigger, and an error, if there is any.
func (c *FakeTimeTriggers) Create(ctx context.Context, _timeTrigger *corev1.TimeTrigger, opts v1.CreateOptions) (result *corev1.TimeTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(timetriggersResource, c.ns, _timeTrigger), &corev1.TimeTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.TimeTrigger), err
}

// Update takes the representation of a _timeTrigger and updates it. Returns the server's representation of the timeTrigger, and an error, if there is any.
func (c *FakeTimeTriggers) Update(ctx context.Context, _timeTrigger *corev1.TimeTrigger, opts v1.UpdateOptions) (result *corev1.TimeTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(timetriggersResource, c.ns, _timeTrigger), &corev1.TimeTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.TimeTrigger), err
}

// Delete takes name of the _timeTrigger and deletes it. Returns an error if one occurs.
func (c *FakeTimeTriggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(timetriggersResource, c.ns, name, opts), &corev1.TimeTrigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTimeTriggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(timetriggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.TimeTriggerList{})
	return err
}

// Patch applies the patch and returns the patched timeTrigger.
func (c *FakeTimeTriggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.TimeTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(timetriggersResource, c.ns, name, pt, data, subresources...), &corev1.TimeTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.TimeTrigger), err
}
