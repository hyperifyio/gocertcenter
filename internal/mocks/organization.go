// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockOrganization is a mock implementation of the IOrganization interface
type MockOrganization struct {
	mock.Mock
}

var _ models.IOrganization = (*MockOrganization)(nil)

func (m *MockOrganization) GetDTO() dtos.OrganizationDTO {
	args := m.Called()
	return args.Get(0).(dtos.OrganizationDTO)
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
