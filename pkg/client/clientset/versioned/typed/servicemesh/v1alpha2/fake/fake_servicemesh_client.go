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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/servicemesh/v1alpha2"
)

type FakeServicemeshV1alpha2 struct {
	*testing.Fake
}

func (c *FakeServicemeshV1alpha2) ServicePolicies(namespace string) v1alpha2.ServicePolicyInterface {
	return &FakeServicePolicies{c, namespace}
}

func (c *FakeServicemeshV1alpha2) Strategies(namespace string) v1alpha2.StrategyInterface {
	return &FakeStrategies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicemeshV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
