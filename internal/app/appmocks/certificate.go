// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"crypto/x509"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificate implements a mock models.Certificate for the Certificate interface
type MockCertificate struct {
	mock.Mock
}

func (c *MockCertificate) NotBefore() time.Time {
	args := c.Called()
	return args.Get(0).(time.Time)
}

func (c *MockCertificate) NotAfter() time.Time {
	args := c.Called()
	return args.Get(0).(time.Time)
}

func (m *MockCertificate) IsCA() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockCertificate) SerialNumber() appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).(appmodels.SerialNumber)
}

func (m *MockCertificate) OrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificate) OrganizationName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificate) Organization() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m *MockCertificate) Parents() []appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).([]appmodels.SerialNumber)
}

func (m *MockCertificate) SignedBy() appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).(appmodels.SerialNumber)
}

func (m *MockCertificate) Certificate() *x509.Certificate {
	args := m.Called()
	return args.Get(0).(*x509.Certificate)
}

func (m *MockCertificate) CommonName() string {
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

var _ appmodels.Certificate = (*MockCertificate)(nil)
