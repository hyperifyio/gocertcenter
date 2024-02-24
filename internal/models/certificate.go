// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
)

// Certificate model implements ICertificate
type Certificate struct {

	// organization is the organization ID this certificate belongs to
	organization string

	// signedBy is the serial number of the certificate this certificate was signed by
	signedBy SerialNumber

	// data is the certificate data
	certificate *x509.Certificate
}

// Compile time assertion for implementing the interface
var _ ICertificate = (*Certificate)(nil)

// NewCertificate creates a certificate model from existing data
func NewCertificate(
	organization string,
	signedBy SerialNumber,
	certificate *x509.Certificate,
) *Certificate {
	return &Certificate{
		organization: organization,
		signedBy:     signedBy,
		certificate:  certificate,
	}
}

func (c *Certificate) IsCA() bool {
	return c.certificate.IsCA
}

func (c *Certificate) GetSerialNumber() SerialNumber {
	return c.certificate.SerialNumber
}

func (c *Certificate) GetOrganizationID() string {
	return c.organization
}

func (c *Certificate) GetOrganizationName() string {
	slice := c.GetOrganization()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

func (c *Certificate) GetOrganization() []string {
	return c.certificate.Subject.Organization
}

func (c *Certificate) GetSignedBy() SerialNumber {
	return c.signedBy
}

func (c *Certificate) GetCertificate() *x509.Certificate {
	return c.certificate
}
