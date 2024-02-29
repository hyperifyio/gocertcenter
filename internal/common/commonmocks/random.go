// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package commonmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockRandomManager is a mock implementation of models.IRandomManager for testing.
type MockRandomManager struct {
	mock.Mock
}

func NewMockRandomManager() *MockRandomManager {
	return &MockRandomManager{}
}

// CreateBigInt calls the mocked function.
func (m *MockRandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	args := m.Called(max)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*big.Int), args.Error(1)
}

var _ managers.IRandomManager = (*MockRandomManager)(nil)
