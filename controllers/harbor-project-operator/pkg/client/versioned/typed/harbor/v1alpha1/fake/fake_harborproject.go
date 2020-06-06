/*
MIT License

Copyright (c) 2020 Fumihiro Ito

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "go.f110.dev/mono/controllers/harbor-project-operator/pkg/api/harbor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHarborProjects implements HarborProjectInterface
type FakeHarborProjects struct {
	Fake *FakeHarborV1alpha1
	ns   string
}

var harborprojectsResource = schema.GroupVersionResource{Group: "harbor.f110.dev", Version: "v1alpha1", Resource: "harborprojects"}

var harborprojectsKind = schema.GroupVersionKind{Group: "harbor.f110.dev", Version: "v1alpha1", Kind: "HarborProject"}

// Get takes name of the harborProject, and returns the corresponding harborProject object, and an error if there is any.
func (c *FakeHarborProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.HarborProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(harborprojectsResource, c.ns, name), &v1alpha1.HarborProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HarborProject), err
}

// List takes label and field selectors, and returns the list of HarborProjects that match those selectors.
func (c *FakeHarborProjects) List(opts v1.ListOptions) (result *v1alpha1.HarborProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(harborprojectsResource, harborprojectsKind, c.ns, opts), &v1alpha1.HarborProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HarborProjectList{ListMeta: obj.(*v1alpha1.HarborProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.HarborProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested harborProjects.
func (c *FakeHarborProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(harborprojectsResource, c.ns, opts))

}

// Create takes the representation of a harborProject and creates it.  Returns the server's representation of the harborProject, and an error, if there is any.
func (c *FakeHarborProjects) Create(harborProject *v1alpha1.HarborProject) (result *v1alpha1.HarborProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(harborprojectsResource, c.ns, harborProject), &v1alpha1.HarborProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HarborProject), err
}

// Update takes the representation of a harborProject and updates it. Returns the server's representation of the harborProject, and an error, if there is any.
func (c *FakeHarborProjects) Update(harborProject *v1alpha1.HarborProject) (result *v1alpha1.HarborProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(harborprojectsResource, c.ns, harborProject), &v1alpha1.HarborProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HarborProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHarborProjects) UpdateStatus(harborProject *v1alpha1.HarborProject) (*v1alpha1.HarborProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(harborprojectsResource, "status", c.ns, harborProject), &v1alpha1.HarborProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HarborProject), err
}

// Delete takes name of the harborProject and deletes it. Returns an error if one occurs.
func (c *FakeHarborProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(harborprojectsResource, c.ns, name), &v1alpha1.HarborProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHarborProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(harborprojectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HarborProjectList{})
	return err
}

// Patch applies the patch and returns the patched harborProject.
func (c *FakeHarborProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HarborProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(harborprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HarborProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HarborProject), err
}
