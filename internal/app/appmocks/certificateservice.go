// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateService is a mock implementation of appmodels.CertificateRepository for testing purposes.
type MockCertificateService struct {
	mock.Mock
}

func (m *MockCertificateService) FindAllByOrganizationAndSerialNumbers(organization string, certificates []appmodels.SerialNumber) ([]appmodels.Certificate, error) {
	args := m.Called(organization, certificates)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateService) FindAllByOrganization(organization string) ([]appmodels.Certificate, error) {
	args := m.Called(organization)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateService) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.SerialNumber) (appmodels.Certificate, error) {
	args := m.Called(organization, certificates)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateService) Save(certificate appmodels.Certificate) (appmodels.Certificate, error) {
	args := m.Called(certificate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

var _ appmodels.CertificateRepository = (*MockCertificateService)(nil)
