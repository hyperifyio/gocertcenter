// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPrivateKeyService is a mock implementation of models.PrivateKeyRepository interface.
type MockPrivateKeyService struct {
	mock.Mock
}

// GetExistingPrivateKey simulates retrieving an existing private key by serial number.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) FindByOrganizationAndSerialNumbers(organization string, certificates []*big.Int) (appmodels.PrivateKey, error) {
	args := m.Called(organization, certificates)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.PrivateKey), args.Error(1)
}

// CreatePrivateKey simulates creating a new private key.
// It uses a function field to allow custom behavior for each test.
func (m *MockPrivateKeyService) Save(key appmodels.PrivateKey) (appmodels.PrivateKey, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.PrivateKey), args.Error(1)
}

var _ appmodels.PrivateKeyRepository = (*MockPrivateKeyService)(nil)
