// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors_test

import (
	"github.com/hyperifyio/gocertcenter/internal/apierrors"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"net/url"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	requestURL, err := url.Parse("https://example.com/nonexistent")
	if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
	}

	request := &mocks.MockRequest{
		Method: "POST",
		URL:    requestURL,
	}

	response := &mocks.MockResponse{}

	// Assuming apierrors.MethodNotAllowed expects IResponse, IRequest, IServer which can be nil for this test
	apierrors.MethodNotAllowed(response, request, nil) // Mock server can be nil if not used in the handler

	// Verify the response
	if !response.MethodNotSupportedError {
		t.Error("Expected MethodNotSupportedError to be true")
	}

}
