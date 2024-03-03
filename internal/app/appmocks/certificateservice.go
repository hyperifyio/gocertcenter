// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateService is a mock implementation of appmodels.CertificateRepository for testing purposes.
type MockCertificateService struct {
	mock.Mock
}

func (m *MockCertificateService) FindAllByOrganizationAndSignedBy(organization *big.Int, certificate *big.Int) ([]appmodels.Certificate, error) {
	args := m.Called(organization, certificate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateService) FindAllByOrganization(organization *big.Int) ([]appmodels.Certificate, error) {
	args := m.Called(organization)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateService) FindByOrganizationAndSerialNumber(organization *big.Int, certificate *big.Int) (appmodels.Certificate, error) {
	args := m.Called(organization, certificate)
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
