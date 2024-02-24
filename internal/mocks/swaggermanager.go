// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/gorilla/mux"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/stretchr/testify/mock"
	"net/http"
)

// MockSwaggerManager defines a mock for apitypes.ISwaggerManager interface.
type MockSwaggerManager struct {
	mock.Mock
}

var _ apitypes.ISwaggerManager = (*MockSwaggerManager)(nil)

// GenerateAndExposeOpenapi mocks the GenerateAndExposeOpenapi method.
func (m *MockSwaggerManager) GenerateAndExposeOpenapi() error {
	args := m.Called()
	return args.Error(0)
}

// AddRoute mocks the AddRoute method.
func (m *MockSwaggerManager) AddRoute(method string, path string, handler http.HandlerFunc, schema swagger.Definitions) (*mux.Route, error) {
	args := m.Called(method, path, handler, schema)
	return args.Get(0).(*mux.Route), args.Error(1)
}
