/*
Copyright 2019 The KubeSphere authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/network/v1alpha1"
)

// FakeNamespaceNetworkPolicies implements NamespaceNetworkPolicyInterface
type FakeNamespaceNetworkPolicies struct {
	Fake *FakeNetworkV1alpha1
	ns   string
}

var namespacenetworkpoliciesResource = schema.GroupVersionResource{Group: "network.kubesphere.io", Version: "v1alpha1", Resource: "namespacenetworkpolicies"}

var namespacenetworkpoliciesKind = schema.GroupVersionKind{Group: "network.kubesphere.io", Version: "v1alpha1", Kind: "NamespaceNetworkPolicy"}

// Get takes name of the namespaceNetworkPolicy, and returns the corresponding namespaceNetworkPolicy object, and an error if there is any.
func (c *FakeNamespaceNetworkPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.NamespaceNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namespacenetworkpoliciesResource, c.ns, name), &v1alpha1.NamespaceNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NamespaceNetworkPolicies that match those selectors.
func (c *FakeNamespaceNetworkPolicies) List(opts v1.ListOptions) (result *v1alpha1.NamespaceNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namespacenetworkpoliciesResource, namespacenetworkpoliciesKind, c.ns, opts), &v1alpha1.NamespaceNetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NamespaceNetworkPolicyList{ListMeta: obj.(*v1alpha1.NamespaceNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.NamespaceNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaceNetworkPolicies.
func (c *FakeNamespaceNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namespacenetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a namespaceNetworkPolicy and creates it.  Returns the server's representation of the namespaceNetworkPolicy, and an error, if there is any.
func (c *FakeNamespaceNetworkPolicies) Create(namespaceNetworkPolicy *v1alpha1.NamespaceNetworkPolicy) (result *v1alpha1.NamespaceNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namespacenetworkpoliciesResource, c.ns, namespaceNetworkPolicy), &v1alpha1.NamespaceNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceNetworkPolicy), err
}

// Update takes the representation of a namespaceNetworkPolicy and updates it. Returns the server's representation of the namespaceNetworkPolicy, and an error, if there is any.
func (c *FakeNamespaceNetworkPolicies) Update(namespaceNetworkPolicy *v1alpha1.NamespaceNetworkPolicy) (result *v1alpha1.NamespaceNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namespacenetworkpoliciesResource, c.ns, namespaceNetworkPolicy), &v1alpha1.NamespaceNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceNetworkPolicy), err
}

// Delete takes name of the namespaceNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNamespaceNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(namespacenetworkpoliciesResource, c.ns, name), &v1alpha1.NamespaceNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamespaceNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namespacenetworkpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NamespaceNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched namespaceNetworkPolicy.
func (c *FakeNamespaceNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NamespaceNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namespacenetworkpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NamespaceNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceNetworkPolicy), err
}
