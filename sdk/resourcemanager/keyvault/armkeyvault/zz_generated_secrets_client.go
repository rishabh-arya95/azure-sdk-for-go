//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

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
	"strconv"
	"strings"
)

// SecretsClient contains the methods for the Secrets group.
// Don't use this type directly, use NewSecretsClient() instead.
type SecretsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecretsClient creates a new instance of SecretsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecretsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecretsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SecretsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create or update a secret in a key vault in the specified subscription. NOTE: This API is intended for
// internal use in ARM deployments. Users should use the data-plane REST service for interaction
// with vault secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - Name of the vault
// secretName - Name of the secret
// parameters - Parameters to create or update the secret
// options - SecretsClientCreateOrUpdateOptions contains the optional parameters for the SecretsClient.CreateOrUpdate method.
func (client *SecretsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters, options *SecretsClientCreateOrUpdateOptions) (SecretsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, secretName, parameters, options)
	if err != nil {
		return SecretsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecretsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SecretsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecretsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters, options *SecretsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SecretsClient) createOrUpdateHandleResponse(resp *http.Response) (SecretsClientCreateOrUpdateResponse, error) {
	result := SecretsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Secret); err != nil {
		return SecretsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the specified secret. NOTE: This API is intended for internal use in ARM deployments. Users should use the data-plane
// REST service for interaction with vault secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - The name of the vault.
// secretName - The name of the secret.
// options - SecretsClientGetOptions contains the optional parameters for the SecretsClient.Get method.
func (client *SecretsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, secretName string, options *SecretsClientGetOptions) (SecretsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, secretName, options)
	if err != nil {
		return SecretsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecretsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecretsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecretsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, options *SecretsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecretsClient) getHandleResponse(resp *http.Response) (SecretsClientGetResponse, error) {
	result := SecretsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Secret); err != nil {
		return SecretsClientGetResponse{}, err
	}
	return result, nil
}

// List - The List operation gets information about the secrets in a vault. NOTE: This API is intended for internal use in
// ARM deployments. Users should use the data-plane REST service for interaction with
// vault secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - The name of the vault.
// options - SecretsClientListOptions contains the optional parameters for the SecretsClient.List method.
func (client *SecretsClient) List(resourceGroupName string, vaultName string, options *SecretsClientListOptions) *SecretsClientListPager {
	return &SecretsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
		},
		advancer: func(ctx context.Context, resp SecretsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecretListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecretsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *SecretsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecretsClient) listHandleResponse(resp *http.Response) (SecretsClientListResponse, error) {
	result := SecretsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretListResult); err != nil {
		return SecretsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update a secret in the specified subscription. NOTE: This API is intended for internal use in ARM deployments.
// Users should use the data-plane REST service for interaction with vault secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - Name of the vault
// secretName - Name of the secret
// parameters - Parameters to patch the secret
// options - SecretsClientUpdateOptions contains the optional parameters for the SecretsClient.Update method.
func (client *SecretsClient) Update(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters, options *SecretsClientUpdateOptions) (SecretsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, secretName, parameters, options)
	if err != nil {
		return SecretsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecretsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SecretsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SecretsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters, options *SecretsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SecretsClient) updateHandleResponse(resp *http.Response) (SecretsClientUpdateResponse, error) {
	result := SecretsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Secret); err != nil {
		return SecretsClientUpdateResponse{}, err
	}
	return result, nil
}
