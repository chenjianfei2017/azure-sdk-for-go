package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// PartitionListsClient is the client for the PartitionLists methods of the Servicefabric service.
type PartitionListsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewPartitionListsClient creates an instance of the PartitionListsClient client.
func NewPartitionListsClient(timeout *int32) PartitionListsClient {
	return NewPartitionListsClientWithBaseURI(DefaultBaseURI, timeout)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewPartitionListsClientWithBaseURI creates an instance of the PartitionListsClient client.
func NewPartitionListsClientWithBaseURI(baseURI string, timeout *int32) PartitionListsClient {
	return PartitionListsClient{NewWithBaseURI(baseURI, timeout)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Repair repair partition lists
// Parameters:
// serviceName - the name of the service
func (client PartitionListsClient) Repair(ctx context.Context, serviceName string) (result String, err error) {
	req, err := client.RepairPreparer(ctx, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionListsClient", "Repair", nil, "Failure preparing request")
		return
	}

	resp, err := client.RepairSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionListsClient", "Repair", resp, "Failure sending request")
		return
	}

	result, err = client.RepairResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionListsClient", "Repair", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RepairPreparer prepares the Repair request.
func (client PartitionListsClient) RepairPreparer(ctx context.Context, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": serviceName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Services/{serviceName}/$/GetPartitions/$/Recover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RepairSender sends the Repair request. The method will close the
// http.Response Body if it receives an error.
func (client PartitionListsClient) RepairSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RepairResponder handles the response to the Repair request. The method always
// closes the http.Response Body.
func (client PartitionListsClient) RepairResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}