// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmocks

import (
	"crypto/x509/pkix"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockRevokedCertificate is a mock type for the IRevokedCertificate interface
type MockRevokedCertificate struct {
	mock.Mock
}

// GetSerialNumber mocks the GetSerialNumber method
func (m *MockRevokedCertificate) GetSerialNumber() appmodels.ISerialNumber {
	args := m.Called()
	return args.Get(0).(appmodels.ISerialNumber) // Ensure the return type matches the interface
}

// GetRevocationTime mocks the GetRevocationTime method
func (m *MockRevokedCertificate) GetRevocationTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time) // Ensure the return type matches the interface
}

// GetExpirationTime mocks the GetExpirationTime method
func (m *MockRevokedCertificate) GetExpirationTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time) // Ensure the return type matches the interface
}

// GetRevokedCertificate mocks the GetRevokedCertificate method
func (m *MockRevokedCertificate) GetRevokedCertificate() pkix.RevokedCertificate {
	args := m.Called()
	return args.Get(0).(pkix.RevokedCertificate) // Ensure the return type matches the interface
}
