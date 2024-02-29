// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"crypto/rsa"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPublicKey is a mock implementation of the PublicKey interface
type MockPublicKey struct {
	mock.Mock
}

func (m *MockPublicKey) GetPublicKey() any {
	args := m.Called()
	return args.Get(0).(any)
}

func NewMockPublicKey() *MockPublicKey {
	mockPublicKey := &MockPublicKey{}
	return mockPublicKey
}

func NewMockRsaPublicKey() *MockPublicKey {
	mockPublicKey := &MockPublicKey{}
	mockPublicKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	return mockPublicKey
}

var _ appmodels.PublicKey = (*MockPublicKey)(nil)
