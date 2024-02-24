// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockCertificate implements a mock models.ICertificate for the ICertificate interface
type MockCertificate struct {
	mock.Mock
}

var _ models.ICertificate = (*MockCertificate)(nil)

func (c *MockCertificate) GetDTO() dtos.CertificateDTO {
	args := c.Called()
	return args.Get(0).(dtos.CertificateDTO)
}

func (m *MockCertificate) IsCA() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) GetSerialNumber() models.ISerialNumber {
	args := m.Called()
	return args.Get(0).(models.ISerialNumber)
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

func (m *MockCertificate) GetSignedBy() models.ISerialNumber {
	args := m.Called()
	return args.Get(0).(models.ISerialNumber)
}

func (m *MockCertificate) GetCertificate() *x509.Certificate {
	args := m.Called()
	return args.Get(0).(*x509.Certificate)
}

func (m *MockCertificate) GetCommonName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificate) IsSelfSigned() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) IsRootCertificate() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) IsIntermediateCertificate() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) IsServerCertificate() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) IsClientCertificate() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) GetPEM() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}
