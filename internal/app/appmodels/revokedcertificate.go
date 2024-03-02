// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509/pkix"
	"math/big"
	"time"
)

// RevokedCertificateModel model implements RevokedCertificate
type RevokedCertificateModel struct {

	// serialNumber is the serial number of the revoked certificate
	serialNumber *big.Int

	// revocationTime is the time when the certificate was revoked
	revocationTime time.Time

	// expirationTime is the original expiration time of the certificate
	expirationTime time.Time
}

func (k *RevokedCertificateModel) SerialNumber() *big.Int {
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
		SerialNumber:   k.serialNumber,
		RevocationTime: k.revocationTime,
	}
}

// NewRevokedCertificate creates a private key model from existing data
func NewRevokedCertificate(
	serialNumber *big.Int,
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
