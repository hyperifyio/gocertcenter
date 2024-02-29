// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPrivateKeyService is a mock implementation of models.IPrivateKeyService interface.
type MockPrivateKeyService struct {
	mock.Mock
}

// GetExistingPrivateKey simulates retrieving an existing private key by serial number.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
	args := m.Called(organization, certificates)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.IPrivateKey), args.Error(1)
}

// CreatePrivateKey simulates creating a new private key.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) Save(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.IPrivateKey), args.Error(1)
}

var _ appmodels.IPrivateKeyService = (*MockPrivateKeyService)(nil)
