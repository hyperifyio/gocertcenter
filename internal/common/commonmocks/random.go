// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package commonmocks

import (
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockRandomManager is a mock implementation of models.IRandomManager for testing.
type MockRandomManager struct {
	CreateBigIntFunc func(max *big.Int) (*big.Int, error)
}

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

var _ managers.IRandomManager = (*MockRandomManager)(nil)
