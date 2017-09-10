/*
Copyright 2017 The KubeDB Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DormantDatabaseLister helps list DormantDatabases.
type DormantDatabaseLister interface {
	// List lists all DormantDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DormantDatabase, err error)
	// DormantDatabases returns an object that can list and get DormantDatabases.
	DormantDatabases(namespace string) DormantDatabaseNamespaceLister
	DormantDatabaseListerExpansion
}

// dormantDatabaseLister implements the DormantDatabaseLister interface.
type dormantDatabaseLister struct {
	indexer cache.Indexer
}

// NewDormantDatabaseLister returns a new DormantDatabaseLister.
func NewDormantDatabaseLister(indexer cache.Indexer) DormantDatabaseLister {
	return &dormantDatabaseLister{indexer: indexer}
}

// List lists all DormantDatabases in the indexer.
func (s *dormantDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.DormantDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DormantDatabase))
	})
	return ret, err
}

// DormantDatabases returns an object that can list and get DormantDatabases.
func (s *dormantDatabaseLister) DormantDatabases(namespace string) DormantDatabaseNamespaceLister {
	return dormantDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DormantDatabaseNamespaceLister helps list and get DormantDatabases.
type DormantDatabaseNamespaceLister interface {
	// List lists all DormantDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DormantDatabase, err error)
	// Get retrieves the DormantDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DormantDatabase, error)
	DormantDatabaseNamespaceListerExpansion
}

// dormantDatabaseNamespaceLister implements the DormantDatabaseNamespaceLister
// interface.
type dormantDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DormantDatabases in the indexer for a given namespace.
func (s dormantDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DormantDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DormantDatabase))
	})
	return ret, err
}

// Get retrieves the DormantDatabase from the indexer for a given namespace and name.
func (s dormantDatabaseNamespaceLister) Get(name string) (*v1alpha1.DormantDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dormantdatabase"), name)
	}
	return obj.(*v1alpha1.DormantDatabase), nil
}
