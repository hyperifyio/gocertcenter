// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
)

// Certificate model implements ICertificate
type Certificate struct {

	// organization is the organization ID this certificate belongs to
	organization string

	// signedBy is the serial number of the certificate this certificate was signed by
	signedBy ISerialNumber

	// data is the certificate data
	certificate *x509.Certificate
}

// Compile time assertion for implementing the interface
var _ ICertificate = (*Certificate)(nil)

// NewCertificate creates a certificate model from existing data
func NewCertificate(
	organization string,
	signedBy ISerialNumber,
	certificate *x509.Certificate,
) *Certificate {
	return &Certificate{
		organization: organization,
		signedBy:     signedBy,
		certificate:  certificate,
	}
}

func (c *Certificate) GetDTO() dtos.CertificateDTO {
	return dtos.NewCertificateDTO(
		c.GetCommonName(),
		c.GetSerialNumber().String(),
		c.GetSignedBy().String(),
		c.GetOrganizationName(),
		c.IsCA(),
		c.IsRootCertificate(),
		c.IsIntermediateCertificate(),
		c.IsServerCertificate(),
		c.IsClientCertificate(),
		string(c.GetPEM()),
	)
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
	return c.certificate.Subject.Organization
}

func (c *Certificate) GetSignedBy() ISerialNumber {
	return c.signedBy
}

func (c *Certificate) GetCertificate() *x509.Certificate {
	return c.certificate
}

func (c *Certificate) GetPEM() []byte {
	// Convert the certificate to a PEM block
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: c.certificate.Raw,
	}
	// Encode the PEM block to memory
	pemBytes := pem.EncodeToMemory(pemBlock)
	return pemBytes
}
