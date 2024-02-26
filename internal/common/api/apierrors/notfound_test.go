// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors_test

import (
	"net/url"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apierrors"
)

func TestNotFound(t *testing.T) {
	requestURL, err := url.Parse("https://example.com/nonexistent")
	if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
	}

	request := &apimocks.MockRequest{
		Method: "GET",
		URL:    requestURL,
	}

	response := &apimocks.MockResponse{}

	// Call the NotFound function with the mocks
	apierrors.NotFound(response, request, nil) // Assuming server is not used in the handler and can be nil

	// Verify the response
	if !response.NotFoundError {
		t.Error("Expected NotFoundError to be true, indicating that SendNotFoundError was called")
	}

	// Optionally, verify that no other error types were mistakenly set
	if response.MethodNotSupportedError || response.SentStatusCode != 0 || response.SentErrorMessage != "" {
		t.Error("Unexpected response behavior; only NotFoundError should be set")
	}

	// Since NotFound does not directly use Send to set status codes or send data,
	// there's no need to check SentData or SentStatusCode unless your implementation of SendNotFoundError does so.
}
