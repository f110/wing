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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/f110/wing/controllers/harbor-project-operator/pkg/api/harbor/v1alpha1"
	scheme "github.com/f110/wing/controllers/harbor-project-operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HarborProjectsGetter has a method to return a HarborProjectInterface.
// A group's client should implement this interface.
type HarborProjectsGetter interface {
	HarborProjects(namespace string) HarborProjectInterface
}

// HarborProjectInterface has methods to work with HarborProject resources.
type HarborProjectInterface interface {
	Create(*v1alpha1.HarborProject) (*v1alpha1.HarborProject, error)
	Update(*v1alpha1.HarborProject) (*v1alpha1.HarborProject, error)
	UpdateStatus(*v1alpha1.HarborProject) (*v1alpha1.HarborProject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HarborProject, error)
	List(opts v1.ListOptions) (*v1alpha1.HarborProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HarborProject, err error)
	HarborProjectExpansion
}

// harborProjects implements HarborProjectInterface
type harborProjects struct {
	client rest.Interface
	ns     string
}

// newHarborProjects returns a HarborProjects
func newHarborProjects(c *HarborV1alpha1Client, namespace string) *harborProjects {
	return &harborProjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the harborProject, and returns the corresponding harborProject object, and an error if there is any.
func (c *harborProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.HarborProject, err error) {
	result = &v1alpha1.HarborProject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("harborprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HarborProjects that match those selectors.
func (c *harborProjects) List(opts v1.ListOptions) (result *v1alpha1.HarborProjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HarborProjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("harborprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested harborProjects.
func (c *harborProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("harborprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a harborProject and creates it.  Returns the server's representation of the harborProject, and an error, if there is any.
func (c *harborProjects) Create(harborProject *v1alpha1.HarborProject) (result *v1alpha1.HarborProject, err error) {
	result = &v1alpha1.HarborProject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("harborprojects").
		Body(harborProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a harborProject and updates it. Returns the server's representation of the harborProject, and an error, if there is any.
func (c *harborProjects) Update(harborProject *v1alpha1.HarborProject) (result *v1alpha1.HarborProject, err error) {
	result = &v1alpha1.HarborProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("harborprojects").
		Name(harborProject.Name).
		Body(harborProject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *harborProjects) UpdateStatus(harborProject *v1alpha1.HarborProject) (result *v1alpha1.HarborProject, err error) {
	result = &v1alpha1.HarborProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("harborprojects").
		Name(harborProject.Name).
		SubResource("status").
		Body(harborProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the harborProject and deletes it. Returns an error if one occurs.
func (c *harborProjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("harborprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *harborProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("harborprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched harborProject.
func (c *harborProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HarborProject, err error) {
	result = &v1alpha1.HarborProject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("harborprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}