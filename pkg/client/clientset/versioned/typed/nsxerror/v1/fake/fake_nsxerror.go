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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	nsxerrorv1 "github.com/yktsubo/nsxerror/pkg/apis/nsxerror/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNSXErrors implements NSXErrorInterface
type FakeNSXErrors struct {
	Fake *FakeVmwareV1
	ns   string
}

var nsxerrorsResource = schema.GroupVersionResource{Group: "vmware.eng.com", Version: "v1", Resource: "nsxerrors"}

var nsxerrorsKind = schema.GroupVersionKind{Group: "vmware.eng.com", Version: "v1", Kind: "NSXError"}

// Get takes name of the nSXError, and returns the corresponding nSXError object, and an error if there is any.
func (c *FakeNSXErrors) Get(name string, options v1.GetOptions) (result *nsxerrorv1.NSXError, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nsxerrorsResource, c.ns, name), &nsxerrorv1.NSXError{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nsxerrorv1.NSXError), err
}

// List takes label and field selectors, and returns the list of NSXErrors that match those selectors.
func (c *FakeNSXErrors) List(opts v1.ListOptions) (result *nsxerrorv1.NSXErrorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nsxerrorsResource, nsxerrorsKind, c.ns, opts), &nsxerrorv1.NSXErrorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nsxerrorv1.NSXErrorList{ListMeta: obj.(*nsxerrorv1.NSXErrorList).ListMeta}
	for _, item := range obj.(*nsxerrorv1.NSXErrorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nSXErrors.
func (c *FakeNSXErrors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nsxerrorsResource, c.ns, opts))

}

// Create takes the representation of a nSXError and creates it.  Returns the server's representation of the nSXError, and an error, if there is any.
func (c *FakeNSXErrors) Create(nSXError *nsxerrorv1.NSXError) (result *nsxerrorv1.NSXError, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nsxerrorsResource, c.ns, nSXError), &nsxerrorv1.NSXError{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nsxerrorv1.NSXError), err
}

// Update takes the representation of a nSXError and updates it. Returns the server's representation of the nSXError, and an error, if there is any.
func (c *FakeNSXErrors) Update(nSXError *nsxerrorv1.NSXError) (result *nsxerrorv1.NSXError, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nsxerrorsResource, c.ns, nSXError), &nsxerrorv1.NSXError{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nsxerrorv1.NSXError), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNSXErrors) UpdateStatus(nSXError *nsxerrorv1.NSXError) (*nsxerrorv1.NSXError, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nsxerrorsResource, "status", c.ns, nSXError), &nsxerrorv1.NSXError{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nsxerrorv1.NSXError), err
}

// Delete takes name of the nSXError and deletes it. Returns an error if one occurs.
func (c *FakeNSXErrors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nsxerrorsResource, c.ns, name), &nsxerrorv1.NSXError{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNSXErrors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nsxerrorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &nsxerrorv1.NSXErrorList{})
	return err
}

// Patch applies the patch and returns the patched nSXError.
func (c *FakeNSXErrors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *nsxerrorv1.NSXError, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nsxerrorsResource, c.ns, name, pt, data, subresources...), &nsxerrorv1.NSXError{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nsxerrorv1.NSXError), err
}
