// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"time"
)

// Organization model
type Organization struct {
	id    string
	names []string
}

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

// GetID returns unique identifier for this organization
func (o *Organization) GetID() string {
	return o.id
}

// GetOrganizationName returns the primary organization name
func (o *Organization) GetName() string {
	slice := o.GetNames()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

// GetOrganizationNames returns the full name of the organization including department
func (o *Organization) GetNames() []string {
	return o.names
}

// NewRootCertificate creates a new root CA certificate
func (o *Organization) NewRootCertificate(
	manager managers.ICertificateManager,
	commonName string, // The name of the root CA
	privateKey PrivateKey,
	expiration time.Duration,
) (*Certificate, error) {

	serialNumber := privateKey.GetSerialNumber()

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber,
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
	manager managers.ICertificateManager,
	commonName string, // commonName The name of the intermediate CA
	serialNumber SerialNumber, // serialNumber Serial Number of the intermediate certificate
	parentCertificate *Certificate, // parentCertificate The parent certificate, typically the root CA
	parentPrivateKey PrivateKey, // parentPrivateKey Private key of the parent
	expiration time.Duration, // The expiration duration
) (*Certificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber,
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
	manager managers.ICertificateManager,
	serialNumber SerialNumber, // Serial Number of the server certificate
	parentCertificate *Certificate, // The parent certificate, typically the intermediate or root certificate
	privateKey PrivateKey, // Private key of the parent
	dnsNames []string, // List of domain names the certificate is valid for
	expiration time.Duration,
) (*Certificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber,
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
	manager managers.ICertificateManager,
	commonName string, // The name of the client
	serialNumber SerialNumber, // Serial Number of the client certificate
	parentCertificate *Certificate, // The parent certificate, typically the intermediate or root certificate
	privateKey PrivateKey, // Private key of the parent
	expiration time.Duration,
) (*Certificate, error) {

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber,
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
