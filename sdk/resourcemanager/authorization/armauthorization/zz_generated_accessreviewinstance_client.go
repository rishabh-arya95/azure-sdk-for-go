//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstanceClient contains the methods for the AccessReviewInstance group.
// Don't use this type directly, use NewAccessReviewInstanceClient() instead.
type AccessReviewInstanceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessReviewInstanceClient creates a new instance of AccessReviewInstanceClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewInstanceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccessReviewInstanceClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AccessReviewInstanceClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// AcceptRecommendations - An action to accept recommendations for decision in an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceClientAcceptRecommendationsOptions contains the optional parameters for the AccessReviewInstanceClient.AcceptRecommendations
// method.
func (client *AccessReviewInstanceClient) AcceptRecommendations(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientAcceptRecommendationsOptions) (AccessReviewInstanceClientAcceptRecommendationsResponse, error) {
	req, err := client.acceptRecommendationsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewInstanceClientAcceptRecommendationsResponse{RawResponse: resp}, nil
}

// acceptRecommendationsCreateRequest creates the AcceptRecommendations request.
func (client *AccessReviewInstanceClient) acceptRecommendationsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientAcceptRecommendationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/acceptRecommendations"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ApplyDecisions - An action to apply all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceClientApplyDecisionsOptions contains the optional parameters for the AccessReviewInstanceClient.ApplyDecisions
// method.
func (client *AccessReviewInstanceClient) ApplyDecisions(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientApplyDecisionsOptions) (AccessReviewInstanceClientApplyDecisionsResponse, error) {
	req, err := client.applyDecisionsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccessReviewInstanceClientApplyDecisionsResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewInstanceClientApplyDecisionsResponse{RawResponse: resp}, nil
}

// applyDecisionsCreateRequest creates the ApplyDecisions request.
func (client *AccessReviewInstanceClient) applyDecisionsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientApplyDecisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/applyDecisions"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ResetDecisions - An action to reset all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceClientResetDecisionsOptions contains the optional parameters for the AccessReviewInstanceClient.ResetDecisions
// method.
func (client *AccessReviewInstanceClient) ResetDecisions(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientResetDecisionsOptions) (AccessReviewInstanceClientResetDecisionsResponse, error) {
	req, err := client.resetDecisionsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccessReviewInstanceClientResetDecisionsResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewInstanceClientResetDecisionsResponse{RawResponse: resp}, nil
}

// resetDecisionsCreateRequest creates the ResetDecisions request.
func (client *AccessReviewInstanceClient) resetDecisionsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientResetDecisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/resetDecisions"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// SendReminders - An action to send reminders for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceClientSendRemindersOptions contains the optional parameters for the AccessReviewInstanceClient.SendReminders
// method.
func (client *AccessReviewInstanceClient) SendReminders(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientSendRemindersOptions) (AccessReviewInstanceClientSendRemindersResponse, error) {
	req, err := client.sendRemindersCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientSendRemindersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceClientSendRemindersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccessReviewInstanceClientSendRemindersResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewInstanceClientSendRemindersResponse{RawResponse: resp}, nil
}

// sendRemindersCreateRequest creates the SendReminders request.
func (client *AccessReviewInstanceClient) sendRemindersCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientSendRemindersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/sendReminders"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Stop - An action to stop an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceClientStopOptions contains the optional parameters for the AccessReviewInstanceClient.Stop
// method.
func (client *AccessReviewInstanceClient) Stop(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientStopOptions) (AccessReviewInstanceClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientStopResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccessReviewInstanceClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewInstanceClientStopResponse{RawResponse: resp}, nil
}

// stopCreateRequest creates the Stop request.
func (client *AccessReviewInstanceClient) stopCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/stop"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
