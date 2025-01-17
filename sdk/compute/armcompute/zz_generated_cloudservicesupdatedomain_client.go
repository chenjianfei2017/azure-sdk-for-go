// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// CloudServicesUpdateDomainClient contains the methods for the CloudServicesUpdateDomain group.
// Don't use this type directly, use NewCloudServicesUpdateDomainClient() instead.
type CloudServicesUpdateDomainClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCloudServicesUpdateDomainClient creates a new instance of CloudServicesUpdateDomainClient with the specified values.
func NewCloudServicesUpdateDomainClient(con *armcore.Connection, subscriptionID string) *CloudServicesUpdateDomainClient {
	return &CloudServicesUpdateDomainClient{con: con, subscriptionID: subscriptionID}
}

// GetUpdateDomain - Gets the specified update domain of a cloud service. Use nextLink property in the response to get the next page of update domains.
// Do this till nextLink is null to fetch all the update domains.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServicesUpdateDomainClient) GetUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainGetUpdateDomainOptions) (UpdateDomainResponse, error) {
	req, err := client.getUpdateDomainCreateRequest(ctx, resourceGroupName, cloudServiceName, updateDomain, options)
	if err != nil {
		return UpdateDomainResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return UpdateDomainResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return UpdateDomainResponse{}, client.getUpdateDomainHandleError(resp)
	}
	return client.getUpdateDomainHandleResponse(resp)
}

// getUpdateDomainCreateRequest creates the GetUpdateDomain request.
func (client *CloudServicesUpdateDomainClient) getUpdateDomainCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainGetUpdateDomainOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{updateDomain}", url.PathEscape(strconv.FormatInt(int64(updateDomain), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUpdateDomainHandleResponse handles the GetUpdateDomain response.
func (client *CloudServicesUpdateDomainClient) getUpdateDomainHandleResponse(resp *azcore.Response) (UpdateDomainResponse, error) {
	var val *UpdateDomain
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return UpdateDomainResponse{}, err
	}
	return UpdateDomainResponse{RawResponse: resp.Response, UpdateDomain: val}, nil
}

// getUpdateDomainHandleError handles the GetUpdateDomain error response.
func (client *CloudServicesUpdateDomainClient) getUpdateDomainHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListUpdateDomains - Gets a list of all update domains in a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServicesUpdateDomainClient) ListUpdateDomains(resourceGroupName string, cloudServiceName string, options *CloudServicesUpdateDomainListUpdateDomainsOptions) UpdateDomainListResultPager {
	return &updateDomainListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listUpdateDomainsCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
		},
		responder: client.listUpdateDomainsHandleResponse,
		errorer:   client.listUpdateDomainsHandleError,
		advancer: func(ctx context.Context, resp UpdateDomainListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.UpdateDomainListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listUpdateDomainsCreateRequest creates the ListUpdateDomains request.
func (client *CloudServicesUpdateDomainClient) listUpdateDomainsCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServicesUpdateDomainListUpdateDomainsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listUpdateDomainsHandleResponse handles the ListUpdateDomains response.
func (client *CloudServicesUpdateDomainClient) listUpdateDomainsHandleResponse(resp *azcore.Response) (UpdateDomainListResultResponse, error) {
	var val *UpdateDomainListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return UpdateDomainListResultResponse{}, err
	}
	return UpdateDomainListResultResponse{RawResponse: resp.Response, UpdateDomainListResult: val}, nil
}

// listUpdateDomainsHandleError handles the ListUpdateDomains error response.
func (client *CloudServicesUpdateDomainClient) listUpdateDomainsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginWalkUpdateDomain - Updates the role instances in the specified update domain.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServicesUpdateDomainClient) BeginWalkUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainBeginWalkUpdateDomainOptions) (HTTPPollerResponse, error) {
	resp, err := client.walkUpdateDomain(ctx, resourceGroupName, cloudServiceName, updateDomain, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CloudServicesUpdateDomainClient.WalkUpdateDomain", "", resp, client.walkUpdateDomainHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeWalkUpdateDomain creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CloudServicesUpdateDomainClient) ResumeWalkUpdateDomain(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("CloudServicesUpdateDomainClient.WalkUpdateDomain", token, client.walkUpdateDomainHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// WalkUpdateDomain - Updates the role instances in the specified update domain.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServicesUpdateDomainClient) walkUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainBeginWalkUpdateDomainOptions) (*azcore.Response, error) {
	req, err := client.walkUpdateDomainCreateRequest(ctx, resourceGroupName, cloudServiceName, updateDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.walkUpdateDomainHandleError(resp)
	}
	return resp, nil
}

// walkUpdateDomainCreateRequest creates the WalkUpdateDomain request.
func (client *CloudServicesUpdateDomainClient) walkUpdateDomainCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainBeginWalkUpdateDomainOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{updateDomain}", url.PathEscape(strconv.FormatInt(int64(updateDomain), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, req.MarshalAsJSON(*options.Parameters)
	}
	return req, nil
}

// walkUpdateDomainHandleError handles the WalkUpdateDomain error response.
func (client *CloudServicesUpdateDomainClient) walkUpdateDomainHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
