// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package indexapi_test

import (
	"github.com/hyperifyio/gocertcenter/internal/api/indexapi"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"testing"
)

func TestIndex(t *testing.T) {
	mockResponse := &mocks.MockResponse{}
	mockRequest := &mocks.MockRequest{IsGet: true}
	mockServer := mocks.NewMockServer()

	indexapi.Index(mockResponse, mockRequest, mockServer)

	if mockResponse.SentStatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", mockResponse.SentStatusCode)
	}
	if data, ok := mockResponse.SentData.(dtos.IndexDTO); !ok || data.Version != "0.0.1" {
		t.Errorf("Expected data version 0.0.1, got %v", mockResponse.SentData)
	}
}
