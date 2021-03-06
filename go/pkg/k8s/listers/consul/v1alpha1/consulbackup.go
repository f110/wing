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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.f110.dev/mono/go/pkg/api/consul/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsulBackupLister helps list ConsulBackups.
// All objects returned here must be treated as read-only.
type ConsulBackupLister interface {
	// List lists all ConsulBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsulBackup, err error)
	// ConsulBackups returns an object that can list and get ConsulBackups.
	ConsulBackups(namespace string) ConsulBackupNamespaceLister
	ConsulBackupListerExpansion
}

// consulBackupLister implements the ConsulBackupLister interface.
type consulBackupLister struct {
	indexer cache.Indexer
}

// NewConsulBackupLister returns a new ConsulBackupLister.
func NewConsulBackupLister(indexer cache.Indexer) ConsulBackupLister {
	return &consulBackupLister{indexer: indexer}
}

// List lists all ConsulBackups in the indexer.
func (s *consulBackupLister) List(selector labels.Selector) (ret []*v1alpha1.ConsulBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsulBackup))
	})
	return ret, err
}

// ConsulBackups returns an object that can list and get ConsulBackups.
func (s *consulBackupLister) ConsulBackups(namespace string) ConsulBackupNamespaceLister {
	return consulBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsulBackupNamespaceLister helps list and get ConsulBackups.
// All objects returned here must be treated as read-only.
type ConsulBackupNamespaceLister interface {
	// List lists all ConsulBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsulBackup, err error)
	// Get retrieves the ConsulBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConsulBackup, error)
	ConsulBackupNamespaceListerExpansion
}

// consulBackupNamespaceLister implements the ConsulBackupNamespaceLister
// interface.
type consulBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConsulBackups in the indexer for a given namespace.
func (s consulBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConsulBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsulBackup))
	})
	return ret, err
}

// Get retrieves the ConsulBackup from the indexer for a given namespace and name.
func (s consulBackupNamespaceLister) Get(name string) (*v1alpha1.ConsulBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("consulbackup"), name)
	}
	return obj.(*v1alpha1.ConsulBackup), nil
}
