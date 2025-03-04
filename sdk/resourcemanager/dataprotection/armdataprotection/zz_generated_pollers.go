//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// BackupInstancesClientAdhocBackupPoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientAdhocBackupPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientAdhocBackupPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientAdhocBackupPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientAdhocBackupResponse will be returned.
func (p *BackupInstancesClientAdhocBackupPoller) FinalResponse(ctx context.Context) (BackupInstancesClientAdhocBackupResponse, error) {
	respType := BackupInstancesClientAdhocBackupResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationJobExtendedInfo)
	if err != nil {
		return BackupInstancesClientAdhocBackupResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientAdhocBackupPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientCreateOrUpdateResponse will be returned.
func (p *BackupInstancesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (BackupInstancesClientCreateOrUpdateResponse, error) {
	respType := BackupInstancesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.BackupInstanceResource)
	if err != nil {
		return BackupInstancesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientDeleteResponse will be returned.
func (p *BackupInstancesClientDeletePoller) FinalResponse(ctx context.Context) (BackupInstancesClientDeleteResponse, error) {
	respType := BackupInstancesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return BackupInstancesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientTriggerRehydratePoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientTriggerRehydratePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientTriggerRehydratePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientTriggerRehydratePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientTriggerRehydrateResponse will be returned.
func (p *BackupInstancesClientTriggerRehydratePoller) FinalResponse(ctx context.Context) (BackupInstancesClientTriggerRehydrateResponse, error) {
	respType := BackupInstancesClientTriggerRehydrateResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return BackupInstancesClientTriggerRehydrateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientTriggerRehydratePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientTriggerRestorePoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientTriggerRestorePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientTriggerRestorePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientTriggerRestorePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientTriggerRestoreResponse will be returned.
func (p *BackupInstancesClientTriggerRestorePoller) FinalResponse(ctx context.Context) (BackupInstancesClientTriggerRestoreResponse, error) {
	respType := BackupInstancesClientTriggerRestoreResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationJobExtendedInfo)
	if err != nil {
		return BackupInstancesClientTriggerRestoreResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientTriggerRestorePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientValidateForBackupPoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientValidateForBackupPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientValidateForBackupPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientValidateForBackupPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientValidateForBackupResponse will be returned.
func (p *BackupInstancesClientValidateForBackupPoller) FinalResponse(ctx context.Context) (BackupInstancesClientValidateForBackupResponse, error) {
	respType := BackupInstancesClientValidateForBackupResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationJobExtendedInfo)
	if err != nil {
		return BackupInstancesClientValidateForBackupResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientValidateForBackupPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupInstancesClientValidateForRestorePoller provides polling facilities until the operation reaches a terminal state.
type BackupInstancesClientValidateForRestorePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupInstancesClientValidateForRestorePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupInstancesClientValidateForRestorePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupInstancesClientValidateForRestoreResponse will be returned.
func (p *BackupInstancesClientValidateForRestorePoller) FinalResponse(ctx context.Context) (BackupInstancesClientValidateForRestoreResponse, error) {
	respType := BackupInstancesClientValidateForRestoreResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationJobExtendedInfo)
	if err != nil {
		return BackupInstancesClientValidateForRestoreResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupInstancesClientValidateForRestorePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupVaultsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type BackupVaultsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupVaultsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupVaultsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupVaultsClientCreateOrUpdateResponse will be returned.
func (p *BackupVaultsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (BackupVaultsClientCreateOrUpdateResponse, error) {
	respType := BackupVaultsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.BackupVaultResource)
	if err != nil {
		return BackupVaultsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupVaultsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// BackupVaultsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type BackupVaultsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *BackupVaultsClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *BackupVaultsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final BackupVaultsClientUpdateResponse will be returned.
func (p *BackupVaultsClientUpdatePoller) FinalResponse(ctx context.Context) (BackupVaultsClientUpdateResponse, error) {
	respType := BackupVaultsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.BackupVaultResource)
	if err != nil {
		return BackupVaultsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *BackupVaultsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ExportJobsClientTriggerPoller provides polling facilities until the operation reaches a terminal state.
type ExportJobsClientTriggerPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ExportJobsClientTriggerPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ExportJobsClientTriggerPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ExportJobsClientTriggerResponse will be returned.
func (p *ExportJobsClientTriggerPoller) FinalResponse(ctx context.Context) (ExportJobsClientTriggerResponse, error) {
	respType := ExportJobsClientTriggerResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ExportJobsClientTriggerResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ExportJobsClientTriggerPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
