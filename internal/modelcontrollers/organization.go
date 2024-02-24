// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"time"
)

// OrganizationController implements models.IOrganizationController to control
// operations for organization models.
//
// It utilizes models.IOrganizationService interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type OrganizationController struct {
	repository models.IOrganizationService
}

var _ models.IOrganizationController = (*OrganizationController)(nil)

func (r *OrganizationController) UsesOrganizationService(service models.IOrganizationService) bool {
	return r.repository == service
}

func (r *OrganizationController) GetExistingOrganization(id string) (models.IOrganization, error) {
	return r.repository.GetExistingOrganization(id)
}

func (r *OrganizationController) CreateOrganization(certificate models.IOrganization) (models.IOrganization, error) {
	return r.repository.CreateOrganization(certificate)
}

// NewRootCertificate creates a new root CA certificate
func (r *OrganizationController) NewRootCertificate(
	o models.IOrganization,
	manager models.ICertificateManager,
	commonName string, // The name of the root CA
	privateKey models.IPrivateKey, // Private key for self signing
	expiration time.Duration,
) (models.ICertificate, error) {

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

	return models.NewCertificate(o.GetID(), serialNumber, cert), nil
}

// NewIntermediateCertificate creates a new intermediate CA certificate
func (r *OrganizationController) NewIntermediateCertificate(
	o models.IOrganization,
	manager models.ICertificateManager,
	commonName string, // commonName The name of the intermediate CA
	serialNumber models.ISerialNumber, // serialNumber Serial Number of the intermediate certificate
	parentCertificate models.ICertificate, // parentCertificate The parent certificate, typically the root CA
	parentPrivateKey models.IPrivateKey, // parentPrivateKey Private key of the parent
	expiration time.Duration, // The expiration duration
) (models.ICertificate, error) {

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

	return models.NewCertificate(o.GetID(), parentCertificate.GetSerialNumber(), cert), nil
}

// NewServerCertificate creates a new server certificate
func (r *OrganizationController) NewServerCertificate(
	o models.IOrganization,
	manager models.ICertificateManager,
	serialNumber models.ISerialNumber, // Serial Number of the server certificate
	parentCertificate models.ICertificate, // The parent certificate, typically the intermediate or root certificate
	privateKey models.IPrivateKey, // Private key of the parent
	dnsNames []string, // List of domain names the certificate is valid for
	expiration time.Duration,
) (models.ICertificate, error) {

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

	return models.NewCertificate(o.GetID(), parentCertificate.GetSerialNumber(), cert), nil
}

// NewClientCertificate creates a new client certificate
func (r *OrganizationController) NewClientCertificate(
	o models.IOrganization,
	manager models.ICertificateManager,
	commonName string, // The name of the client
	serialNumber models.ISerialNumber, // Serial Number of the client certificate
	parentCertificate models.ICertificate, // The parent certificate, typically the intermediate or root certificate
	privateKey models.IPrivateKey, // Private key of the parent
	expiration time.Duration,
) (models.ICertificate, error) {

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

	return models.NewCertificate(o.GetID(), parentCertificate.GetSerialNumber(), cert), nil
}

// NewOrganizationController creates a new instance of OrganizationController
//
//	injecting the specified models.IOrganizationService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewOrganizationController(
	service models.IOrganizationService,
) *OrganizationController {
	return &OrganizationController{
		repository: service,
	}
}
