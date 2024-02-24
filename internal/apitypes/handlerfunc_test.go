// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apitypes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/apitypes"
)

func TestToGorillaHandlerFunc(t *testing.T) {
	// Define a flag to check if handler is called
	handlerCalled := false

	// Define an http.HandlerFunc that sets the flag to true when called
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK) // Optionally set a response status to verify later
	})

	// Convert the http.HandlerFunc to a gorilla.HandlerFunc
	convertedHandler := apitypes.ToGorillaHandlerFunc(originalHandler)

	// Create a new HTTP request that will be handled by the converted handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Creating request failed:", err)
	}

	// Create a ResponseRecorder to capture the handler's HTTP response
	rr := httptest.NewRecorder()

	// Invoke the converted handler with the request and ResponseRecorder
	convertedHandler(rr, req)

	// Verify the handler was called by checking the flag
	if !handlerCalled {
		t.Error("Expected the original http.HandlerFunc to be called, but it wasn't")
	}

	// Optionally, verify the response status code if you set one in your handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
