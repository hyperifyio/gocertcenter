// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apiserver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apiserver"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func TestResponseHandler_Error(t *testing.T) {
	// Mock request handler that always returns an error
	mockHandler := func(response apitypes.Response, request apitypes.Request) error {
		return errors.New("test error")
	}

	// Wrap the mock handler with the responseHandler
	handler := apiserver.ResponseHandler(mockHandler)

	// Create a test HTTP request and response recorder
	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// Call the wrapped handler
	handler.ServeHTTP(w, req)

	// Verify the response
	resp := w.Result()
	defer resp.Body.Close()

	// Check the status code is 500 Internal ApplicationServer Error
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Expected HTTP status code 500")

	// Optionally, you can also verify the response body or the log output if necessary
	// This might require additional setup to capture log output or parse the response body
}
