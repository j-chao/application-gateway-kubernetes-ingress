/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureapplicationgatewayinstanceupdatestatus/v1beta1"
)

// AzureApplicationGatewayInstanceUpdateStatusLister helps list AzureApplicationGatewayInstanceUpdateStatuses.
// All objects returned here must be treated as read-only.
type AzureApplicationGatewayInstanceUpdateStatusLister interface {
	// List lists all AzureApplicationGatewayInstanceUpdateStatuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.AzureApplicationGatewayInstanceUpdateStatus, err error)
	// Get retrieves the AzureApplicationGatewayInstanceUpdateStatus from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.AzureApplicationGatewayInstanceUpdateStatus, error)
	AzureApplicationGatewayInstanceUpdateStatusListerExpansion
}

// azureApplicationGatewayInstanceUpdateStatusLister implements the AzureApplicationGatewayInstanceUpdateStatusLister interface.
type azureApplicationGatewayInstanceUpdateStatusLister struct {
	indexer cache.Indexer
}

// NewAzureApplicationGatewayInstanceUpdateStatusLister returns a new AzureApplicationGatewayInstanceUpdateStatusLister.
func NewAzureApplicationGatewayInstanceUpdateStatusLister(indexer cache.Indexer) AzureApplicationGatewayInstanceUpdateStatusLister {
	return &azureApplicationGatewayInstanceUpdateStatusLister{indexer: indexer}
}

// List lists all AzureApplicationGatewayInstanceUpdateStatuses in the indexer.
func (s *azureApplicationGatewayInstanceUpdateStatusLister) List(selector labels.Selector) (ret []*v1beta1.AzureApplicationGatewayInstanceUpdateStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.AzureApplicationGatewayInstanceUpdateStatus))
	})
	return ret, err
}

// Get retrieves the AzureApplicationGatewayInstanceUpdateStatus from the index for a given name.
func (s *azureApplicationGatewayInstanceUpdateStatusLister) Get(name string) (*v1beta1.AzureApplicationGatewayInstanceUpdateStatus, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("azureapplicationgatewayinstanceupdatestatus"), name)
	}
	return obj.(*v1beta1.AzureApplicationGatewayInstanceUpdateStatus), nil
}