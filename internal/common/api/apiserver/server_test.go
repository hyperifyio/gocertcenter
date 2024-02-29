// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package apiserver_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apiserver"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestNewServer(t *testing.T) {
	listen := "localhost:8080"
	server, err := apiserver.NewServer(listen, nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	if server.GetAddress() != listen {
		t.Errorf("NewServer listen = %s; want %s", server.GetAddress(), listen)
	}
}

// Assuming Start and Stop methods are updated to start/stop an HTTP server
func TestServer_StartStop(t *testing.T) {

	listenAddr := "localhost:8080"

	server, err := apiserver.NewServer(listenAddr, nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Wait for the server to start
	time.Sleep(time.Second)

	// Attempt to connect to the server
	resp, err := http.Get("http://" + listenAddr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	resp.Body.Close() // Close the body to avoid resource leaks

	// Now test stopping the server
	server.Stop()

	// Attempt to connect to the server after stopping it
	_, err = http.Get("http://" + listenAddr)
	if err == nil {
		t.Fatal("Server should not be reachable after stopping")
	}
}

func TestIsStarted(t *testing.T) {
	server, err := apiserver.NewServer("localhost:8080", nil)

	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if server.IsStarted() {
		t.Error("Expected server to be not started")
	}

	// Mock starting the server (set server.server to non-nil)
	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Stop()

	if !server.IsStarted() {
		t.Error("Expected server to be marked as started")
	}
}

func TestGetURL(t *testing.T) {
	tests := []struct {
		listen string
		want   string
	}{
		{":8080", "http://localhost:8080"},
		{"localhost:8080", "http://localhost:8080"},
	}

	for _, tt := range tests {
		server, err := apiserver.NewServer(tt.listen, nil)
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}

		got := server.GetURL()
		if got != tt.want {
			t.Errorf("GetURL() = %v, want %v", got, tt.want)
		}
	}
}

func TestServer_StartError(t *testing.T) {
	server, err := apiserver.NewServer("invalidAddress", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	err = server.Start()
	if err == nil {
		t.Fatal("Expected error when starting server with invalid address")
	}

	// Clean up
	server.Stop()
}

func TestServer_StartTwice(t *testing.T) {
	listenAddr := "localhost:8081" // Use a different port to avoid conflicts

	server, err := apiserver.NewServer(listenAddr, nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Try to start again
	err = server.Start()
	if err == nil {
		t.Fatal("Expected error when starting server twice")
	}

	// Clean up
	server.Stop()
}

func TestStart_ServerAlreadyStarted(t *testing.T) {
	server, err := apiserver.NewServer("localhost:8082", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Stop()

	// First start should succeed
	if err := server.Start(); err != nil {
		t.Fatalf("First start of the server failed: %v", err)
	}

	// Second start should fail
	if err := server.Start(); err == nil {
		t.Fatal("Expected error when starting an already started server, got nil")
	} else if !strings.Contains(err.Error(), "[server] Already started") {
		t.Fatalf("Expected '[server] Already started' error, got: %v", err)
	}
}

func TestStart_ListenerAlreadyExists(t *testing.T) {

	server, err := apiserver.NewServer("localhost:8083", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	defer server.Stop()

	// Manually create a listener to simulate already listening state
	ln, err := net.Listen("tcp", server.GetAddress())
	if err != nil {
		t.Fatalf("Failed to manually listen on address: %v", err)
	}
	defer ln.Close()

	err = server.SetListener(&ln)
	if err != nil {
		t.Fatalf("Failed to set listener on address: %v", err)
	}

	// Attempt to start the server should fail because listener already exists
	if err := server.Start(); err == nil {
		t.Fatal("Expected error when starting server with existing listener, got nil")
	} else if !strings.Contains(err.Error(), "[server] Already listening") {
		t.Fatalf("Expected '[server] Already listening' error, got: %v", err)
	}
}

func TestStart_DefaultAddress(t *testing.T) {
	server, err := apiserver.NewServer("", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Stop()

	// Starting server with default address
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server with default address: %v", err)
	}

	// You might want to make a test HTTP request here to ensure the server is listening
	// Note: This test assumes ":http" is a valid listen address and might need adjustment
}

func TestStart_FailAfterShortWait(t *testing.T) {

	// Create an occupied port or use a common one like 8081 for localhost
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal("Expected port 8081 to be free: ", err)
	}
	defer ln.Close()

	server, err := apiserver.NewServer(":8081", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	err = server.Start()
	if err == nil {
		t.Fatal("Expected server to fail starting on an occupied port, but it did not")
	}

	if !strings.Contains(err.Error(), "address already in use") {
		t.Fatalf("Expected an 'address already in use' error, got: %v", err)
	}
}

func TestServer_InitSetup_Idempotent(t *testing.T) {
	server, err := apiserver.NewServer(":8080", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	err = server.InitSetup() // First initialization
	if err != nil {
		t.Fatalf("Failed to init server: %v", err)
	}

	// Get the initial hash of the server state
	initialHash, err := server.GetInternalHash()
	if err != nil {
		t.Fatalf("Failed to hash server state: %v", err)
	}

	server.InitSetup() // Second initialization

	secondHash, err := server.GetInternalHash()
	if err != nil {
		t.Fatalf("Failed to hash server state again: %v", err)
	}

	// Compare the initial hash with the hash after the second initialization
	if secondHash != initialHash {
		t.Fatal("InitSetup should not reinitialize router or swaggerRouter")
	}

}

func TestServer_SetListener_AlreadySet(t *testing.T) {
	server, err := apiserver.NewServer(":8080", nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// Set listener for the first time
	err = server.SetListener(&listener)
	if err != nil {
		t.Fatalf("Failed to set listener: %v", err)
	}

	// Attempt to set another listener
	err = server.SetListener(&listener)
	if err == nil {
		t.Fatal("Expected error when setting listener a second time, got nil")
	}
}

func TestServer_SetupHandler(t *testing.T) {
	// Setup the server on a test port
	listenAddr := ":8084" // Use a unique port to avoid conflicts with other tests
	server, err := apiserver.NewServer(listenAddr, nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Define a test route and handler
	testPath := "/test-handler"
	testResponse := "handler response"
	err = server.SetupHandler(http.MethodGet, testPath, testHandler(testResponse), swagger.Definitions{})
	if err != nil {
		t.Fatalf("Failed to setup the server: %v", err)
	}

	// Start the server
	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Stop()

	// Wait a moment for the server to start
	time.Sleep(100 * time.Millisecond)

	// Make a request to the test route
	resp, err := http.Get("http://localhost" + listenAddr + testPath)
	if err != nil {
		t.Fatalf("Failed to make request to test route: %v", err)
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Verify the response body
	var responseBody map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	if responseBody["message"] != testResponse {
		t.Errorf("Expected response message to be '%s', got '%s'", testResponse, responseBody["message"])
	}
}

func TestServer_SetupRoutes(t *testing.T) {

	// Define multiple test routes with unique response messages
	var testRoutes = []apitypes.Route{
		{
			Method:      http.MethodGet,
			Path:        "/route1",
			Handler:     testHandler("response from route1"),
			Definitions: swagger.Definitions{}, // Simplified for example; adjust as needed
		},
		{
			Method:      http.MethodPost,
			Path:        "/route2",
			Handler:     testHandler("response from route2"),
			Definitions: swagger.Definitions{}, // Simplified for example; adjust as needed
		},
		// Add more routes as needed for comprehensive testing
	}

	// Setup server on a test port
	listenAddr := ":8085" // Ensure this port is unique and not used by other tests or services
	server, err := apiserver.NewServer(listenAddr, nil)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Call SetupRoutes with the defined test routes
	err = server.SetupRoutes(testRoutes)
	if err != nil {
		t.Fatalf("Failed to setup routes: %v", err)
	}

	// Start the server
	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Stop()

	// Wait a moment for the server to start
	time.Sleep(100 * time.Millisecond)

	// Test each route
	for _, route := range testRoutes {
		// Create a request to the current route's path
		resp, err := HttpRequest(route.Method, fmt.Sprintf("http://localhost%s%s", listenAddr, route.Path), nil)
		if err != nil {
			t.Fatalf("Failed to make request to %s: %v", route.Path, err)
		}
		defer resp.Body.Close()

		// Verify the response status code
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("%s: Expected status code %d, got %d", route.Path, http.StatusOK, resp.StatusCode)
		}

		// Verify the response body
		var responseBody map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			t.Fatalf("Failed to decode response body for %s: %v", route.Path, err)
		}
		expectedMessage := "response from " + strings.Trim(route.Path, "/")
		if responseBody["message"] != expectedMessage {
			t.Errorf("%s: Expected response message to be '%s', got '%s'", route.Path, expectedMessage, responseBody["message"])
		}
	}
}

// testHandler returns a simple HTTP handler function for testing purposes.
func testHandler(responseContent string) apitypes.RequestHandlerFunc {
	return func(response apitypes.IResponse, request apitypes.IRequest) error {
		response.Send(http.StatusOK, map[string]string{"message": responseContent})
		return nil
	}
}

func HttpRequest(method string, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func TestSetupHandler(t *testing.T) {
	server, err := apiserver.NewServer(":8086", nil) // Ensure this port is unique and not used by other tests
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Stop()

	// Example route to test
	testPath := "/setup-handler-test"
	testMethod := http.MethodGet
	testResponse := "SetupHandler works!"

	// Setup the handler
	err = server.SetupHandler(testMethod, testPath, func(response apitypes.IResponse, request apitypes.IRequest) error {
		response.Send(http.StatusOK, map[string]interface{}{"message": testResponse})
		return nil
	}, swagger.Definitions{})
	if err != nil {
		t.Fatalf("SetupHandler failed: %v", err)
	}

	// Start the server
	if err := server.Start(); err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}

	// Wait a moment for the server to start
	time.Sleep(100 * time.Millisecond)

	// Make a request to the test route
	resp, err := http.Get("http://localhost:8086" + testPath)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Verify response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if got, want := responseBody["message"], testResponse; got != want {
		t.Errorf("Got response message %v, want %v", got, want)
	}
}

func TestStop(t *testing.T) {
	server, err := apiserver.NewServer(":8087", nil) // Ensure this port is unique
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Ensure server is running
	_, err = http.Get("http://localhost:8087")
	if err != nil {
		t.Fatalf("Server should be running, but got error: %v", err)
	}

	// Stop the server
	if err := server.Stop(); err != nil {
		t.Fatalf("Failed to stop server: %v", err)
	}

	// Ensure server has stopped
	_, err = http.Get("http://localhost:8087")
	if err == nil {
		t.Fatal("Server should be stopped, but request succeeded")
	}
}

func TestServer_FinalizeSetup(t *testing.T) {

	// Mock swagger factory function that always returns an error
	mockSwaggerManager := new(apimocks.MockSwaggerManager)
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	server, err := apiserver.NewServer(":8080", mockFactory)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Expect the FinalizeSetup to call GenerateAndExposeOpenapi and succeed
	mockSwaggerManager.On("GenerateAndExposeOpenapi").Return(nil)

	err = server.FinalizeSetup()
	if err != nil {
		t.Errorf("FinalizeSetup failed: %v", err)
	}

	mockSwaggerManager.AssertCalled(t, "GenerateAndExposeOpenapi")
}

func TestServer_SetupHandlerWithMock(t *testing.T) {

	mockRoute := mux.NewRouter().NewRoute()

	// Initialize the mock SwaggerManager
	// Mock swagger factory function that always returns an error
	mockSwaggerManager := new(apimocks.MockSwaggerManager)
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	// Creating server with the mocked SwaggerManager
	server, err := apiserver.NewServer(":8080", mockFactory)
	if err != nil {
		t.Fatalf("NewServer failed: %v", err)
	}

	// Define test parameters
	path := "/test"
	method := "GET"
	definitions := swagger.Definitions{}
	var handler apitypes.RequestHandlerFunc = func(response apitypes.IResponse, request apitypes.IRequest) error {
		return nil
	}

	// Mock the AddRoute method
	// Since AddRoute is expected to return (*mux.Route, error), ensure the mock does the same.
	// For simplicity, nil is returned for *mux.Route because it's not directly used or checked in this test scenario.
	mockSwaggerManager.On("AddRoute", method, path, mock.AnythingOfType("http.HandlerFunc"), definitions).Return(mockRoute, nil)

	// Attempt to setup a handler
	err = server.SetupHandler(method, path, handler, definitions)
	if err != nil {
		t.Errorf("SetupHandler failed: %v", err)
	}

	// Verify that AddRoute was called with expected parameters
	// Correct the method name from "AddRouter" to "AddRoute"
	mockSwaggerManager.AssertCalled(t, "AddRoute", method, path, mock.AnythingOfType("http.HandlerFunc"), definitions)
}

func TestSetupHandlerEmptyMethodPath(t *testing.T) {

	// Mock the AddRoute method to return a mock IRoute
	mockRoute := mux.NewRouter().NewRoute()

	// Create a new instance of MockSwaggerManager
	mockSwaggerManager := new(apimocks.MockSwaggerManager)
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	// Create server instance with the mock
	server, err := apiserver.NewServer(":8080", mockFactory)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Set up expectations
	mockSwaggerManager.On("AddRoute", "GET", "/", mock.AnythingOfType("http.HandlerFunc"), swagger.Definitions{}).Return(mockRoute, nil)

	var handler apitypes.RequestHandlerFunc = func(response apitypes.IResponse, request apitypes.IRequest) error {
		return nil
	}

	// Call the method under test
	err = server.SetupHandler(
		"",
		"",
		handler,
		swagger.Definitions{},
	)

	// Assert expectations
	mockSwaggerManager.AssertExpectations(t)

	// Check for errors
	if err != nil {
		t.Errorf("SetupHandler failed: %v", err)
	}
}

func TestInitSetup_SwaggerFactoryError(t *testing.T) {

	expectedErrorMsg := "mock error"

	// Mock swagger factory function that always returns an error
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return nil, errors.New(expectedErrorMsg)
	}

	// Create a new server instance
	server, err := apiserver.NewServer(
		":8080",
		mockFactory,
	)

	if server != nil {
		t.Fatalf("Expected InitSetup to fail when swaggerFactory fails, but got a server")
	}

	// NewServer calls InitSetup(), which should fail
	if err == nil {
		t.Fatalf("Expected an error from InitSetup when swaggerFactory fails, but got nil")
	}

	// Check if the actual error message contains the expected substring
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedErrorMsg, err.Error())
	}

}

func TestSetupRoutes_HandlerError(t *testing.T) {

	expectedErrorMsg := "add route failed"

	mockRoute := mux.NewRouter().NewRoute()

	// Create a new instance of MockSwaggerManager
	mockSwaggerManager := new(apimocks.MockSwaggerManager)
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	mockSwaggerManager.On("AddRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRoute, errors.New(expectedErrorMsg))

	// Create a server and inject the mock swagger manager
	server, err := apiserver.NewServer(
		":8080",
		mockFactory,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	var handler apitypes.RequestHandlerFunc = func(response apitypes.IResponse, request apitypes.IRequest) error {
		return nil
	}

	// Define a test route that would trigger the AddRoute call
	testRoutes := []apitypes.Route{
		{
			Method:      "GET",
			Path:        "/test",
			Handler:     handler,
			Definitions: swagger.Definitions{},
		},
	}

	// Call SetupRoutes and expect it to return an error
	err = server.SetupRoutes(testRoutes)
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedErrorMsg, err.Error())
	}

}

func TestServerInitAfterUnSetup(t *testing.T) {
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Unsetup the server
	if err := server.UnSetup(); err != nil {
		t.Fatalf("Failed to unsetup the server: %v", err)
	}

	// Attempt to restart the server. This should call InitSetup() again.
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to restart server: %v", err)
	}

}

func TestServer_Start_FailsFinalizeSetup(t *testing.T) {

	// Create a new instance of MockSwaggerManager
	mockSwaggerManager := new(apimocks.MockSwaggerManager)
	mockFactory := func(
		router *mux.Router,
		context *context.Context,
		url string,
		description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	// Create a mock Swagger manager that will fail on GenerateAndExposeOpenapi
	mockSwaggerManager.On("GenerateAndExposeOpenapi").Return(errors.New("finalize setup failed"))

	// Initialize the server with the mock swagger manager
	server, err := apiserver.NewServer(
		"localhost:8080",
		mockFactory,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Attempt to start the server
	err = server.Start()
	if err == nil {
		t.Fatal("Expected Start to fail due to FinalizeSetup error, but it did not")
	}

	// Check if the error message is as expected
	expectedErrorMsg := "finalize setup failed"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedErrorMsg, err.Error())
	}

}

// Custom factory function that always returns nil
func nilServerFactory(address string, router *mux.Router) managers.IServerManager {
	return nil
}

func TestServer_StartWithNilServerFactory(t *testing.T) {
	// Setup
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Inject the custom nil server factory function
	server.SetServerFactory(nilServerFactory)

	// Attempt to start the server
	err = server.Start()

	// Verify
	if err == nil {
		t.Errorf("Expected Start to fail due to nil serverFactory result, but it did not")
	} else if !strings.Contains(err.Error(), "failed to initialize server") {
		t.Errorf("Expected failure to initialize server error, got: %v", err)
	}
}

// mockServerManager simulates the IServerManager behavior for testing.
type slowMockServerManager struct {
	Time time.Duration
}

// Serve simulates a server operation that takes 1.5 seconds and then fails.
func (m *slowMockServerManager) Serve(l net.Listener) error {
	time.Sleep(m.Time) // Simulate delay
	return errors.New("server failed intentionally for test")
}

// Shutdown simulates a successful shutdown.
func (m *slowMockServerManager) Shutdown() error {
	return nil
}

func TestServer_StartWithShortDelayedServeFailure(t *testing.T) {
	// Create the server with a mock server factory function.
	server, err := apiserver.NewServer(
		"localhost:0",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Inject the mock server manager factory that returns our custom mockServerManager.
	server.SetServerFactory(func(address string, router *mux.Router) managers.IServerManager {
		return &slowMockServerManager{Time: 100 * time.Millisecond}
	})

	// Attempt to start the server.
	err = server.Start()

	// Verify that the server attempted to start and failed after the expected delay.
	if err == nil {
		t.Errorf("Expected Start to fail due to simulated short server failure, but it did not")
	} else if !strings.Contains(err.Error(), "server failed intentionally for test") {
		t.Errorf("Expected specific failure message, got: %v", err)
	}

	time.Sleep(2 * time.Second)

}

func TestServer_StartWithLongDelayedServeFailure(t *testing.T) {

	// Create the server with a mock server factory function.
	server, err := apiserver.NewServer(
		"localhost:0",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Inject the mock server manager factory that returns our custom mockServerManager.
	server.SetServerFactory(func(address string, router *mux.Router) managers.IServerManager {
		return &slowMockServerManager{Time: 1600 * time.Millisecond}
	})

	// Attempt to start the server.
	err = server.Start()

	// Verify that the server attempted to start and did not fail in the expected short delay.
	if err != nil {
		t.Errorf("Expected Start to not fail due to simulated long server failure, but it did: %v", err)
	}

	time.Sleep(2 * time.Second)

}

func TestServer_SetInfo(t *testing.T) {
	// Initialize the server with minimal setup.
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Create a new Info object to set on the server.
	newInfo := &openapi3.Info{
		Title:       "New API Title",
		Version:     "1.0.1",
		Description: "A new description for the API.",
	}

	// Call SetInfo on the server with the new Info object.
	server.SetInfo(newInfo)

	// Retrieve the Info from the server to verify it was set correctly.
	setInfo := server.GetInfo()

	// Assert that the Info set on the server matches the new Info object.
	if setInfo.Title != newInfo.Title || setInfo.Version != newInfo.Version || setInfo.Description != newInfo.Description {
		t.Errorf("SetInfo did not correctly update the server's Info field")
	}
}

func TestServer_SetSwaggerFactory(t *testing.T) {

	mockSwaggerManager := &apimocks.MockSwaggerManager{}

	// Mock factory function that returns an instance of our mockSwaggerManager
	factory := func(router *mux.Router, ctx *context.Context, url, description string, info *openapi3.Info) (managers.ISwaggerManager, error) {
		return mockSwaggerManager, nil
	}

	// Create the server with minimal setup.
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	hash, err := server.GetInternalHash()
	if err != nil {
		t.Fatalf("Failed to hash server state: %v", err)
	}

	// Call SetSwaggerFactory on the server with the mock factory function.
	server.SetSwaggerFactory(factory)

	nextHash, err := server.GetInternalHash()
	if err != nil {
		t.Fatalf("Failed to hash server state again: %v", err)
	}

	if hash != nextHash {
		t.Errorf("SetSwaggerFactory did not correctly update internal state")
	}
}

func TestServer_StopShutdownFails(t *testing.T) {
	// Define the error you expect to be returned by the Shutdown method
	shutdownErr := errors.New("shutdown failed")

	// Create an instance of the mock server manager
	mockServerManager := new(apimocks.MockServerManager)

	// Setup the expectation for the Shutdown method to be called and to return the defined error
	mockServerManager.On("Shutdown").Return(shutdownErr)

	// Initialize the server, passing any required dependencies
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Inject the mock server manager factory that returns our custom mockServerManager.
	server.SetServerFactory(func(address string, router *mux.Router) managers.IServerManager {
		return mockServerManager
	})

	err = server.Start()

	// Attempt to stop the server
	err = server.Stop()

	// Verify the error is as expected
	if err == nil {
		t.Errorf("Expected failure due to simulated shutdown failure, but it did not")
	} else if !strings.Contains(err.Error(), "shutdown failed") {
		t.Errorf("Expected failure to shutdown, got: %v", err)
	}

	// Assert that Shutdown was called exactly once
	mockServerManager.AssertNumberOfCalls(t, "Shutdown", 1)

}

func TestServer_UnSetupWhileRunning(t *testing.T) {
	// Create an instance of the mock server manager
	mockServerManager := new(apimocks.MockServerManager)

	// No need to mock the Serve method for this test, but we initialize the server as if it's running.

	// Initialize the server, passing any required dependencies
	server, err := apiserver.NewServer(
		"localhost:8080",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Inject the mock server manager factory to return our custom mockServerManager.
	// This step simulates the server being in a running state by directly setting the mock manager.
	server.SetServerFactory(func(address string, router *mux.Router) managers.IServerManager {
		return mockServerManager
	})

	err = server.Start()

	// Attempt to unsetup the server while it is "running"
	err = server.UnSetup()

	// Verify the error is as expected: it should fail since the server is considered running.
	if err == nil {
		t.Errorf("Expected UnSetup to fail because the server is running, but it did not")
	} else if !strings.Contains(err.Error(), "Cannot revert route setup, server still running") {
		t.Errorf("Expected failure to shutdown, got: %v", err)
	}
}

func TestStart_UninitializedHandle_SwaggerFactoryFails(t *testing.T) {

	// Custom swagger factory function that simulates failure
	failingSwaggerFactory := func(
		router *mux.Router,
		context *context.Context,
		basePath, description string,
		info *openapi3.Info,
	) (managers.ISwaggerManager, error) {
		return nil, errors.New("swagger factory initialization failed")
	}

	// Create the server with the failing swagger factory
	server := apiserver.NewUninitializedServer("localhost:8080", failingSwaggerFactory)

	// Attempt to start the server, expecting it to fail due to the failing swagger factory
	err := server.Start()
	assert.NotNil(t, err, "Expected Start to fail due to swagger factory failure")
	assert.Contains(t, err.Error(), "swagger factory initialization failed", "Error message should indicate failure of swagger factory initialization")
}
