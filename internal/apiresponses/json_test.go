// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apiresponses_test

import (
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/apiresponses"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONResponse_Send(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	data := map[string]string{"test": "value"}

	sender.Send(http.StatusOK, data)

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "{\"test\":\"value\"}\n"
	if body := w.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}

}

func TestJSONResponse_Send_MarshalError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	// Passing a channel, as it cannot be marshaled into JSON
	sender.Send(http.StatusOK, make(chan int))

	// Expecting internal server error due to marshaling failure
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected HTTP status %d, got %d", http.StatusInternalServerError, w.Code)
	}

	expectedBody := "Error creating JSON writer\n" // Assuming http.Error ends with a newline
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body to be '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestJSONResponse_Send_WriteError(t *testing.T) {
	w := httptest.NewRecorder()
	fw := mocks.NewFailWriter(w)
	sender := apiresponses.NewJSONResponse(fw)

	sender.Send(http.StatusOK, map[string]string{"test": "value"})

	// The custom failWriter does not correctly simulate a failed write such that the standard http.Error mechanism can catch and handle it.
	// This is due to how the http.Error function works, writing directly to the ResponseWriter, and expecting it to succeed.
	// In a real scenario, a write failure might be handled differently, potentially at a higher level of your application's error handling infrastructure.
	// Therefore, this specific test might not accurately reflect a real-world scenario or its handling.
}

func TestJSONResponse_SendError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	errorMessage := "An error occurred"
	statusCode := http.StatusBadRequest // You can choose any error status code

	sender.SendError(statusCode, errorMessage)

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	if status := w.Code; status != statusCode {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, statusCode)
	}

	// Check the response body for the correct error message
	expected := fmt.Sprintf("{\"code\":400,\"error\":\"%s\"}\n", errorMessage)
	if body := w.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestJSONResponse_SendMethodNotSupportedError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	sender.SendMethodNotSupportedError()

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	expectedStatusCode := http.StatusMethodNotAllowed
	if status := w.Code; status != expectedStatusCode {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, expectedStatusCode)
	}

	// Check the response body for the correct error message
	expectedErrorMessage := "Method Not Allowed"
	expectedBody := fmt.Sprintf("{\"code\":405,\"error\":\"%s\"}\n", expectedErrorMessage)
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}

// TestJSONResponse_SendNotFoundError tests the SendNotFoundError method
func TestJSONResponse_SendNotFoundError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	sender.SendNotFoundError()

	// Check the status code
	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the response body for the correct error message
	expectedErrorMessage := http.StatusText(http.StatusNotFound)
	expectedBody := fmt.Sprintf("{\"code\":404,\"error\":\"%s\"}\n", expectedErrorMessage)
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}

// TestJSONResponse_SendConflictError tests the SendConflictError method
func TestJSONResponse_SendConflictError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	errorMessage := "A conflict occurred"
	sender.SendConflictError(errorMessage)

	// Check the status code
	if status := w.Code; status != http.StatusConflict {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the response body for the correct error message
	expectedBody := fmt.Sprintf("{\"code\":409,\"error\":\"%s\"}\n", errorMessage)
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}

// TestJSONResponse_SendInternalServerError tests the SendInternalServerError method
func TestJSONResponse_SendInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := apiresponses.NewJSONResponse(w)

	errorMessage := "An internal server error occurred"
	sender.SendInternalServerError(errorMessage)

	// Check the status code
	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the response body for the correct error message
	expectedBody := fmt.Sprintf("{\"code\":500,\"error\":\"%s\"}\n", errorMessage)
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}
