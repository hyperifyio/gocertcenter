// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apirequests_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apirequests"
)

func TestRequestImpl_IsMethodGet(t *testing.T) {

	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{"GET Method", http.MethodGet, true},
		{"POST Method", http.MethodPost, false},
		{"PUT Method", http.MethodPut, false},
		{"DELETE Method", http.MethodDelete, false},
		{"PATCH Method", http.MethodPatch, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, "/", nil)
			if err != nil {
				t.Fatal("Creating request failed:", err)
			}

			requestImpl := apirequests.NewRequest(req)
			if got := requestImpl.IsMethodGet(); got != test.expected {
				t.Errorf("IsMethodGet() = %v, want %v for method %s", got, test.expected, test.method)
			}
		})
	}

}

func TestRequestImpl_GetURL(t *testing.T) {
	expectedURL := "http://example.com/test"
	req, err := http.NewRequest(http.MethodGet, expectedURL, nil)
	if err != nil {
		t.Fatal("Creating request failed:", err)
	}

	requestImpl := apirequests.NewRequest(req)
	if gotURL := requestImpl.GetURL(); gotURL.String() != expectedURL {
		t.Errorf("GetURL() = %v, want %v", gotURL, expectedURL)
	}
}

func TestRequestImpl_GetMethod(t *testing.T) {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req, err := http.NewRequest(method, "/", nil)
			if err != nil {
				t.Fatal("Creating request failed:", err)
			}

			requestImpl := apirequests.NewRequest(req)
			if gotMethod := requestImpl.GetMethod(); gotMethod != method {
				t.Errorf("GetMethod() = %v, want %v", gotMethod, method)
			}
		})
	}
}

func TestRequestImpl_GetVariable(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/test/{var}", func(w http.ResponseWriter, r *http.Request) {
		requestImpl := apirequests.NewRequest(r)
		value := requestImpl.GetVariable("var")
		if value != "value" {
			t.Errorf("GetVars() did not return the expected value, got %v", value)
		}
	})

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	url := fmt.Sprintf("%s/test/value", testServer.URL)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal("Making GET request failed:", err)
	}
	defer resp.Body.Close()
}

func TestRequestImpl_GetHeader(t *testing.T) {
	headerName := "Content-Type"
	headerValue := "application/json"
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(headerName, headerValue)

	requestImpl := apirequests.NewRequest(req)
	gotHeaderValue := requestImpl.GetHeader(headerName)
	if gotHeaderValue != headerValue {
		t.Errorf("GetHeader() got = %v, want %v", gotHeaderValue, headerValue)
	}
}

func TestRequestImpl_BodyAndGetBodyBytes(t *testing.T) {
	bodyContent := "test body content"
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(bodyContent))

	requestImpl := apirequests.NewRequest(req)

	// Testing GetBodyBytes which also verifies Body() indirectly
	bodyBytes, err := requestImpl.GetBodyBytes()
	if err != nil {
		t.Fatalf("GetBodyBytes() error = %v", err)
	}
	if gotBodyContent := string(bodyBytes); gotBodyContent != bodyContent {
		t.Errorf("GetBodyBytes() got = %v, want %v", gotBodyContent, bodyContent)
	}
}

func TestRequestImpl_GetQueryParam(t *testing.T) {
	queryParamName := "param"
	queryParamValue := "value"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/?%s=%s", queryParamName, queryParamValue), nil)

	requestImpl := apirequests.NewRequest(req)
	gotQueryParamValue := requestImpl.GetQueryParam(queryParamName)
	if gotQueryParamValue != queryParamValue {
		t.Errorf("GetQueryParam() got = %v, want %v", gotQueryParamValue, queryParamValue)
	}
}

func TestRequestImpl_Body(t *testing.T) {
	bodyContent := "test body content"
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(bodyContent))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	requestImpl := apirequests.NewRequest(req)

	body := requestImpl.Body()
	defer body.Close() // Ensure the body is closed after reading

	readContent, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	if string(readContent) != bodyContent {
		t.Errorf("Body content mismatch, got: %s, want: %s", string(readContent), bodyContent)
	}
}

func TestRequestImpl_GetBodyBytes_ReadError(t *testing.T) {
	// Create a new http.Request with the body set to our error-producing reader
	req, err := http.NewRequest(http.MethodGet, "/", errorReader{})
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	requestImpl := apirequests.NewRequest(req)

	// Attempt to get body bytes, expecting an error
	_, err = requestImpl.GetBodyBytes()
	if err == nil {
		t.Error("Expected an error from GetBodyBytes, got nil")
	} else if !strings.Contains(err.Error(), "GetBodyBytes: failed") {
		t.Errorf("Error message does not match expected, got: %v", err)
	}
}

type errorReader struct{}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func (e errorReader) Close() error {
	return nil
}
