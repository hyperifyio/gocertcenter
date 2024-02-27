// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors_test

import (
	"net/url"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apierrors"
)

func TestMethodNotAllowed(t *testing.T) {
	requestURL, err := url.Parse("https://example.com/nonexistent")
	if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
	}

	request := &apimocks.MockRequest{
		Method: "POST",
		URL:    requestURL,
	}

	response := &apimocks.MockResponse{}

	// Assuming apierrors.MethodNotAllowed expects IResponse, IRequest, IServer which can be nil for this test
	apierrors.MethodNotAllowed(response, request) // Mock server can be nil if not used in the handler

	// Verify the response
	if !response.MethodNotSupportedError {
		t.Error("Expected MethodNotSupportedError to be true")
	}

}
