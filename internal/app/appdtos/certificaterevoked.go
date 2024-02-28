// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

import (
	"time"
)

type CertificateRevokedDTO struct {

	// SerialNumber is the serial number of revoked certificate
	SerialNumber string `json:"serialNumber"`

	// RevocationTime is the time when the certificate was revoked
	RevocationTime time.Time `json:"revocationTime"`

	// ExpirationTime is the original expiration time of the certificate
	ExpirationTime time.Time `json:"expirationTime"`
}

func NewCertificateRevokedDTO(
	serialNumber string,
	revocationTime time.Time,
	expirationTime time.Time,
) CertificateRevokedDTO {
	return CertificateRevokedDTO{
		SerialNumber:   serialNumber,
		RevocationTime: revocationTime,
		ExpirationTime: expirationTime,
	}
}
