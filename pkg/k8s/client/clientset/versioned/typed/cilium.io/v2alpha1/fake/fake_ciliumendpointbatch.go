// Copyright 2017-2021 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumEndpointBatches implements CiliumEndpointBatchInterface
type FakeCiliumEndpointBatches struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumendpointbatchesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumendpointbatches"}

var ciliumendpointbatchesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2alpha1", Kind: "CiliumEndpointBatch"}

// Get takes name of the ciliumEndpointBatch, and returns the corresponding ciliumEndpointBatch object, and an error if there is any.
func (c *FakeCiliumEndpointBatches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumEndpointBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumendpointbatchesResource, name), &v2alpha1.CiliumEndpointBatch{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointBatch), err
}

// List takes label and field selectors, and returns the list of CiliumEndpointBatches that match those selectors.
func (c *FakeCiliumEndpointBatches) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumEndpointBatchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumendpointbatchesResource, ciliumendpointbatchesKind, opts), &v2alpha1.CiliumEndpointBatchList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumEndpointBatchList{ListMeta: obj.(*v2alpha1.CiliumEndpointBatchList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumEndpointBatchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEndpointBatches.
func (c *FakeCiliumEndpointBatches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumendpointbatchesResource, opts))
}

// Create takes the representation of a ciliumEndpointBatch and creates it.  Returns the server's representation of the ciliumEndpointBatch, and an error, if there is any.
func (c *FakeCiliumEndpointBatches) Create(ctx context.Context, ciliumEndpointBatch *v2alpha1.CiliumEndpointBatch, opts v1.CreateOptions) (result *v2alpha1.CiliumEndpointBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumendpointbatchesResource, ciliumEndpointBatch), &v2alpha1.CiliumEndpointBatch{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointBatch), err
}

// Update takes the representation of a ciliumEndpointBatch and updates it. Returns the server's representation of the ciliumEndpointBatch, and an error, if there is any.
func (c *FakeCiliumEndpointBatches) Update(ctx context.Context, ciliumEndpointBatch *v2alpha1.CiliumEndpointBatch, opts v1.UpdateOptions) (result *v2alpha1.CiliumEndpointBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumendpointbatchesResource, ciliumEndpointBatch), &v2alpha1.CiliumEndpointBatch{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointBatch), err
}

// Delete takes name of the ciliumEndpointBatch and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEndpointBatches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ciliumendpointbatchesResource, name), &v2alpha1.CiliumEndpointBatch{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEndpointBatches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumendpointbatchesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumEndpointBatchList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEndpointBatch.
func (c *FakeCiliumEndpointBatches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEndpointBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumendpointbatchesResource, name, pt, data, subresources...), &v2alpha1.CiliumEndpointBatch{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointBatch), err
}