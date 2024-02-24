// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
	"time"
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

func (m *MockOrganization) NewRootCertificate(manager models.ICertificateManager, commonName string, privateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(manager, commonName, privateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganization) NewIntermediateCertificate(manager models.ICertificateManager, commonName string, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, parentPrivateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(manager, commonName, serialNumber, parentCertificate, parentPrivateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganization) NewServerCertificate(manager models.ICertificateManager, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, privateKey models.IPrivateKey, dnsNames []string, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(manager, serialNumber, parentCertificate, privateKey, dnsNames, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}

func (m *MockOrganization) NewClientCertificate(manager models.ICertificateManager, commonName string, serialNumber models.ISerialNumber, parentCertificate models.ICertificate, privateKey models.IPrivateKey, expiration time.Duration) (models.ICertificate, error) {
	args := m.Called(manager, commonName, serialNumber, parentCertificate, privateKey, expiration)
	return args.Get(0).(models.ICertificate), args.Error(1)
}
