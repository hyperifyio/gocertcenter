// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"bytes"
	"crypto/x509"
	"math/big"
	"time"
)

// CertificateModel model implements Certificate
type CertificateModel struct {

	// organization is the organization ID this certificate belongs to
	organization string

	// signedBy is the serial number of the root/intermediate certificate which signed this one
	signedBy *big.Int

	// data is the certificate data
	certificate *x509.Certificate
}

// NewCertificate creates a certificate model from existing data
func NewCertificate(
	organization string,
	signedBy *big.Int,
	certificate *x509.Certificate,
) *CertificateModel {
	return &CertificateModel{
		organization: organization,
		signedBy:     signedBy,
		certificate:  certificate,
	}
}

func (c *CertificateModel) NotBefore() time.Time {
	return c.certificate.NotBefore
}

func (c *CertificateModel) NotAfter() time.Time {
	return c.certificate.NotAfter
}

func (c *CertificateModel) IsCA() bool {
	return c.certificate.IsCA
}

func (c *CertificateModel) IsSelfSigned() bool {
	if len(c.certificate.AuthorityKeyId) > 0 && len(c.certificate.SubjectKeyId) > 0 {
		return bytes.Equal(c.certificate.AuthorityKeyId, c.certificate.SubjectKeyId)
	}
	return c.certificate.Issuer.String() == c.certificate.Subject.String()
}

func (c *CertificateModel) IsIntermediateCertificate() bool {
	return c.certificate.BasicConstraintsValid && c.certificate.IsCA && !c.IsSelfSigned()
}

func (c *CertificateModel) IsServerCertificate() bool {
	for _, usage := range c.certificate.ExtKeyUsage {
		if usage == x509.ExtKeyUsageServerAuth {
			return true
		}
	}
	return false
}

func (c *CertificateModel) IsClientCertificate() bool {
	for _, usage := range c.certificate.ExtKeyUsage {
		if usage == x509.ExtKeyUsageClientAuth {
			return true
		}
	}
	return false
}

func (c *CertificateModel) IsRootCertificate() bool {
	return c.certificate.BasicConstraintsValid && c.certificate.IsCA && c.IsSelfSigned()
}

func (c *CertificateModel) SerialNumber() *big.Int {
	return c.certificate.SerialNumber
}

func (c *CertificateModel) OrganizationID() string {
	return c.organization
}

func (c *CertificateModel) CommonName() string {
	return c.certificate.Subject.CommonName
}

func (c *CertificateModel) OrganizationName() string {
	slice := c.Organization()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

func (c *CertificateModel) Organization() []string {
	originalSlice := c.certificate.Subject.Organization
	sliceCopy := make([]string, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (c *CertificateModel) SignedBy() *big.Int {
	return c.signedBy
}

func (c *CertificateModel) Certificate() *x509.Certificate {
	return c.certificate
}

// Compile time assertion for implementing the interface
var _ Certificate = (*CertificateModel)(nil)
