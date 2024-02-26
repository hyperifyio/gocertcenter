// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"bytes"
	"crypto/x509"
)

// Certificate model implements ICertificate
type Certificate struct {

	// organization is the organization ID this certificate belongs to
	organization string

	// parents is all parent certificates in the chain
	parents []ISerialNumber

	// data is the certificate data
	certificate *x509.Certificate
}

// Compile time assertion for implementing the interface
var _ ICertificate = (*Certificate)(nil)

// NewCertificate creates a certificate model from existing data
func NewCertificate(
	organization string,
	parents []ISerialNumber,
	certificate *x509.Certificate,
) *Certificate {
	return &Certificate{
		organization: organization,
		parents:      parents,
		certificate:  certificate,
	}
}

func (c *Certificate) IsCA() bool {
	return c.certificate.IsCA
}

func (c *Certificate) IsSelfSigned() bool {
	if len(c.certificate.AuthorityKeyId) > 0 && len(c.certificate.SubjectKeyId) > 0 {
		return bytes.Equal(c.certificate.AuthorityKeyId, c.certificate.SubjectKeyId)
	}
	return c.certificate.Issuer.String() == c.certificate.Subject.String()
}

func (c *Certificate) IsIntermediateCertificate() bool {
	return c.certificate.BasicConstraintsValid && c.certificate.IsCA && !c.IsSelfSigned()
}

func (c *Certificate) IsServerCertificate() bool {
	for _, usage := range c.certificate.ExtKeyUsage {
		if usage == x509.ExtKeyUsageServerAuth {
			return true
		}
	}
	return false
}

func (c *Certificate) IsClientCertificate() bool {
	for _, usage := range c.certificate.ExtKeyUsage {
		if usage == x509.ExtKeyUsageClientAuth {
			return true
		}
	}
	return false
}

func (c *Certificate) IsRootCertificate() bool {
	return c.certificate.BasicConstraintsValid && c.certificate.IsCA && c.IsSelfSigned()
}

func (c *Certificate) GetSerialNumber() ISerialNumber {
	return NewSerialNumber(c.certificate.SerialNumber)
}

func (c *Certificate) GetOrganizationID() string {
	return c.organization
}

func (c *Certificate) GetParents() []ISerialNumber {
	originalSlice := c.parents
	sliceCopy := make([]ISerialNumber, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (c *Certificate) GetCommonName() string {
	return c.certificate.Subject.CommonName
}

func (c *Certificate) GetOrganizationName() string {
	slice := c.GetOrganization()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

func (c *Certificate) GetOrganization() []string {
	originalSlice := c.certificate.Subject.Organization
	sliceCopy := make([]string, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (c *Certificate) GetSignedBy() ISerialNumber {
	if len(c.parents) >= 1 {
		return c.parents[len(c.parents)-1]
	} else {
		return c.GetSerialNumber()
	}
}

func (c *Certificate) GetCertificate() *x509.Certificate {
	return c.certificate
}
