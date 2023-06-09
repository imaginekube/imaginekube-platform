/*
Copyright 2023 The ImagineKube Authors.

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

package calicov3

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	calicov3 "imaginekube.com/api/network/calicov3"

	scheme "imaginekube.com/imaginekube/pkg/simple/client/network/ippool/calico/client/clientset/versioned/scheme"
)

// BlockAffinitiesGetter has a method to return a BlockAffinityInterface.
// A group's client should implement this interface.
type BlockAffinitiesGetter interface {
	BlockAffinities() BlockAffinityInterface
}

// BlockAffinityInterface has methods to work with BlockAffinity resources.
type BlockAffinityInterface interface {
	Create(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.CreateOptions) (*calicov3.BlockAffinity, error)
	Update(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.UpdateOptions) (*calicov3.BlockAffinity, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*calicov3.BlockAffinity, error)
	List(ctx context.Context, opts v1.ListOptions) (*calicov3.BlockAffinityList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *calicov3.BlockAffinity, err error)
	BlockAffinityExpansion
}

// blockAffinities implements BlockAffinityInterface
type blockAffinities struct {
	client rest.Interface
}

// newBlockAffinities returns a BlockAffinities
func newBlockAffinities(c *CrdCalicov3Client) *blockAffinities {
	return &blockAffinities{
		client: c.RESTClient(),
	}
}

// Get takes name of the blockAffinity, and returns the corresponding blockAffinity object, and an error if there is any.
func (c *blockAffinities) Get(ctx context.Context, name string, options v1.GetOptions) (result *calicov3.BlockAffinity, err error) {
	result = &calicov3.BlockAffinity{}
	err = c.client.Get().
		Resource("blockaffinities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BlockAffinities that match those selectors.
func (c *blockAffinities) List(ctx context.Context, opts v1.ListOptions) (result *calicov3.BlockAffinityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &calicov3.BlockAffinityList{}
	err = c.client.Get().
		Resource("blockaffinities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested blockAffinities.
func (c *blockAffinities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("blockaffinities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a blockAffinity and creates it.  Returns the server's representation of the blockAffinity, and an error, if there is any.
func (c *blockAffinities) Create(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.CreateOptions) (result *calicov3.BlockAffinity, err error) {
	result = &calicov3.BlockAffinity{}
	err = c.client.Post().
		Resource("blockaffinities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blockAffinity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a blockAffinity and updates it. Returns the server's representation of the blockAffinity, and an error, if there is any.
func (c *blockAffinities) Update(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.UpdateOptions) (result *calicov3.BlockAffinity, err error) {
	result = &calicov3.BlockAffinity{}
	err = c.client.Put().
		Resource("blockaffinities").
		Name(blockAffinity.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blockAffinity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the blockAffinity and deletes it. Returns an error if one occurs.
func (c *blockAffinities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("blockaffinities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *blockAffinities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("blockaffinities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched blockAffinity.
func (c *blockAffinities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *calicov3.BlockAffinity, err error) {
	result = &calicov3.BlockAffinity{}
	err = c.client.Patch(pt).
		Resource("blockaffinities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
