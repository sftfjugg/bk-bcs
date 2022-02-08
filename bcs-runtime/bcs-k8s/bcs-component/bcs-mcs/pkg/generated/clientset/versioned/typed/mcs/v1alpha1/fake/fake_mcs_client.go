/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-mcs/pkg/generated/clientset/versioned/typed/mcs/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMcsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMcsV1alpha1) Manifests(namespace string) v1alpha1.ManifestInterface {
	return &FakeManifests{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMcsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
