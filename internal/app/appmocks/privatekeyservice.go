// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPrivateKeyService is a mock implementation of models.IPrivateKeyService interface.
type MockPrivateKeyService struct {
	// These fields allow you to specify the behavior and output of the mock methods.
	GetExistingPrivateKeyFunc func(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error)
	CreatePrivateKeyFunc      func(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error)
}

// GetExistingPrivateKey simulates retrieving an existing private key by serial number.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) GetExistingPrivateKey(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
	if m.GetExistingPrivateKeyFunc != nil {
		return m.GetExistingPrivateKeyFunc(organization, certificates)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}

// CreatePrivateKey simulates creating a new private key.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) CreatePrivateKey(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
	if m.CreatePrivateKeyFunc != nil {
		return m.CreatePrivateKeyFunc(key)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}

var _ appmodels.IPrivateKeyService = (*MockPrivateKeyService)(nil)
