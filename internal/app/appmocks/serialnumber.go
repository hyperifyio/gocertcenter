// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

type MockSerialNumber struct {
	mock.Mock
}

func (m *MockSerialNumber) String() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockSerialNumber) Value() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockSerialNumber) Cmp(s2 appmodels.ISerialNumber) int {
	args := m.Called(s2)
	return args.Int(0)
}

func (m *MockSerialNumber) Sign() int {
	args := m.Called()
	return args.Int(0)
}

// NewMockSerialNumber creates an instance of MockSerialNumber. This function can be used to instantiate the mock and set up expectations or return values.
func NewMockSerialNumber() *MockSerialNumber {
	return &MockSerialNumber{}
}

var _ appmodels.ISerialNumber = (*MockSerialNumber)(nil)
