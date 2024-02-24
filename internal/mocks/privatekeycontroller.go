// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// MockPrivateKeyController is a mock implementation of models.IPrivateKeyController interface.
type MockPrivateKeyController struct {
	UsesPrivateKeyServiceFunc func(service models.IPrivateKeyService) bool
	// These fields allow you to specify the behavior and output of the mock methods.
	GetExistingPrivateKeyFunc func(serialNumber models.ISerialNumber) (models.IPrivateKey, error)
	CreatePrivateKeyFunc      func(key models.IPrivateKey) (models.IPrivateKey, error)
}

var _ models.IPrivateKeyController = (*MockPrivateKeyController)(nil)

func (m *MockPrivateKeyController) UsesPrivateKeyService(service models.IPrivateKeyService) bool {
	return m.UsesPrivateKeyServiceFunc(service)
}

// GetExistingPrivateKey simulates retrieving an existing private key by serial number.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyController) GetExistingPrivateKey(serialNumber models.ISerialNumber) (models.IPrivateKey, error) {
	if m.GetExistingPrivateKeyFunc != nil {
		return m.GetExistingPrivateKeyFunc(serialNumber)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}

// CreatePrivateKey simulates creating a new private key.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyController) CreatePrivateKey(key models.IPrivateKey) (models.IPrivateKey, error) {
	if m.CreatePrivateKeyFunc != nil {
		return m.CreatePrivateKeyFunc(key)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}
