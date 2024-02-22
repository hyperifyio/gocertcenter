// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import "math/big"

// MockRandomManager is a mock implementation of IRandomManager for testing.
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
