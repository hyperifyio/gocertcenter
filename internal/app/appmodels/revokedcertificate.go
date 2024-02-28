// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509/pkix"
	"time"
)

// RevokedCertificate model implements IRevokedCertificate
type RevokedCertificate struct {

	// serialNumber is the serial number of the revoked certificate
	serialNumber ISerialNumber

	// revocationTime is the time when the certificate was revoked
	revocationTime time.Time

	// expirationTime is the original expiration time of the certificate
	expirationTime time.Time
}

func (k *RevokedCertificate) GetSerialNumber() ISerialNumber {
	return k.serialNumber
}

func (k *RevokedCertificate) GetRevocationTime() time.Time {
	return k.revocationTime
}

func (k *RevokedCertificate) GetExpirationTime() time.Time {
	return k.expirationTime
}

func (k *RevokedCertificate) GetRevokedCertificate() pkix.RevokedCertificate {
	return pkix.RevokedCertificate{
		SerialNumber:   k.serialNumber.Value(),
		RevocationTime: k.revocationTime,
	}
}

// NewRevokedCertificate creates a private key model from existing data
func NewRevokedCertificate(
	serialNumber ISerialNumber,
	revocationTime time.Time,
	expirationTime time.Time,
) *RevokedCertificate {
	return &RevokedCertificate{
		serialNumber:   serialNumber,
		revocationTime: revocationTime,
		expirationTime: expirationTime,
	}
}

// Compile time assertion for implementing the interface
var _ IRevokedCertificate = (*RevokedCertificate)(nil)
