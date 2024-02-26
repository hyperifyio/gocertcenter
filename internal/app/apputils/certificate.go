// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CreateSignedCertificate creates a new certificate signed by a root or
// intermediate certificate
//   - manager: Certificate manager
//   - template: Certificate template
//   - signingCertificate: The certificate to use for signing
//   - signingPublicKey: The public key to use for signing
//   - signingPrivateKey: The private key to use for signing
func CreateSignedCertificate(
	manager managers.ICertificateManager,
	template *x509.Certificate,
	signingCertificate *x509.Certificate,
	signingPublicKey any,
	signingPrivateKey any,
) (*x509.Certificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("CreateSignedCertificate: manager: must be defined")
	}

	bytes, err := manager.CreateCertificate(
		rand.Reader,
		template,
		signingCertificate,
		signingPublicKey,
		signingPrivateKey,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %w", err)
	}

	cert, err := manager.ParseCertificate(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate after creating it: %w", err)
	}

	return cert, nil
}

func GetCertificateDTO(c appmodels.ICertificate) appdtos.CertificateDTO {
	parents := c.GetParents()
	strings := make([]string, len(parents))
	for i, p := range parents {
		strings[i] = p.String()
	}
	return appdtos.NewCertificateDTO(
		c.GetCommonName(),
		c.GetSerialNumber().String(),
		strings,
		c.GetSignedBy().String(),
		c.GetOrganizationName(),
		c.IsCA(),
		c.IsRootCertificate(),
		c.IsIntermediateCertificate(),
		c.IsServerCertificate(),
		c.IsClientCertificate(),
		string(GetCertificatePEMBytes(c)),
	)
}

func GetCertificatePEMBytes(c appmodels.ICertificate) []byte {
	// Convert the certificate to a PEM block
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: c.GetCertificate().Raw,
	}
	// Encode the PEM block to memory
	pemBytes := pem.EncodeToMemory(pemBlock)
	return pemBytes
}

// NewIntermediateCertificate creates an intermediate certificate
//   - manager: Certificate manager
//   - serialNumber: Serial number for the new certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - parentCertificate: The certificate to use for signing
//   - parentPrivateKey: The private key to use for signing
//   - commonName: The common name for the new certificate
//
// Returns the new certificate or an error
func NewIntermediateCertificate(
	manager managers.ICertificateManager,
	serialNumber appmodels.ISerialNumber,
	organization appmodels.IOrganization,
	expiration time.Duration,
	parentCertificate appmodels.ICertificate,
	parentPrivateKey appmodels.IPrivateKey,
	commonName string,
) (appmodels.ICertificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: manager: must be defined")
	}

	if serialNumber == nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: serialNumber: must be defined")
	}

	if organization == nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: organization: must be defined")
	}

	if parentCertificate == nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: parentCertificate: must be defined")
	}

	if parentPrivateKey == nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: parentPrivateKey: must be defined")
	}

	if commonName == "" {
		return nil, fmt.Errorf("NewIntermediateCertificate: commonName: must be defined")
	}

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: organization.GetNames(),
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
	cert, err := CreateSignedCertificate(
		manager,
		&certificateTemplate,
		parentCertificate.GetCertificate(),
		parentPrivateKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(organization.GetID(), append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber()), cert), nil

}

