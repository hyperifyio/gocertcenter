// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers_test

import (
	"context"
	"net/http"
	"testing"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/managers"
)

func TestNewSwaggerManager(t *testing.T) {
	router := mux.NewRouter()
	ctx := context.Background()
	info := &openapi3.Info{Title: "Test API", Version: "1.0"}

	manager, err := managers.NewSwaggerManager(router, &ctx, "http://example.com", "Test API Server", info)
	if err != nil {
		t.Fatalf("Failed to create SwaggerManager: %v", err)
	}

	if manager == nil {
		t.Fatal("Expected SwaggerManager instance, got nil")
	}
}

func TestSwaggerManager_AddRoute(t *testing.T) {
	router := mux.NewRouter()
	ctx := context.Background()
	info := &openapi3.Info{Title: "Test API", Version: "1.0"}
	manager, _ := managers.NewSwaggerManager(router, &ctx, "http://example.com", "Test API Server", info)

	testPath := "/test"
	testMethod := http.MethodGet
	definitions := swagger.Definitions{}

	route, err := manager.AddRoute(testMethod, testPath, func(w http.ResponseWriter, r *http.Request) {}, definitions)
	if err != nil {
		t.Fatalf("Failed to add route: %v", err)
	}

	if route == nil {
		t.Fatal("Expected non-nil *mux.Route, got nil")
	}

	// Optionally, you could make a request to the testPath and verify the handler is called
	// This would require starting the router in a test server
}

func TestSwaggerManager_GenerateAndExposeOpenapi(t *testing.T) {
	router := mux.NewRouter()
	ctx := context.Background()
	info := &openapi3.Info{Title: "Test API", Version: "1.0"}
	manager, _ := managers.NewSwaggerManager(router, &ctx, "http://example.com", "Test API Server", info)

	err := manager.GenerateAndExposeOpenapi()
	if err != nil {
		t.Errorf("Expected no error from GenerateAndExposeOpenapi, got %v", err)
	}

	// Optionally, verify the OpenAPI spec is accessible
	// This would involve making a request to the OpenAPI spec endpoint and checking the response
}

type UnsupportedType struct {
	Func func() `json:"func"` // Functions are not supported by JSON encoding.
}

func TestSwaggerManager_AddRouteWithError(t *testing.T) {

	// Setup: Assuming you have a mock implementation of swagger.Router that can simulate an AddRoute error
	router := mux.NewRouter()
	ctx := context.Background()
	url := "http://example.com"
	description := "Test API Server"
	info := &openapi3.Info{Title: "Test API", Version: "1.0"}

	manager, err := managers.NewSwaggerManager(router, &ctx, url, description, info)
	if err != nil {
		t.Fatalf("Failed to create SwaggerManager: %v", err)
	}

	// Act: Attempt to add a route which is designed to fail within the mock
	_, err = manager.AddRoute(
		"GET",
		"/test",
		func(w http.ResponseWriter, r *http.Request) {
		},
		// This Definitions should create an error
		swagger.Definitions{
			Summary: "",
			Extensions: map[string]interface{}{
				"s": nil,
			},
		},
	)
	if err == nil {
		t.Error("Expected an error when adding a route, but got nil")
	}
}

func TestNewSwaggerManagerWithError(t *testing.T) {

	// Setup: Assuming there's a way to force swagger.NewRouter to return an error, perhaps via dependency injection or mocking
	router := mux.NewRouter()
	ctx := context.Background()
	url := "http://example.com"
	description := "Unable to create API server"
	info := &openapi3.Info{} // Title and version are required

	// Act: Attempt to create a SwaggerManager which is designed to fail
	_, err := managers.NewSwaggerManager(router, &ctx, url, description, info)
	if err == nil {
		t.Error("Expected an error when creating SwaggerManager, but got nil")
	}

}
