// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509/pkix"
	"time"
)

// RevokedCertificateModel model implements RevokedCertificate
type RevokedCertificateModel struct {

	// serialNumber is the serial number of the revoked certificate
	serialNumber SerialNumber

	// revocationTime is the time when the certificate was revoked
	revocationTime time.Time

	// expirationTime is the original expiration time of the certificate
	expirationTime time.Time
}

func (k *RevokedCertificateModel) SerialNumber() SerialNumber {
	return k.serialNumber
}

func (k *RevokedCertificateModel) RevocationTime() time.Time {
	return k.revocationTime
}

func (k *RevokedCertificateModel) ExpirationTime() time.Time {
	return k.expirationTime
}

func (k *RevokedCertificateModel) RevokedCertificate() pkix.RevokedCertificate {
	return pkix.RevokedCertificate{
		SerialNumber:   k.serialNumber.Value(),
		RevocationTime: k.revocationTime,
	}
}

// NewRevokedCertificate creates a private key model from existing data
func NewRevokedCertificate(
	serialNumber SerialNumber,
	revocationTime time.Time,
	expirationTime time.Time,
) *RevokedCertificateModel {
	return &RevokedCertificateModel{
		serialNumber:   serialNumber,
		revocationTime: revocationTime,
		expirationTime: expirationTime,
	}
}

// Compile time assertion for implementing the interface
var _ RevokedCertificate = (*RevokedCertificateModel)(nil)
