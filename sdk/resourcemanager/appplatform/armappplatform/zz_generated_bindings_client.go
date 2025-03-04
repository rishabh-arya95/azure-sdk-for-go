//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BindingsClient contains the methods for the Bindings group.
// Don't use this type directly, use NewBindingsClient() instead.
type BindingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBindingsClient creates a new instance of BindingsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBindingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BindingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &BindingsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create a new Binding or update an exiting Binding.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// appName - The name of the App resource.
// bindingName - The name of the Binding resource.
// bindingResource - Parameters for the create or update operation
// options - BindingsClientBeginCreateOrUpdateOptions contains the optional parameters for the BindingsClient.BeginCreateOrUpdate
// method.
func (client *BindingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginCreateOrUpdateOptions) (BindingsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, appName, bindingName, bindingResource, options)
	if err != nil {
		return BindingsClientCreateOrUpdatePollerResponse{}, err
	}
	result := BindingsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BindingsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return BindingsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BindingsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create a new Binding or update an exiting Binding.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BindingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, appName, bindingName, bindingResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BindingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if bindingName == "" {
		return nil, errors.New("parameter bindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bindingName}", url.PathEscape(bindingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, bindingResource)
}

// BeginDelete - Operation to delete a Binding.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// appName - The name of the App resource.
// bindingName - The name of the Binding resource.
// options - BindingsClientBeginDeleteOptions contains the optional parameters for the BindingsClient.BeginDelete method.
func (client *BindingsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, options *BindingsClientBeginDeleteOptions) (BindingsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, appName, bindingName, options)
	if err != nil {
		return BindingsClientDeletePollerResponse{}, err
	}
	result := BindingsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BindingsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return BindingsClientDeletePollerResponse{}, err
	}
	result.Poller = &BindingsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a Binding.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BindingsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, options *BindingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, appName, bindingName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BindingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, options *BindingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if bindingName == "" {
		return nil, errors.New("parameter bindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bindingName}", url.PathEscape(bindingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a Binding and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// appName - The name of the App resource.
// bindingName - The name of the Binding resource.
// options - BindingsClientGetOptions contains the optional parameters for the BindingsClient.Get method.
func (client *BindingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, options *BindingsClientGetOptions) (BindingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, appName, bindingName, options)
	if err != nil {
		return BindingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BindingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BindingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BindingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, options *BindingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if bindingName == "" {
		return nil, errors.New("parameter bindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bindingName}", url.PathEscape(bindingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BindingsClient) getHandleResponse(resp *http.Response) (BindingsClientGetResponse, error) {
	result := BindingsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BindingResource); err != nil {
		return BindingsClientGetResponse{}, err
	}
	return result, nil
}

// List - Handles requests to list all resources in an App.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// appName - The name of the App resource.
// options - BindingsClientListOptions contains the optional parameters for the BindingsClient.List method.
func (client *BindingsClient) List(resourceGroupName string, serviceName string, appName string, options *BindingsClientListOptions) *BindingsClientListPager {
	return &BindingsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, appName, options)
		},
		advancer: func(ctx context.Context, resp BindingsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BindingResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BindingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *BindingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BindingsClient) listHandleResponse(resp *http.Response) (BindingsClientListResponse, error) {
	result := BindingsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BindingResourceCollection); err != nil {
		return BindingsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update an exiting Binding.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// appName - The name of the App resource.
// bindingName - The name of the Binding resource.
// bindingResource - Parameters for the update operation
// options - BindingsClientBeginUpdateOptions contains the optional parameters for the BindingsClient.BeginUpdate method.
func (client *BindingsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginUpdateOptions) (BindingsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serviceName, appName, bindingName, bindingResource, options)
	if err != nil {
		return BindingsClientUpdatePollerResponse{}, err
	}
	result := BindingsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BindingsClient.Update", "azure-async-operation", resp, client.pl)
	if err != nil {
		return BindingsClientUpdatePollerResponse{}, err
	}
	result.Poller = &BindingsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update an exiting Binding.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BindingsClient) update(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, appName, bindingName, bindingResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *BindingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource BindingResource, options *BindingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if bindingName == "" {
		return nil, errors.New("parameter bindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bindingName}", url.PathEscape(bindingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, bindingResource)
}
