//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

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

// DataServicesClient contains the methods for the DataServices group.
// Don't use this type directly, use NewDataServicesClient() instead.
type DataServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataServicesClient creates a new instance of DataServicesClient with the specified values.
// subscriptionID - The Subscription Id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DataServicesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets the data service that matches the data service name given.
// If the operation fails it returns an *azcore.ResponseError type.
// dataServiceName - The name of the data service that is being queried.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - DataServicesClientGetOptions contains the optional parameters for the DataServicesClient.Get method.
func (client *DataServicesClient) Get(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *DataServicesClientGetOptions) (DataServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, dataServiceName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return DataServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataServicesClient) getCreateRequest(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *DataServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataServicesClient) getHandleResponse(resp *http.Response) (DataServicesClientGetResponse, error) {
	result := DataServicesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataService); err != nil {
		return DataServicesClientGetResponse{}, err
	}
	return result, nil
}

// ListByDataManager - This method gets all the data services.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - DataServicesClientListByDataManagerOptions contains the optional parameters for the DataServicesClient.ListByDataManager
// method.
func (client *DataServicesClient) ListByDataManager(resourceGroupName string, dataManagerName string, options *DataServicesClientListByDataManagerOptions) *DataServicesClientListByDataManagerPager {
	return &DataServicesClientListByDataManagerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
		},
		advancer: func(ctx context.Context, resp DataServicesClientListByDataManagerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DataServiceList.NextLink)
		},
	}
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *DataServicesClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *DataServicesClientListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *DataServicesClient) listByDataManagerHandleResponse(resp *http.Response) (DataServicesClientListByDataManagerResponse, error) {
	result := DataServicesClientListByDataManagerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataServiceList); err != nil {
		return DataServicesClientListByDataManagerResponse{}, err
	}
	return result, nil
}
