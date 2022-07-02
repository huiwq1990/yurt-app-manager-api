/*
Copyright 2020 The OpenYurt Authors.

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
	v1alpha1 "github.com/openyurtio/yurt-app-manager-api/pkg/yurtappmanager/client/clientset/versioned/typed/apps/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAppsV1alpha1) NodePools() v1alpha1.NodePoolInterface {
	return &FakeNodePools{c}
}

func (c *FakeAppsV1alpha1) YurtAppDaemons(namespace string) v1alpha1.YurtAppDaemonInterface {
	return &FakeYurtAppDaemons{c, namespace}
}

func (c *FakeAppsV1alpha1) YurtAppSets(namespace string) v1alpha1.YurtAppSetInterface {
	return &FakeYurtAppSets{c, namespace}
}

func (c *FakeAppsV1alpha1) YurtIngresses() v1alpha1.YurtIngressInterface {
	return &FakeYurtIngresses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
