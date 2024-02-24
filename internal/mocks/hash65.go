// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"github.com/stretchr/testify/mock"
	"hash"
)

// MockHash64 is a mock type for hash.Hash64 interface
type MockHash64 struct {
	mock.Mock
}

var _ hash.Hash64 = (*MockHash64)(nil)

// Write is the mock implementation of the Write method from the hash.Hash interface
func (m *MockHash64) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

// Sum is the mock implementation of the Sum method from the hash.Hash interface
func (m *MockHash64) Sum(b []byte) []byte {
	args := m.Called(b)
	return args.Get(0).([]byte)
}

// Reset is the mock implementation of the Reset method from the hash.Hash interface
func (m *MockHash64) Reset() {
	m.Called()
}

// Size is the mock implementation of the Size method from the hash.Hash interface
func (m *MockHash64) Size() int {
	args := m.Called()
	return args.Int(0)
}

// BlockSize is the mock implementation of the BlockSize method from the hash.Hash interface
func (m *MockHash64) BlockSize() int {
	args := m.Called()
	return args.Int(0)
}

// Sum64 is the mock implementation of the Sum64 method from the Hash64 interface
func (m *MockHash64) Sum64() uint64 {
	args := m.Called()
	return args.Get(0).(uint64)
}
