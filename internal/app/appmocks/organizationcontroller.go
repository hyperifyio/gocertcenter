// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockOrganizationController is a mock implementation of models.OrganizationController for testing purposes.
type MockOrganizationController struct {
	mock.Mock
}

func (m *MockOrganizationController) RevokeCertificate(certificate appmodels.Certificate) (appmodels.RevokedCertificate, error) {
	args := m.Called(certificate)
	return args.Get(0).(appmodels.RevokedCertificate), args.Error(1)
}

func (m *MockOrganizationController) CertificateCollection() ([]appmodels.Certificate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.Certificate), args.Error(1)
}

func (m *MockOrganizationController) OrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganizationController) Organization() appmodels.Organization {
	args := m.Called()
	return args.Get(0).(appmodels.Organization)
}

func (m *MockOrganizationController) ApplicationController() appmodels.ApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.ApplicationController)
}

func (m *MockOrganizationController) CertificateController(serialNumber appmodels.SerialNumber) (appmodels.CertificateController, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.CertificateController), args.Error(1)
}

func (m *MockOrganizationController) Certificate(serialNumber appmodels.SerialNumber) (appmodels.Certificate, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

func (m *MockOrganizationController) SetExpirationDuration(expiration time.Duration) {
	m.Called(expiration)
}

func (m *MockOrganizationController) NewRootCertificate(commonName string) (appmodels.Certificate, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

func (m *MockOrganizationController) UsesOrganizationService(service appmodels.OrganizationRepository) bool {
	args := m.Called(service)
	return args.Bool(0)
}

func (m *MockOrganizationController) UsesApplicationController(service appmodels.ApplicationController) bool {
	args := m.Called(service)
	return args.Bool(0)
}

var _ appmodels.OrganizationController = (*MockOrganizationController)(nil)
