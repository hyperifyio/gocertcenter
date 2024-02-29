// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmocks

import (
	"crypto/x509/pkix"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockRevokedCertificate is a mock type for the RevokedCertificate interface
type MockRevokedCertificate struct {
	mock.Mock
}

// GetSerialNumber mocks the GetSerialNumber method
func (m *MockRevokedCertificate) SerialNumber() appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).(appmodels.SerialNumber) // Ensure the return type matches the interface
}

// GetRevocationTime mocks the GetRevocationTime method
func (m *MockRevokedCertificate) RevocationTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time) // Ensure the return type matches the interface
}

// GetExpirationTime mocks the GetExpirationTime method
func (m *MockRevokedCertificate) ExpirationTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time) // Ensure the return type matches the interface
}

// GetRevokedCertificate mocks the GetRevokedCertificate method
func (m *MockRevokedCertificate) RevokedCertificate() pkix.RevokedCertificate {
	args := m.Called()
	return args.Get(0).(pkix.RevokedCertificate) // Ensure the return type matches the interface
}
