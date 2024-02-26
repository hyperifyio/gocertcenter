// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateService is a mock implementation of appmodels.ICertificateService for testing purposes.
type MockCertificateService struct {
	mock.Mock
}

func (m *MockCertificateService) GetExistingCertificate(organization string, certificates []appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	args := m.Called(organization, certificates)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockCertificateService) CreateCertificate(certificate appmodels.ICertificate) (appmodels.ICertificate, error) {
	args := m.Called(certificate)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

var _ appmodels.ICertificateService = (*MockCertificateService)(nil)
