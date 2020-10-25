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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	miniocontrollerv1beta1 "github.com/minio/minio-operator/pkg/apis/miniocontroller/v1beta1"
	versioned "go.f110.dev/mono/controllers/minio-extra-operator/pkg/client/versioned"
	internalinterfaces "go.f110.dev/mono/controllers/minio-extra-operator/pkg/informers/externalversions/internalinterfaces"
	v1beta1 "go.f110.dev/mono/controllers/minio-extra-operator/pkg/listers/miniocontroller/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MinIOInstanceInformer provides access to a shared informer and lister for
// MinIOInstances.
type MinIOInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.MinIOInstanceLister
}

type minIOInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMinIOInstanceInformer constructs a new informer for MinIOInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMinIOInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMinIOInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMinIOInstanceInformer constructs a new informer for MinIOInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMinIOInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MinV1beta1().MinIOInstances(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MinV1beta1().MinIOInstances(namespace).Watch(context.TODO(), options)
			},
		},
		&miniocontrollerv1beta1.MinIOInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *minIOInstanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMinIOInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *minIOInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&miniocontrollerv1beta1.MinIOInstance{}, f.defaultInformer)
}

func (f *minIOInstanceInformer) Lister() v1beta1.MinIOInstanceLister {
	return v1beta1.NewMinIOInstanceLister(f.Informer().GetIndexer())
}
