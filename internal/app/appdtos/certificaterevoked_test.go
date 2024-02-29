// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewCertificateRevokedDTO(t *testing.T) {
	// Define test data for the CertificateRevokedDTO fields
	serialNumber := "123456789"
	revocationTime := time.Now()
	expirationTime := revocationTime.Add(365 * 24 * time.Hour) // 1 year from now

	// Create a CertificateRevokedDTO instance using the constructor
	certificateRevokedDTO := appdtos.NewCertificateRevokedDTO(serialNumber, revocationTime, expirationTime)

	// Assert that the fields are correctly assigned
	assert.Equal(t, serialNumber, certificateRevokedDTO.SerialNumber, "SerialNumber should match the input serial number")
	assert.Equal(t, revocationTime, certificateRevokedDTO.RevocationTime, "RevocationTime should match the input revocation time")
	assert.Equal(t, expirationTime, certificateRevokedDTO.ExpirationTime, "ExpirationTime should match the input expiration time")
}
