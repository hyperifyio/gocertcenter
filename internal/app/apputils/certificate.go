// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CreateSignedCertificate creates a new certificate signed by a root or
// intermediate certificate
//   - manager: CertificateModel manager
//   - template: CertificateModel template
//   - signingCertificate: The certificate to use for signing
//   - signingPublicKey: The public key to use for signing
//   - signingPrivateKey: The private key to use for signing
func CreateSignedCertificate(
	manager managers.CertificateManager,
	template *x509.Certificate,
	signingCertificate *x509.Certificate,
	publicKey any,
	signingPrivateKey any,
) (*x509.Certificate, error) {

	if manager == nil {
		return nil, fmt.Errorf("CreateSignedCertificate: manager: must be defined")
	}

	bytes, err := manager.CreateCertificate(
		rand.Reader,
		template,
		signingCertificate,
		publicKey,
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

func ToCertificateDTO(c appmodels.Certificate) appdtos.CertificateDTO {
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

func ToCertificateRevokedDTO(
	c appmodels.RevokedCertificate,
) appdtos.CertificateRevokedDTO {
	return appdtos.NewCertificateRevokedDTO(
		c.GetSerialNumber().String(),
		c.GetRevocationTime(),
		c.GetExpirationTime(),
	)
}

func ToRevokedCertificate(
	c appmodels.Certificate,
	revocationTime time.Time,
) appmodels.RevokedCertificate {
	return appmodels.NewRevokedCertificate(
		c.GetSerialNumber(),
		revocationTime,
		c.NotAfter(),
	)
}

func ToCertificateCreatedDTO(
	certManager managers.CertificateManager,
	c appmodels.Certificate,
	k appmodels.PrivateKey,
) (appdtos.CertificateCreatedDTO, error) {

	if certManager == nil {
		return appdtos.CertificateCreatedDTO{}, fmt.Errorf("ToCertificateCreatedDTO: cert manager not defined")
	}

	if c == nil {
		return appdtos.CertificateCreatedDTO{}, fmt.Errorf("ToCertificateCreatedDTO: certificate not defined")
	}

	if k == nil {
		return appdtos.CertificateCreatedDTO{}, fmt.Errorf("ToCertificateCreatedDTO: private key not defined")
	}

	dto, err := ToPrivateKeyDTO(certManager, k)
	if err != nil {
		return appdtos.CertificateCreatedDTO{}, fmt.Errorf("ToCertificateCreatedDTO: failed: %w", err)
	}
	return appdtos.NewCertificateCreatedDTO(
		ToCertificateDTO(c),
		dto,
	), nil
}

func ToListOfCertificateDTO(list []appmodels.Certificate) []appdtos.CertificateDTO {
	result := make([]appdtos.CertificateDTO, len(list))
	for i, v := range list {
		result[i] = ToCertificateDTO(v)
	}
	return result
}

func FilterCertificatesByType(list []appmodels.Certificate, certificateType string) []appmodels.Certificate {
	if certificateType == "root" {
		return FilterRootCertificates(list)
	}
	if certificateType == "client" {
		return FilterClientCertificates(list)
	}
	if certificateType == "server" {
		return FilterServerCertificates(list)
	}
	if certificateType == "intermediate" {
		return FilterIntermediateCertificates(list)
	}
	return []appmodels.Certificate{}
}

func FilterRootCertificates(list []appmodels.Certificate) []appmodels.Certificate {
	result := make([]appmodels.Certificate, 0)
	for _, v := range list {
		if v.IsRootCertificate() {
			result = append(result, v)
		}
	}
	return result
}

func FilterClientCertificates(list []appmodels.Certificate) []appmodels.Certificate {
	result := make([]appmodels.Certificate, 0)
	for _, v := range list {
		if v.IsClientCertificate() {
			result = append(result, v)
		}
	}
	return result
}

func FilterServerCertificates(list []appmodels.Certificate) []appmodels.Certificate {
	result := make([]appmodels.Certificate, 0)
	for _, v := range list {
		if v.IsServerCertificate() {
			result = append(result, v)
		}
	}
	return result
}

func FilterIntermediateCertificates(list []appmodels.Certificate) []appmodels.Certificate {
	result := make([]appmodels.Certificate, 0)
	for _, v := range list {
		if v.IsIntermediateCertificate() {
			result = append(result, v)
		}
	}
	return result
}

func ToCertificateListDTO(list []appmodels.Certificate) appdtos.CertificateListDTO {
	payload := ToListOfCertificateDTO(list)
	return appdtos.NewCertificateListDTO(payload)
}

func GetCertificatePEMBytes(c appmodels.Certificate) []byte {
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
//   - manager managers.CertificateManager is the certificate manager
//   - serialNumber appmodels.SerialNumber is the serial number for the new certificate
//   - organization appmodels.Organization is the organization for the new certificate
//   - expiration time.Duration is the expiration duration of the new certificate
//   - publicKey appmodels.PublicKey is public key of the new certificate
//   - parentCertificate appmodels.Certificate is the certificate of the part who signs this certificate
//   - parentPrivateKey appmodels.PrivateKey is the private key of the part who signs this certificate
//   - commonName string is the common name for the new certificate
//
// Returns the new certificate or an error
func NewIntermediateCertificate(
	manager managers.CertificateManager,
	serialNumber appmodels.SerialNumber,
	organization appmodels.Organization,
	expiration time.Duration,
	publicKey appmodels.PublicKey,
	parentCertificate appmodels.Certificate,
	parentPrivateKey appmodels.PrivateKey,
	commonName string,
) (appmodels.Certificate, error) {

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

	if err := ValidateRootCertificateCommonName(commonName); err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: commonName: %s: %s", err, commonName)
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
		ExtKeyUsage:           []x509.ExtKeyUsage{},
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
		publicKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(
		organization.GetID(),
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber()),
		cert,
	), nil

}

// NewServerCertificate creates an intermediate certificate
//   - manager: CertificateModel manager
//   - serialNumber: Serial number for the new certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - publicKey appmodels.PublicKey is public key of the new certificate
//   - parentCertificate: The certificate to use for signing
//   - parentPrivateKey: The private key to use for signing
//   - commonName: The common name for the new certificate
//   - dnsNames: The dns names for the new certificate
//
// Returns the new certificate or an error
func NewServerCertificate(
	manager managers.CertificateManager,
	serialNumber appmodels.SerialNumber,
	organization appmodels.Organization,
	expiration time.Duration,
	publicKey appmodels.PublicKey,
	parentCertificate appmodels.Certificate,
	parentPrivateKey appmodels.PrivateKey,
	commonName string,
	dnsNames ...string,
) (appmodels.Certificate, error) {

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

	if err := ValidateServerCertificateCommonName(commonName); err != nil {
		return nil, fmt.Errorf("NewServerCertificate: commonName: %s: %s", err, commonName)
	}

	if dnsNames == nil || len(dnsNames) <= 0 {
		return nil, fmt.Errorf("NewServerCertificate: dnsNames: must be defined")
	}

	if err := ValidateDNSNames(dnsNames); err != nil {
		return nil, fmt.Errorf("NewServerCertificate: dnsNames: %s: %s", err, strings.Join(dnsNames, " | "))
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
		publicKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewServerCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(
		organization.GetID(),
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber()),
		cert,
	), nil
}

// NewClientCertificate creates an intermediate certificate
//   - manager: CertificateModel manager
//   - serialNumber: Serial number for the new certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - publicKey appmodels.PublicKey is public key of the new certificate
//   - parentCertificate: The certificate to use for signing
//   - parentPrivateKey: The private key to use for signing
//   - commonName: The common name for the new certificate
//   - dnsNames: The dns names for the new certificate
//
// Returns the new certificate or an error
func NewClientCertificate(
	manager managers.CertificateManager,
	serialNumber appmodels.SerialNumber,
	organization appmodels.Organization,
	expiration time.Duration,
	publicKey appmodels.PublicKey,
	parentCertificate appmodels.Certificate,
	parentPrivateKey appmodels.PrivateKey,
	commonName string,
) (appmodels.Certificate, error) {

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

	if err := ValidateClientCertificateCommonName(commonName); err != nil {
		return nil, fmt.Errorf("NewClientCertificate: commonName: %s: %s", err, commonName)
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
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// Use the parent certificate to sign the intermediate certificate
	cert, err := CreateSignedCertificate(
		manager,
		&certificateTemplate,
		parentCertificate.GetCertificate(),
		publicKey.GetPublicKey(),
		parentPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("NewClientCertificate: failed: %w", err)
	}

	return appmodels.NewCertificate(
		organization.GetID(),
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber()),
		cert,
	), nil
}

// NewRootCertificate creates a new root certificate
//   - manager: CertificateModel manager
//   - serialNumber: Serial number for the new root certificate
//   - organization: The organization for the new certificate
//   - expiration: The expiration duration
//   - privateKey: The private key to use for signing
//   - commonName: The common name for the new root certificate
//
// Returns the new certificate or an error
func NewRootCertificate(
	manager managers.CertificateManager,
	serialNumber appmodels.SerialNumber,
	organization appmodels.Organization,
	expiration time.Duration,
	privateKey appmodels.PrivateKey,
	commonName string,
) (appmodels.Certificate, error) {

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

	if err := ValidateRootCertificateCommonName(commonName); err != nil {
		return nil, fmt.Errorf("NewRootCertificate: commonName: %s: %s", err, commonName)
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
		ExtKeyUsage:           []x509.ExtKeyUsage{},
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

	return appmodels.NewCertificate(
		organization.GetID(),
		[]appmodels.SerialNumber{},
		cert,
	), nil
}
