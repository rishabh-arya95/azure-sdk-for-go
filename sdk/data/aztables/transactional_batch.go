// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"sort"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	generated "github.com/Azure/azure-sdk-for-go/sdk/data/aztables/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
)

// TransactionType is the type for a specific transaction operation.
type TransactionType string

const (
	TransactionTypeAdd           TransactionType = "add"
	TransactionTypeUpdateMerge   TransactionType = "updatemerge"
	TransactionTypeUpdateReplace TransactionType = "updatereplace"
	TransactionTypeDelete        TransactionType = "delete"
	TransactionTypeInsertMerge   TransactionType = "insertmerge"
	TransactionTypeInsertReplace TransactionType = "insertreplace"
)

// PossibleTransactionTypeValues returns the possible values for the TransactionType const type.
func PossibleTransactionTypeValues() []TransactionType {
	return []TransactionType{
		TransactionTypeAdd,
		TransactionTypeUpdateMerge,
		TransactionTypeUpdateReplace,
		TransactionTypeDelete,
		TransactionTypeInsertMerge,
		TransactionTypeInsertReplace,
	}
}

type TransactionAction struct {
	ActionType TransactionType
	Entity     []byte
	IfMatch    *azcore.ETag
}

type TransactionResponse struct {
	// The response for a single table.
	TransactionResponses []http.Response
}

type SubmitTransactionOptions struct {
	RequestID *string
}

// SubmitTransaction submits the table transactional batch according to the slice of TransactionActions provided. All transactionActions must be for entities
// with the same PartitionKey. There can only be one transaction action for a row key, a duplicated row key will return an error. The TransactionResponse object
// contains the response for each sub-request in the same order that they are made in the transactionActions parameter.
func (t *Client) SubmitTransaction(ctx context.Context, transactionActions []TransactionAction, tableSubmitTransactionOptions *SubmitTransactionOptions) (TransactionResponse, error) {
	u1, err := uuid.New()
	if err != nil {
		return TransactionResponse{}, err
	}
	u2, err := uuid.New()
	if err != nil {
		return TransactionResponse{}, err
	}
	return t.submitTransactionInternal(ctx, &transactionActions, u1, u2, tableSubmitTransactionOptions)
}

// submitTransactionInternal is the internal implementation for SubmitTransaction. It allows for explicit configuration of the batch and changeset UUID values for testing.
func (t *Client) submitTransactionInternal(ctx context.Context, transactionActions *[]TransactionAction, batchUuid uuid.UUID, changesetUuid uuid.UUID, tableSubmitTransactionOptions *SubmitTransactionOptions) (TransactionResponse, error) {
	if len(*transactionActions) == 0 {
		return TransactionResponse{}, errEmptyTransaction
	}
	changesetBoundary := fmt.Sprintf("changeset_%s", changesetUuid.String())
	changeSetBody, err := t.generateChangesetBody(changesetBoundary, transactionActions)
	if err != nil {
		return TransactionResponse{}, err
	}
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(t.con.Endpoint(), "$batch"))
	if err != nil {
		return TransactionResponse{}, err
	}
	req.Raw().Header.Set("x-ms-version", "2019-02-02")
	if tableSubmitTransactionOptions != nil && tableSubmitTransactionOptions.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *tableSubmitTransactionOptions.RequestID)
	}
	req.Raw().Header.Set("DataServiceVersion", "3.0")
	req.Raw().Header.Set("Accept", string(generated.ODataMetadataFormatApplicationJSONODataMinimalmetadata))

	boundary := fmt.Sprintf("batch_%s", batchUuid.String())
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err = writer.SetBoundary(boundary)
	if err != nil {
		return TransactionResponse{}, err
	}
	h := make(textproto.MIMEHeader)
	h.Set(headerContentType, fmt.Sprintf("multipart/mixed; boundary=%s", changesetBoundary))
	batchWriter, err := writer.CreatePart(h)
	if err != nil {
		return TransactionResponse{}, err
	}
	_, err = batchWriter.Write(changeSetBody.Bytes())
	if err != nil {
		return TransactionResponse{}, err
	}
	writer.Close()

	err = req.SetBody(streaming.NopCloser(bytes.NewReader(body.Bytes())), fmt.Sprintf("multipart/mixed; boundary=%s", boundary))
	if err != nil {
		return TransactionResponse{}, err
	}

	resp, err := t.con.Pipeline().Do(req)
	if err != nil {
		return TransactionResponse{}, err
	}

	transactionResponse, err := buildTransactionResponse(req, resp, len(*transactionActions))
	if err != nil {
		return *transactionResponse, err
	}

	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return TransactionResponse{}, runtime.NewResponseError(resp)
	}
	return *transactionResponse, nil
}

// create the transaction response. This will read the inner responses
func buildTransactionResponse(req *policy.Request, resp *http.Response, itemCount int) (*TransactionResponse, error) {
	innerResponses := make([]http.Response, itemCount)
	result := TransactionResponse{TransactionResponses: innerResponses}

	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &TransactionResponse{}, err
	}
	reader := bytes.NewReader(bytesBody)
	if bytes.IndexByte(bytesBody, '{') == 0 {
		// This is a failure and the body is json
		return &TransactionResponse{}, runtime.NewResponseError(resp)
	}

	outerBoundary := getBoundaryName(bytesBody)
	mpReader := multipart.NewReader(reader, outerBoundary)
	outerPart, err := mpReader.NextPart()
	if err != nil {
		return &TransactionResponse{}, err
	}

	innerBytes, err := ioutil.ReadAll(outerPart)
	if err != nil && err != io.ErrUnexpectedEOF { // Cosmos specific error handling
		return &TransactionResponse{}, err
	}
	innerBoundary := getBoundaryName(innerBytes)
	reader = bytes.NewReader(innerBytes)
	mpReader = multipart.NewReader(reader, innerBoundary)
	i := 0
	innerPart, err := mpReader.NextPart()
	for ; err == nil; innerPart, err = mpReader.NextPart() {
		part, err := ioutil.ReadAll(innerPart)
		if err != nil {
			break
		}
		r, err := http.ReadResponse(bufio.NewReader(bytes.NewBuffer(part)), req.Raw())
		if err != nil {
			return &TransactionResponse{}, err
		}
		if r.StatusCode >= 400 {
			if err != nil {
				return &TransactionResponse{}, err
			} else {
				return &result, runtime.NewResponseError(resp)
			}
		}
		innerResponses[i] = *r
		i++
	}

	return &result, nil
}

