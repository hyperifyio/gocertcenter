// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockOrganization is a mock implementation of the IOrganization interface
type MockOrganization struct {
	mock.Mock
}

func (m *MockOrganization) GetID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganization) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganization) GetNames() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

var _ appmodels.IOrganization = (*MockOrganization)(nil)
