/*
Copyright The KCP Authors.

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

package v1alpha1

import (
	"context"
	"time"

	logicalcluster "github.com/kcp-dev/logicalcluster"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	scheme "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/scheme"
)

// APIResourceSchemasGetter has a method to return a APIResourceSchemaInterface.
// A group's client should implement this interface.
type APIResourceSchemasGetter interface {
	APIResourceSchemas() APIResourceSchemaInterface
}

// APIResourceSchemaInterface has methods to work with APIResourceSchema resources.
type APIResourceSchemaInterface interface {
	Create(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.CreateOptions) (*v1alpha1.APIResourceSchema, error)
	Update(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.UpdateOptions) (*v1alpha1.APIResourceSchema, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIResourceSchema, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIResourceSchemaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIResourceSchema, err error)
	APIResourceSchemaExpansion
}

// aPIResourceSchemas implements APIResourceSchemaInterface
type aPIResourceSchemas struct {
	client  rest.Interface
	cluster logicalcluster.Name
}

// newAPIResourceSchemas returns a APIResourceSchemas
func newAPIResourceSchemas(c *ApisV1alpha1Client) *aPIResourceSchemas {
	return &aPIResourceSchemas{
		client:  c.RESTClient(),
		cluster: c.cluster,
	}
}

// Get takes name of the aPIResourceSchema, and returns the corresponding aPIResourceSchema object, and an error if there is any.
func (c *aPIResourceSchemas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIResourceSchema, err error) {
	result = &v1alpha1.APIResourceSchema{}
	err = c.client.Get().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIResourceSchemas that match those selectors.
func (c *aPIResourceSchemas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIResourceSchemaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIResourceSchemaList{}
	err = c.client.Get().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIResourceSchemas.
func (c *aPIResourceSchemas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIResourceSchema and creates it.  Returns the server's representation of the aPIResourceSchema, and an error, if there is any.
func (c *aPIResourceSchemas) Create(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.CreateOptions) (result *v1alpha1.APIResourceSchema, err error) {
	result = &v1alpha1.APIResourceSchema{}
	err = c.client.Post().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIResourceSchema).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIResourceSchema and updates it. Returns the server's representation of the aPIResourceSchema, and an error, if there is any.
func (c *aPIResourceSchemas) Update(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.UpdateOptions) (result *v1alpha1.APIResourceSchema, err error) {
	result = &v1alpha1.APIResourceSchema{}
	err = c.client.Put().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		Name(aPIResourceSchema.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIResourceSchema).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIResourceSchema and deletes it. Returns an error if one occurs.
func (c *aPIResourceSchemas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIResourceSchemas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIResourceSchema.
func (c *aPIResourceSchemas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIResourceSchema, err error) {
	result = &v1alpha1.APIResourceSchema{}
	err = c.client.Patch(pt).
		Cluster(c.cluster).
		Resource("apiresourceschemas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
