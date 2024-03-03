// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockApplicationController is a mock type for the ApplicationController interface
type MockApplicationController struct {
	mock.Mock
}

func (m *MockApplicationController) OrganizationCollection() ([]appmodels.Organization, error) {
	// TODO implement me
	panic("implement me")
}

// UsesOrganizationService mocks the UsesOrganizationService method
func (m *MockApplicationController) UsesOrganizationService(service appmodels.OrganizationRepository) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// UsesCertificateService mocks the UsesCertificateService method
func (m *MockApplicationController) UsesCertificateService(service appmodels.CertificateRepository) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// UsesPrivateKeyService mocks the UsesPrivateKeyService method
func (m *MockApplicationController) UsesPrivateKeyService(service appmodels.PrivateKeyRepository) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// Organization mocks the Organization method
func (m *MockApplicationController) Organization(organization *big.Int) (appmodels.Organization, error) {
	args := m.Called(organization)
	return args.Get(0).(appmodels.Organization), args.Error(1)
}

// OrganizationController mocks the OrganizationController method
func (m *MockApplicationController) OrganizationController(name *big.Int) (appmodels.OrganizationController, error) {
	args := m.Called(name)
	return args.Get(0).(appmodels.OrganizationController), args.Error(1)
}

// NewOrganization mocks the NewOrganization method
func (m *MockApplicationController) NewOrganization(certificate appmodels.Organization) (appmodels.Organization, error) {
	args := m.Called(certificate)
	return args.Get(0).(appmodels.Organization), args.Error(1)
}

var _ appmodels.ApplicationController = (*MockApplicationController)(nil)
