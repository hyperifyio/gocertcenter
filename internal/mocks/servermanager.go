// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/stretchr/testify/mock"
	"net"
)

// MockServerManager is a mock type for apitypes.IServerManager interface.
type MockServerManager struct {
	mock.Mock
}

var _ apitypes.IServerManager = (*MockServerManager)(nil)

// Serve is a mock method that simulates starting the server.
func (m *MockServerManager) Serve(l net.Listener) error {
	args := m.Called(l)
	return args.Error(0)
}

// Shutdown is a mock method that simulates shutting down the server.
func (m *MockServerManager) Shutdown() error {
	args := m.Called()
	return args.Error(0)
}
