/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/platform/v1"
)

// ClusterCredentialLister helps list ClusterCredentials.
type ClusterCredentialLister interface {
	// List lists all ClusterCredentials in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterCredential, err error)
	// Get retrieves the ClusterCredential from the index for a given name.
	Get(name string) (*v1.ClusterCredential, error)
	ClusterCredentialListerExpansion
}

// clusterCredentialLister implements the ClusterCredentialLister interface.
type clusterCredentialLister struct {
	indexer cache.Indexer
}

// NewClusterCredentialLister returns a new ClusterCredentialLister.
func NewClusterCredentialLister(indexer cache.Indexer) ClusterCredentialLister {
	return &clusterCredentialLister{indexer: indexer}
}

// List lists all ClusterCredentials in the indexer.
func (s *clusterCredentialLister) List(selector labels.Selector) (ret []*v1.ClusterCredential, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterCredential))
	})
	return ret, err
}

// Get retrieves the ClusterCredential from the index for a given name.
func (s *clusterCredentialLister) Get(name string) (*v1.ClusterCredential, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustercredential"), name)
	}
	return obj.(*v1.ClusterCredential), nil
}
