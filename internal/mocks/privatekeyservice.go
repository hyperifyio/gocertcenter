// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// MockPrivateKeyService is a mock implementation of the IPrivateKeyService interface.
type MockPrivateKeyService struct {
	// These fields allow you to specify the behavior and output of the mock methods.
	GetExistingPrivateKeyFunc func(serialNumber models.SerialNumber) (*models.PrivateKey, error)
	CreatePrivateKeyFunc      func(key *models.PrivateKey) (*models.PrivateKey, error)
}

// GetExistingPrivateKey simulates retrieving an existing private key by serial number.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) GetExistingPrivateKey(serialNumber models.SerialNumber) (*models.PrivateKey, error) {
	if m.GetExistingPrivateKeyFunc != nil {
		return m.GetExistingPrivateKeyFunc(serialNumber)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}

// CreatePrivateKey simulates creating a new private key.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) CreatePrivateKey(key *models.PrivateKey) (*models.PrivateKey, error) {
	if m.CreatePrivateKeyFunc != nil {
		return m.CreatePrivateKeyFunc(key)
	}
	// Default behavior or error can be returned here if not overridden by a test.
	return nil, nil
}
