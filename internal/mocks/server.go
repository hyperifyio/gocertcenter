// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockServer is a mock implementation of the IServer interface for testing purposes.
type MockServer struct {
	mock.Mock
	Address string
}

// Start simulates starting the server. It doesn't do anything in the mock.
func (m *MockServer) Start() error {
	args := m.Called()
	return args.Error(0) // Return nil or an error depending on what's set in your test
}

// SetupRoutes simulates setting up routes. It doesn't do anything in the mock.
func (m *MockServer) SetupRoutes() {
	m.Called()
}

// GetAddress returns a mock server address.
func (m *MockServer) GetAddress() string {
	args := m.Called()
	return args.String(0) // Return the Address field or a test-specific address
}

// NewMockServer creates an instance of MockServer with default values for testing.
func NewMockServer() *MockServer {
	mockServer := &MockServer{}
	// Setup default return values for methods if needed
	mockServer.On("Start").Return(nil)
	mockServer.On("GetAddress").Return("http://localhost:8080")
	// Ensure you setup the mock ruleset here or in your tests directly
	return mockServer
}
