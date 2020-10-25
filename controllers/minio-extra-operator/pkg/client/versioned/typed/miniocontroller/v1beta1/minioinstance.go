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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/minio/minio-operator/pkg/apis/miniocontroller/v1beta1"
	scheme "go.f110.dev/mono/controllers/minio-extra-operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MinIOInstancesGetter has a method to return a MinIOInstanceInterface.
// A group's client should implement this interface.
type MinIOInstancesGetter interface {
	MinIOInstances(namespace string) MinIOInstanceInterface
}

// MinIOInstanceInterface has methods to work with MinIOInstance resources.
type MinIOInstanceInterface interface {
	Create(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.CreateOptions) (*v1beta1.MinIOInstance, error)
	Update(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.UpdateOptions) (*v1beta1.MinIOInstance, error)
	UpdateStatus(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.UpdateOptions) (*v1beta1.MinIOInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.MinIOInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MinIOInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MinIOInstance, err error)
	MinIOInstanceExpansion
}

// minIOInstances implements MinIOInstanceInterface
type minIOInstances struct {
	client rest.Interface
	ns     string
}

// newMinIOInstances returns a MinIOInstances
func newMinIOInstances(c *MinV1beta1Client, namespace string) *minIOInstances {
	return &minIOInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the minIOInstance, and returns the corresponding minIOInstance object, and an error if there is any.
func (c *minIOInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MinIOInstance, err error) {
	result = &v1beta1.MinIOInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("minioinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MinIOInstances that match those selectors.
func (c *minIOInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MinIOInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MinIOInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("minioinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested minIOInstances.
func (c *minIOInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("minioinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a minIOInstance and creates it.  Returns the server's representation of the minIOInstance, and an error, if there is any.
func (c *minIOInstances) Create(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.CreateOptions) (result *v1beta1.MinIOInstance, err error) {
	result = &v1beta1.MinIOInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("minioinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(minIOInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a minIOInstance and updates it. Returns the server's representation of the minIOInstance, and an error, if there is any.
func (c *minIOInstances) Update(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.UpdateOptions) (result *v1beta1.MinIOInstance, err error) {
	result = &v1beta1.MinIOInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("minioinstances").
		Name(minIOInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(minIOInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *minIOInstances) UpdateStatus(ctx context.Context, minIOInstance *v1beta1.MinIOInstance, opts v1.UpdateOptions) (result *v1beta1.MinIOInstance, err error) {
	result = &v1beta1.MinIOInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("minioinstances").
		Name(minIOInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(minIOInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the minIOInstance and deletes it. Returns an error if one occurs.
func (c *minIOInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("minioinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *minIOInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("minioinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched minIOInstance.
func (c *minIOInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MinIOInstance, err error) {
	result = &v1beta1.MinIOInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("minioinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
