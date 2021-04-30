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

package versioned

import (
	"fmt"

	consulv1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/consul/v1alpha1"
	grafanav1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/grafana/v1alpha1"
	harborv1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/harbor/v1alpha1"
	miniov1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/minio/v1alpha1"
	miniocontrollerv1beta1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/miniocontroller/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ConsulV1alpha1() consulv1alpha1.ConsulV1alpha1Interface
	GrafanaV1alpha1() grafanav1alpha1.GrafanaV1alpha1Interface
	HarborV1alpha1() harborv1alpha1.HarborV1alpha1Interface
	MinioV1alpha1() miniov1alpha1.MinioV1alpha1Interface
	MiniocontrollerV1beta1() miniocontrollerv1beta1.MiniocontrollerV1beta1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	consulV1alpha1         *consulv1alpha1.ConsulV1alpha1Client
	grafanaV1alpha1        *grafanav1alpha1.GrafanaV1alpha1Client
	harborV1alpha1         *harborv1alpha1.HarborV1alpha1Client
	minioV1alpha1          *miniov1alpha1.MinioV1alpha1Client
	miniocontrollerV1beta1 *miniocontrollerv1beta1.MiniocontrollerV1beta1Client
}

// ConsulV1alpha1 retrieves the ConsulV1alpha1Client
func (c *Clientset) ConsulV1alpha1() consulv1alpha1.ConsulV1alpha1Interface {
	return c.consulV1alpha1
}

// GrafanaV1alpha1 retrieves the GrafanaV1alpha1Client
func (c *Clientset) GrafanaV1alpha1() grafanav1alpha1.GrafanaV1alpha1Interface {
	return c.grafanaV1alpha1
}

// HarborV1alpha1 retrieves the HarborV1alpha1Client
func (c *Clientset) HarborV1alpha1() harborv1alpha1.HarborV1alpha1Interface {
	return c.harborV1alpha1
}

// MinioV1alpha1 retrieves the MinioV1alpha1Client
func (c *Clientset) MinioV1alpha1() miniov1alpha1.MinioV1alpha1Interface {
	return c.minioV1alpha1
}

// MiniocontrollerV1beta1 retrieves the MiniocontrollerV1beta1Client
func (c *Clientset) MiniocontrollerV1beta1() miniocontrollerv1beta1.MiniocontrollerV1beta1Interface {
	return c.miniocontrollerV1beta1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.consulV1alpha1, err = consulv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.grafanaV1alpha1, err = grafanav1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.harborV1alpha1, err = harborv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.minioV1alpha1, err = miniov1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.miniocontrollerV1beta1, err = miniocontrollerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.consulV1alpha1 = consulv1alpha1.NewForConfigOrDie(c)
	cs.grafanaV1alpha1 = grafanav1alpha1.NewForConfigOrDie(c)
	cs.harborV1alpha1 = harborv1alpha1.NewForConfigOrDie(c)
	cs.minioV1alpha1 = miniov1alpha1.NewForConfigOrDie(c)
	cs.miniocontrollerV1beta1 = miniocontrollerv1beta1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.consulV1alpha1 = consulv1alpha1.New(c)
	cs.grafanaV1alpha1 = grafanav1alpha1.New(c)
	cs.harborV1alpha1 = harborv1alpha1.New(c)
	cs.minioV1alpha1 = miniov1alpha1.New(c)
	cs.miniocontrollerV1beta1 = miniocontrollerv1beta1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
