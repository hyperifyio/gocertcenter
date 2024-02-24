// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
	"math/big"
)

// MockRandomManager is a mock implementation of models.IRandomManager for testing.
type MockRandomManager struct {
	CreateBigIntFunc func(max *big.Int) (*big.Int, error)
}

var _ models.IRandomManager = (*MockRandomManager)(nil)

func NewMockRandomManager() *MockRandomManager {
	return &MockRandomManager{}
}

// CreateBigInt calls the mocked function.
func (m *MockRandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	if m.CreateBigIntFunc != nil {
		return m.CreateBigIntFunc(max)
	}
	// Return nil or some default value if not specifically mocked
	return nil, nil
}
