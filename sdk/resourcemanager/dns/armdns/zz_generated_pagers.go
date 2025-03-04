//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// RecordSetsClientListAllByDNSZonePager provides operations for iterating over paged responses.
type RecordSetsClientListAllByDNSZonePager struct {
	client    *RecordSetsClient
	current   RecordSetsClientListAllByDNSZoneResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecordSetsClientListAllByDNSZoneResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecordSetsClientListAllByDNSZonePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecordSetsClientListAllByDNSZonePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecordSetListResult.NextLink == nil || len(*p.current.RecordSetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAllByDNSZoneHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecordSetsClientListAllByDNSZoneResponse page.
func (p *RecordSetsClientListAllByDNSZonePager) PageResponse() RecordSetsClientListAllByDNSZoneResponse {
	return p.current
}

// RecordSetsClientListByDNSZonePager provides operations for iterating over paged responses.
type RecordSetsClientListByDNSZonePager struct {
	client    *RecordSetsClient
	current   RecordSetsClientListByDNSZoneResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecordSetsClientListByDNSZoneResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecordSetsClientListByDNSZonePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecordSetsClientListByDNSZonePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecordSetListResult.NextLink == nil || len(*p.current.RecordSetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDNSZoneHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecordSetsClientListByDNSZoneResponse page.
func (p *RecordSetsClientListByDNSZonePager) PageResponse() RecordSetsClientListByDNSZoneResponse {
	return p.current
}

// RecordSetsClientListByTypePager provides operations for iterating over paged responses.
type RecordSetsClientListByTypePager struct {
	client    *RecordSetsClient
	current   RecordSetsClientListByTypeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecordSetsClientListByTypeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecordSetsClientListByTypePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecordSetsClientListByTypePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecordSetListResult.NextLink == nil || len(*p.current.RecordSetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByTypeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecordSetsClientListByTypeResponse page.
func (p *RecordSetsClientListByTypePager) PageResponse() RecordSetsClientListByTypeResponse {
	return p.current
}

// ZonesClientListByResourceGroupPager provides operations for iterating over paged responses.
type ZonesClientListByResourceGroupPager struct {
	client    *ZonesClient
	current   ZonesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ZonesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ZonesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ZonesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ZoneListResult.NextLink == nil || len(*p.current.ZoneListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ZonesClientListByResourceGroupResponse page.
func (p *ZonesClientListByResourceGroupPager) PageResponse() ZonesClientListByResourceGroupResponse {
	return p.current
}

// ZonesClientListPager provides operations for iterating over paged responses.
type ZonesClientListPager struct {
	client    *ZonesClient
	current   ZonesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ZonesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ZonesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ZonesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ZoneListResult.NextLink == nil || len(*p.current.ZoneListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ZonesClientListResponse page.
func (p *ZonesClientListPager) PageResponse() ZonesClientListResponse {
	return p.current
}
