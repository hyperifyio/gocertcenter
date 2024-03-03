// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockOrganization is a mock implementation of the Organization interface
type MockOrganization struct {
	mock.Mock
}

func (m *MockOrganization) ID() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockOrganization) Slug() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganization) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganization) Names() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

var _ appmodels.Organization = (*MockOrganization)(nil)
