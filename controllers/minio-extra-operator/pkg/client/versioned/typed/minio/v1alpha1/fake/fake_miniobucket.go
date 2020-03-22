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
	v1alpha1 "github.com/f110/tools/controllers/minio-extra-operator/pkg/api/minio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMinIOBuckets implements MinIOBucketInterface
type FakeMinIOBuckets struct {
	Fake *FakeMinioV1alpha1
	ns   string
}

var miniobucketsResource = schema.GroupVersionResource{Group: "minio.f110.dev", Version: "v1alpha1", Resource: "miniobuckets"}

var miniobucketsKind = schema.GroupVersionKind{Group: "minio.f110.dev", Version: "v1alpha1", Kind: "MinIOBucket"}

// Get takes name of the minIOBucket, and returns the corresponding minIOBucket object, and an error if there is any.
func (c *FakeMinIOBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.MinIOBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(miniobucketsResource, c.ns, name), &v1alpha1.MinIOBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MinIOBucket), err
}

// List takes label and field selectors, and returns the list of MinIOBuckets that match those selectors.
func (c *FakeMinIOBuckets) List(opts v1.ListOptions) (result *v1alpha1.MinIOBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(miniobucketsResource, miniobucketsKind, c.ns, opts), &v1alpha1.MinIOBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MinIOBucketList{ListMeta: obj.(*v1alpha1.MinIOBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.MinIOBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested minIOBuckets.
func (c *FakeMinIOBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(miniobucketsResource, c.ns, opts))

}

// Create takes the representation of a minIOBucket and creates it.  Returns the server's representation of the minIOBucket, and an error, if there is any.
func (c *FakeMinIOBuckets) Create(minIOBucket *v1alpha1.MinIOBucket) (result *v1alpha1.MinIOBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(miniobucketsResource, c.ns, minIOBucket), &v1alpha1.MinIOBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MinIOBucket), err
}

// Update takes the representation of a minIOBucket and updates it. Returns the server's representation of the minIOBucket, and an error, if there is any.
func (c *FakeMinIOBuckets) Update(minIOBucket *v1alpha1.MinIOBucket) (result *v1alpha1.MinIOBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(miniobucketsResource, c.ns, minIOBucket), &v1alpha1.MinIOBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MinIOBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMinIOBuckets) UpdateStatus(minIOBucket *v1alpha1.MinIOBucket) (*v1alpha1.MinIOBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(miniobucketsResource, "status", c.ns, minIOBucket), &v1alpha1.MinIOBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MinIOBucket), err
}

// Delete takes name of the minIOBucket and deletes it. Returns an error if one occurs.
func (c *FakeMinIOBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(miniobucketsResource, c.ns, name), &v1alpha1.MinIOBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMinIOBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(miniobucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MinIOBucketList{})
	return err
}

// Patch applies the patch and returns the patched minIOBucket.
func (c *FakeMinIOBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MinIOBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(miniobucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MinIOBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MinIOBucket), err
}
