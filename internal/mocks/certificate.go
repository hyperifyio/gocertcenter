// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockCertificate implements a mock models.ICertificate for the ICertificate interface
type MockCertificate struct {
	mock.Mock
}

var _ models.ICertificate = (*MockCertificate)(nil)

func (m *MockCertificate) IsCA() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) GetSerialNumber() models.SerialNumber {
	args := m.Called()
	return args.Get(0).(models.SerialNumber)
}

func (m *MockCertificate) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificate) GetOrganizationName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificate) GetOrganization() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m *MockCertificate) GetSignedBy() models.SerialNumber {
	args := m.Called()
	return args.Get(0).(models.SerialNumber)
}

func (m *MockCertificate) GetCertificate() *x509.Certificate {
	args := m.Called()
	return args.Get(0).(*x509.Certificate)
}