// NewServerCertificate creates an intermediate certificate
//   - manager: Certificate manager
//   - serialNumber: Serial number for the new certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - parentCertificate: The certificate to use for signing
//   - parentPrivateKey: The private key to use for signing
//   - commonName: The common name for the new certificate
//   - dnsNames: The dns names for the new certificate
//
// Returns the new certificate or an error
func NewServerCertificate(
	manager managers.ICertificateManager,
	serialNumber appmodels.ISerialNumber,
	organization appmodels.IOrganization,
	expiration time.Duration,
	parentCertificate appmodels.ICertificate,
	parentPrivateKey appmodels.IPrivateKey,
	commonName string,
	dnsNames ...string,
) (appmodels.ICertificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("NewServerCertificate: manager: must be defined")
	}

	if serialNumber == nil {
		return nil, fmt.Errorf("NewServerCertificate: serialNumber: must be defined")
	}

	if organization == nil {
		return nil, fmt.Errorf("NewServerCertificate: organization: must be defined")
	}

	if parentCertificate == nil {
		return nil, fmt.Errorf("NewServerCertificate: parentCertificate: must be defined")
	}

	if parentPrivateKey == nil {
		return nil, fmt.Errorf("NewServerCertificate: parentPrivateKey: must be defined")
	}

	if commonName == "" {
		return nil, fmt.Errorf("NewServerCertificate: commonName: must be defined")
	}

	if dnsNames == nil || len(dnsNames) <= 0 {
		return nil, fmt.Errorf("NewServerCertificate: dnsNames: must be defined")
	}

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: organization.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              dnsNames,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := CreateSignedCertificate(
		manager,
		&certificateTemplate,
		parentCertificate.GetCertificate(),
		parentPrivateKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewServerCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(organization.GetID(), parentCertificate.GetParents(), cert), nil
}

// NewClientCertificate creates an intermediate certificate
//   - manager: Certificate manager
//   - serialNumber: Serial number for the new certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - parentCertificate: The certificate to use for signing
//   - parentPrivateKey: The private key to use for signing
//   - commonName: The common name for the new certificate
//   - dnsNames: The dns names for the new certificate
//
// Returns the new certificate or an error
func NewClientCertificate(
	manager managers.ICertificateManager,
	serialNumber appmodels.ISerialNumber,
	organization appmodels.IOrganization,
	expiration time.Duration,
	parentCertificate appmodels.ICertificate,
	parentPrivateKey appmodels.IPrivateKey,
	commonName string,
) (appmodels.ICertificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("NewClientCertificate: manager: must be defined")
	}

	if serialNumber == nil {
		return nil, fmt.Errorf("NewClientCertificate: serialNumber: must be defined")
	}

	if organization == nil {
		return nil, fmt.Errorf("NewClientCertificate: organization: must be defined")
	}

	if parentCertificate == nil {
		return nil, fmt.Errorf("NewClientCertificate: parentCertificate: must be defined")
	}

	if parentPrivateKey == nil {
		return nil, fmt.Errorf("NewClientCertificate: parentPrivateKey: must be defined")
	}

	if commonName == "" {
		return nil, fmt.Errorf("NewClientCertificate: commonName: must be defined")
	}

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: organization.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, // Key difference for client auth
		BasicConstraintsValid: true,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := CreateSignedCertificate(
		manager,
		&certificateTemplate,
		parentCertificate.GetCertificate(),
		parentPrivateKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewClientCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(organization.GetID(), parentCertificate.GetParents(), cert), nil
}

// NewRootCertificate creates a new root certificate
//   - manager: Certificate manager
//   - serialNumber: Serial number for the new root certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - privateKey: The private key to use for signing
//   - commonName: The common name for the new root certificate
//
// Returns the new certificate or an error
func NewRootCertificate(
	manager managers.ICertificateManager,
	serialNumber appmodels.ISerialNumber,
	organization appmodels.IOrganization,
	expiration time.Duration,
	privateKey appmodels.IPrivateKey,
	commonName string,
) (appmodels.ICertificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("NewRootCertificate: manager: must be defined")
	}

	if serialNumber == nil {
		return nil, fmt.Errorf("NewRootCertificate: serialNumber: must be defined")
	}

	if organization == nil {
		return nil, fmt.Errorf("NewRootCertificate: organization: must be defined")
	}

	if privateKey == nil {
		return nil, fmt.Errorf("NewRootCertificate: privateKey: must be defined")
	}

	if commonName == "" {
		return nil, fmt.Errorf("NewRootCertificate: commonName: must be defined")
	}

	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber.Value(),
		Subject: pkix.Name{
			Organization: organization.GetNames(),
			CommonName:   commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := CreateSignedCertificate(
		manager,
		&certificateTemplate,
		&certificateTemplate,
		privateKey.GetPublicKey(),
		privateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewRootCertificate: failed to create certificate: %w", err)
	}

	return appmodels.NewCertificate(organization.GetID(), []appmodels.ISerialNumber{}, cert), nil
}
