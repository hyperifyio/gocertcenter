// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package indexapi_test

import (
	"github.com/hyperifyio/gocertcenter/internal/api/indexapi"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"reflect"
	"testing"
)

const ExpectedStatusCode = 200
const ExpectedVersion = "0.0.1"
const ExpectedContentType = "application/json"
const ExpectedSummaryText = "Returns information about the running server"
const ExpectedDescriptionText = "This includes the software name and a version"

func TestIndex(t *testing.T) {
	mockResponse := &mocks.MockResponse{}
	mockRequest := &mocks.MockRequest{IsGet: true}
	mockServer := mocks.NewMockServer()

	indexapi.Index(mockResponse, mockRequest, mockServer)

	if mockResponse.SentStatusCode != ExpectedStatusCode {
		t.Errorf("Expected status code %d, got %d", ExpectedStatusCode, mockResponse.SentStatusCode)
	}
	if data, ok := mockResponse.SentData.(dtos.IndexDTO); !ok || data.Version != ExpectedVersion {
		t.Errorf("Expected data version %s, got %v", ExpectedVersion, mockResponse.SentData)
	}
}

func TestIndexDefinitions(t *testing.T) {

	defs := indexapi.IndexDefinitions()

	// Check the summary and description
	if defs.Summary != ExpectedSummaryText {
		t.Errorf("Expected summary to be '%s', got '%s'", ExpectedSummaryText, defs.Summary)
	}

	if defs.Description != ExpectedDescriptionText {
		t.Errorf("Expected description to be '%s', got '%s'", ExpectedDescriptionText, defs.Description)
	}

	// Check the response for HTTP 200 status code
	resp200, ok := defs.Responses[ExpectedStatusCode]
	if !ok {
		t.Fatalf("Expected a %d response definition", ExpectedStatusCode)
	}

	contentType, ok := resp200.Content[ExpectedContentType]
	if !ok {
		t.Fatalf("Expected '%s' content type in %d response", ExpectedContentType, ExpectedStatusCode)
	}

	// Since we cannot directly compare complex structs in the test, use reflection or type assertion
	// Here we use reflection to check if the type of contentType.Value is IndexDTO
	expectedDTOType := reflect.TypeOf(dtos.IndexDTO{})
	contentValueType := reflect.TypeOf(contentType.Value)
	if contentValueType != expectedDTOType {
		t.Errorf("Expected response content value type to be '%v', got '%v'", expectedDTOType, contentValueType)
	}

}
