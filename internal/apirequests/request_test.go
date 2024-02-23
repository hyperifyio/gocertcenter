// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apirequests_test

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hyperifyio/gocertcenter/internal/apirequests"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestRequestImpl_GetVars(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/test/{var}", func(w http.ResponseWriter, r *http.Request) {
		requestImpl := apirequests.NewRequest(r)
		vars := requestImpl.GetVars()
		if val, ok := vars["var"]; !ok || val != "value" {
			t.Errorf("GetVars() did not return the expected value, got %v", vars)
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
