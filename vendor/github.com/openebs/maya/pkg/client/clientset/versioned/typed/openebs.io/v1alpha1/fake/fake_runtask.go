/*
Copyright 2018 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRunTasks implements RunTaskInterface
type FakeRunTasks struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var runtasksResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "runtasks"}

var runtasksKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "RunTask"}

// Get takes name of the runTask, and returns the corresponding runTask object, and an error if there is any.
func (c *FakeRunTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.RunTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(runtasksResource, c.ns, name), &v1alpha1.RunTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RunTask), err
}

// List takes label and field selectors, and returns the list of RunTasks that match those selectors.
func (c *FakeRunTasks) List(opts v1.ListOptions) (result *v1alpha1.RunTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(runtasksResource, runtasksKind, c.ns, opts), &v1alpha1.RunTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RunTaskList{ListMeta: obj.(*v1alpha1.RunTaskList).ListMeta}
	for _, item := range obj.(*v1alpha1.RunTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runTasks.
func (c *FakeRunTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(runtasksResource, c.ns, opts))

}

// Create takes the representation of a runTask and creates it.  Returns the server's representation of the runTask, and an error, if there is any.
func (c *FakeRunTasks) Create(runTask *v1alpha1.RunTask) (result *v1alpha1.RunTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(runtasksResource, c.ns, runTask), &v1alpha1.RunTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RunTask), err
}

// Update takes the representation of a runTask and updates it. Returns the server's representation of the runTask, and an error, if there is any.
func (c *FakeRunTasks) Update(runTask *v1alpha1.RunTask) (result *v1alpha1.RunTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(runtasksResource, c.ns, runTask), &v1alpha1.RunTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RunTask), err
}

// Delete takes name of the runTask and deletes it. Returns an error if one occurs.
func (c *FakeRunTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(runtasksResource, c.ns, name), &v1alpha1.RunTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRunTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(runtasksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RunTaskList{})
	return err
}

// Patch applies the patch and returns the patched runTask.
func (c *FakeRunTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RunTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(runtasksResource, c.ns, name, pt, data, subresources...), &v1alpha1.RunTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RunTask), err
}