func getBoundaryName(bytesBody []byte) string {
	end := bytes.Index(bytesBody, []byte("\n"))
	if end > 0 && bytesBody[end-1] == '\r' {
		end -= 1
	}
	return string(bytesBody[2:end])
}

// generateChangesetBody generates the individual changesets for the various operations within the batch request.
// There is a changeset for Insert, Delete, Merge etc.
func (t *Client) generateChangesetBody(changesetBoundary string, transactionActions *[]TransactionAction) (*bytes.Buffer, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.SetBoundary(changesetBoundary)
	if err != nil {
		return nil, err
	}

	for _, be := range *transactionActions {
		err := t.generateEntitySubset(&be, writer)
		if err != nil {
			return nil, err
		}
	}

	writer.Close()
	return body, nil
}

// generateEntitySubset generates body payload for particular batch entity
func (t *Client) generateEntitySubset(transactionAction *TransactionAction, writer *multipart.Writer) error {
	h := make(textproto.MIMEHeader)
	h.Set(headerContentTransferEncoding, "binary")
	h.Set(headerContentType, "application/http")
	qo := &generated.QueryOptions{Format: generated.ODataMetadataFormatApplicationJSONODataMinimalmetadata.ToPtr()}

	operationWriter, err := writer.CreatePart(h)
	if err != nil {
		return err
	}
	var req *policy.Request
	var entity map[string]interface{}
	err = json.Unmarshal(transactionAction.Entity, &entity)
	if err != nil {
		return err
	}

	if _, ok := entity[partitionKey]; !ok {
		return fmt.Errorf("entity properties must contain a %s property", partitionKey)
	}
	if _, ok := entity[rowKey]; !ok {
		return fmt.Errorf("entity properties must contain a %s property", rowKey)
	}

	switch transactionAction.ActionType {
	case TransactionTypeDelete:
		ifMatch := string(azcore.ETagAny)
		if transactionAction.IfMatch != nil {
			ifMatch = string(*transactionAction.IfMatch)
		}
		req, err = t.client.DeleteEntityCreateRequest(
			ctx,
			generated.Enum1Three0,
			t.name,
			entity[partitionKey].(string),
			entity[rowKey].(string),
			ifMatch,
			&generated.TableClientDeleteEntityOptions{},
			qo,
		)
		if err != nil {
			return err
		}
	case TransactionTypeAdd:
		req, err = t.client.InsertEntityCreateRequest(
			ctx,
			generated.Enum1Three0,
			t.name,
			&generated.TableClientInsertEntityOptions{
				TableEntityProperties: entity,
				ResponsePreference:    generated.ResponseFormatReturnNoContent.ToPtr(),
			},
			qo,
		)
		if err != nil {
			return err
		}
	case TransactionTypeUpdateMerge:
		fallthrough
	case TransactionTypeInsertMerge:
		opts := &generated.TableClientMergeEntityOptions{TableEntityProperties: entity}
		if transactionAction.IfMatch != nil {
			opts.IfMatch = to.StringPtr(string(*transactionAction.IfMatch))
		}
		req, err = t.client.MergeEntityCreateRequest(
			ctx,
			generated.Enum1Three0,
			t.name,
			entity[partitionKey].(string),
			entity[rowKey].(string),
			opts,
			&generated.QueryOptions{},
		)
		if err != nil {
			return err
		}
		if isCosmosEndpoint(t.con.Endpoint()) {
			transformPatchToCosmosPost(req)
		}
	case TransactionTypeUpdateReplace:
		fallthrough
	case TransactionTypeInsertReplace:
		opts := &generated.TableClientUpdateEntityOptions{TableEntityProperties: entity}
		if transactionAction.IfMatch != nil {
			opts.IfMatch = to.StringPtr(string(*transactionAction.IfMatch))
		}
		req, err = t.client.UpdateEntityCreateRequest(
			ctx,
			generated.Enum1Three0,
			t.name,
			entity[partitionKey].(string),
			entity[rowKey].(string),
			opts,
			&generated.QueryOptions{},
		)
		if err != nil {
			return err
		}
	}

	urlAndVerb := fmt.Sprintf("%s %s HTTP/1.1\r\n", req.Raw().Method, req.Raw().URL)
	_, err = operationWriter.Write([]byte(urlAndVerb))
	if err != nil {
		return err
	}
	err = writeHeaders(req.Raw().Header, &operationWriter)
	if err != nil {
		return err
	}
	_, err = operationWriter.Write([]byte("\r\n")) // additional \r\n is needed per changeset separating the "headers" and the body.
	if err != nil {
		return err
	}
	if req.Raw().Body != nil {
		_, err = io.Copy(operationWriter, req.Body())
	}

	return err
}

func writeHeaders(h http.Header, writer *io.Writer) error {
	// This way it is guaranteed the headers will be written in a sorted order
	var keys []string
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var err error
	for _, k := range keys {
		_, err = (*writer).Write([]byte(fmt.Sprintf("%s: %s\r\n", k, h.Get(k))))

	}
	return err
}
