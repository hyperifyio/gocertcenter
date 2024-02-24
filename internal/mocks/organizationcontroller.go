// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
	"time"
)

// MockOrganizationController is a mock implementation of models.IOrganizationController for testing purposes.
type MockOrganizationController struct {
	mock.Mock
}

var _ models.IOrganizationController = (*MockOrganizationController)(nil)

func (m *MockOrganizationController) UsesOrganizationService(service models.IOrganizationService) bool {
	args := m.Called(service)
	return args.Bool(0)
}

func (m *MockOrganizationController) GetExistingOrganization(id string) (models.IOrganization, error) {
	args := m.Called(id)
	return args.Get(0).(models.IOrganization), args.Error(1)
}

func (m *MockOrganizationController) CreateOrganization(organization models.IOrganization) (models.IOrganization, error) {
	args := m.Called(organization)
	return args.Get(0).(models.IOrganization), args.Error(1)
}

func (m *MockOrganizationController) NewRootCertificate(o models.IOrganization, manager models.ICertificateManager, commonName string, privateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(o, manager, commonName, privateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) NewIntermediateCertificate(o models.IOrganization, manager models.ICertificateManager, commonName string, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, parentPrivateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(o, manager, commonName, serialNumber, parentCertificate, parentPrivateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) NewServerCertificate(o models.IOrganization, manager models.ICertificateManager, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, privateKey models.IPrivateKey, dnsNames []string, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(o, manager, serialNumber, parentCertificate, privateKey, dnsNames, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) NewClientCertificate(o models.IOrganization, manager models.ICertificateManager, commonName string, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, privateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(o, manager, commonName, serialNumber, parentCertificate, privateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}
