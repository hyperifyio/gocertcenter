// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	"net/http"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockSwaggerManager defines a mock for apitypes.ISwaggerManager interface.
type MockSwaggerManager struct {
	mock.Mock
}

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

var _ managers.ISwaggerManager = (*MockSwaggerManager)(nil)
