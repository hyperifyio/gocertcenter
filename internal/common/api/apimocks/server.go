// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// MockServer is a mock implementation of apitypes.IServer interface for testing purposes.
type MockServer struct {
	mock.Mock
	Address string
	URL     string
	Info    *openapi3.Info
}

func (m *MockServer) IsStarted() bool {
	args := m.Called()
	return args.Bool(0)
}

// GetAddress returns a mock server address.
func (m *MockServer) GetAddress() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockServer) SetInfo(info *openapi3.Info) {
	m.Called(info)
}

func (m *MockServer) GetInfo() *openapi3.Info {
	args := m.Called()
	info, _ := args.Get(0).(*openapi3.Info)
	return info
}

// GetURL returns a mock server URL.
func (m *MockServer) GetURL() string {
	args := m.Called()
	return args.String(0)
}

// InitSetup simulates starting the server. It doesn't do anything in the mock.
func (m *MockServer) InitSetup() error {
	args := m.Called()
	return args.Error(0)
}

// SetupRoutes simulates setting up routes. It doesn't do anything in the mock.
func (m *MockServer) SetupRoutes(routes []apitypes.Route) error {
	args := m.Called(routes)
	return args.Error(0)
}

// SetupHandler simulates setting up routes. It doesn't do anything in the mock.
func (m *MockServer) SetupHandler(
	method string,
	path string,
	apiHandler apitypes.RequestHandlerFunc,
	definitions swagger.Definitions,
) error {
	args := m.Called(method, path, apiHandler, definitions)
	return args.Error(0)
}

func (m *MockServer) SetupNotFoundHandler(
	handler apitypes.RequestHandlerFunc,
) {
	m.Called(handler)
}

func (m *MockServer) SetupMethodNotAllowedHandler(
	handler apitypes.RequestHandlerFunc,
) {
	m.Called(handler)
}

func (m *MockServer) FinalizeSetup() error {
	args := m.Called()
	return args.Error(0)
}

// Start simulates starting the server. It doesn't do anything in the mock.
func (m *MockServer) Start() error {
	args := m.Called()
	return args.Error(0)
}

// Stop simulates stopping the server. It doesn't do anything in the mock.
func (m *MockServer) Stop() error {
	args := m.Called()
	return args.Error(0)
}

// NewMockServer creates an instance of MockServer with default values for testing.
func NewMockServer() *MockServer {
	mockServer := &MockServer{}
	// Setup default return values for methods if needed
	mockServer.On("IsStarted").Return(false)
	mockServer.On("GetAddress").Return(":8080")
	mockServer.On("GetURL").Return("http://localhost:8080")
	mockServer.On("GetInfo").Return(&openapi3.Info{
		Title:   "Mock title",
		Version: "0.0.0",
	})
	mockServer.On("SetInfo").Return(nil)
	mockServer.On("InitSetup").Return(nil)
	mockServer.On("SetupRoutes").Return(nil)
	mockServer.On("SetupHandler").Return(nil)
	mockServer.On("SetupNotFoundHandler").Return(nil)
	mockServer.On("SetupMethodNotAllowedHandler").Return(nil)
	mockServer.On("FinalizeSetup").Return(nil)
	mockServer.On("Start").Return(nil)
	mockServer.On("Stop").Return(nil)
	// Ensure you setup the mock ruleset here or in your tests directly
	return mockServer
}

var _ apitypes.IServer = (*MockServer)(nil)
