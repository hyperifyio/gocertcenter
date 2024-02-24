// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"time"
)

// Organization model implements IOrganization
type Organization struct {
	id    string
	names []string
}

// Compile time assertion for implementing the interface
var _ IOrganization = (*Organization)(nil)

// NewOrganization creates a organization model from existing data
func NewOrganization(
	id string,
	names []string,
) *Organization {
	return &Organization{
		id:    id,
		names: names,
	}
}

func (o *Organization) GetDTO() dtos.OrganizationDTO {
	return dtos.NewOrganizationDTO(
		o.GetID(),
		o.GetName(),
		o.GetNames(),
	)
}

// GetID returns unique identifier for this organization
func (o *Organization) GetID() string {
	return o.id
}

// GetName returns the primary organization name
func (o *Organization) GetName() string {
	slice := o.GetNames()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

// GetNames returns the full name of the organization including department
func (o *Organization) GetNames() []string {
	return o.names
}

// NewRootCertificate creates a new root CA certificate
func (o *Organization) NewRootCertificate(
	manager ICertificateManager,
	commonName string, // The name of the root CA
	privateKey IPrivateKey, // Private key for self signing
	expiration time.Duration,
) (ICertificate, error) {

	serialNumber := privateKey.GetSerialNumber()

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: o.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	cert, err := privateKey.CreateCertificate(manager, &certificateTemplate, &certificateTemplate)
	if err != nil {
		return nil, err
	}

	return NewCertificate(o.id, serialNumber, cert), nil
}

// NewIntermediateCertificate creates a new intermediate CA certificate
func (o *Organization) NewIntermediateCertificate(
	manager ICertificateManager,
	commonName string, // commonName The name of the intermediate CA
	serialNumber ISerialNumber, // serialNumber Serial Number of the intermediate certificate
	parentCertificate ICertificate, // parentCertificate The parent certificate, typically the root CA
	parentPrivateKey IPrivateKey, // parentPrivateKey Private key of the parent
	expiration time.Duration, // The expiration duration
) (ICertificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: o.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,

		// Restrict this intermediate CA from issuing further intermediate CAs
		MaxPathLenZero: true,
		MaxPathLen:     0,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := parentPrivateKey.CreateCertificate(manager, &certificateTemplate, parentCertificate.GetCertificate())
	if err != nil {
		// FIXME: Add test for this scope
		return nil, err
	}

	return NewCertificate(o.id, parentCertificate.GetSerialNumber(), cert), nil
}

// NewServerCertificate creates a new server certificate
func (o *Organization) NewServerCertificate(
	manager ICertificateManager,
	serialNumber ISerialNumber, // Serial Number of the server certificate
	parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
	privateKey IPrivateKey, // Private key of the parent
	dnsNames []string, // List of domain names the certificate is valid for
	expiration time.Duration,
) (ICertificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: o.GetNames(),
			CommonName:   dnsNames[0],
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              dnsNames,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := privateKey.CreateCertificate(manager, &certificateTemplate, parentCertificate.GetCertificate())
	if err != nil {
		// FIXME: Add test for this scope
		return nil, err
	}

	return NewCertificate(o.id, parentCertificate.GetSerialNumber(), cert), nil
}

// NewClientCertificate creates a new client certificate
func (o *Organization) NewClientCertificate(
	manager ICertificateManager,
	commonName string, // The name of the client
	serialNumber ISerialNumber, // Serial Number of the client certificate
	parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
	privateKey IPrivateKey, // Private key of the parent
	expiration time.Duration,
) (ICertificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: o.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, // Key difference for client auth
		BasicConstraintsValid: true,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := privateKey.CreateCertificate(manager, &certificateTemplate, parentCertificate.GetCertificate())
	if err != nil {
		// FIXME: Add test for this scope
		return nil, err
	}

	return NewCertificate(o.id, parentCertificate.GetSerialNumber(), cert), nil
}
