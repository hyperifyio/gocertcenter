// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"crypto/x509"
	"errors"
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CertificateController implements appmodels.ICertificateController to control
// operations for certificate models.
//
// It utilizes appmodels.ICertificateService implementation to abstract
// and inject the underlying storage mechanism (e.g., database, memory). This
// design promotes separation of concerns by decoupling the business logic from
// the specific details of data persistence.
type CertificateController struct {
	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager
	repository    appmodels.ICertificateService
	expiration    time.Duration
}

func (r *CertificateController) GetApplicationController() appmodels.IApplicationController {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetOrganizationID() string {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetOrganizationModel() appmodels.IOrganization {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetCertificateModel() appmodels.ICertificate {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetChildCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetChildCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetParentCertificateModel() appmodels.ICertificate {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetParentCertificateController() appmodels.ICertificateController {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetPrivateKeyModel() appmodels.IPrivateKey {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) GetPrivateKeyController() appmodels.IPrivateKeyController {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) NewCertificate(template *x509.Certificate) (appmodels.ICertificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) SetExpirationDuration(expiration time.Duration) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateController) NewIntermediateCertificate(commonName string) (appmodels.ICertificate, error) {
	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create serial number: %w", err)
	}
	cert, err := apputils.NewIntermediateCertificate(
		r.certManager,
		serialNumber,
		r.GetOrganizationModel(),
		r.expiration,
		r.GetCertificateModel(),
		r.GetPrivateKeyModel(),
		commonName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create intermediate certificate: %w", err)
	}
	return cert, nil
}

func (r *CertificateController) NewServerCertificate(dnsNames ...string) (appmodels.ICertificate, error) {
	if len(dnsNames) <= 0 {
		return nil, errors.New("server certificate must have at least one dns name")
	}
	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create serial number: %w", err)
	}
	cert, err := apputils.NewServerCertificate(
		r.certManager,
		serialNumber,
		r.GetOrganizationModel(),
		r.expiration,
		r.GetCertificateModel(),
		r.GetPrivateKeyModel(),
		dnsNames[0],
		dnsNames...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create intermediate certificate: %w", err)
	}
	return cert, nil
}

func (r *CertificateController) NewClientCertificate(commonName string) (appmodels.ICertificate, error) {
	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create serial number: %w", err)
	}
	cert, err := apputils.NewServerCertificate(
		r.certManager,
		serialNumber,
		r.GetOrganizationModel(),
		r.expiration,
		r.GetCertificateModel(),
		r.GetPrivateKeyModel(),
		commonName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create intermediate certificate: %w", err)
	}
	return cert, nil
}

func (r *CertificateController) UsesCertificateService(service appmodels.ICertificateService) bool {
	return r.repository == service
}

func (r *CertificateController) GetExistingCertificate(organization string, certificates []appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	return r.repository.GetExistingCertificate(organization, certificates)
}

func (r *CertificateController) CreateSignedCertificate(
	manager managers.ICertificateManager,
	organization string,
	signingCertificate appmodels.ICertificate,
	signingPrivateKey appmodels.IPrivateKey,
	template *x509.Certificate,
) (appmodels.ICertificate, error) {

	certParents := append(signingCertificate.GetParents(), signingCertificate.GetSerialNumber())

	cert, err := apputils.CreateSignedCertificate(
		manager,
		template,
		signingCertificate.GetCertificate(),
		signingPrivateKey.GetPublicKey(),
		signingPrivateKey.GetPrivateKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %w", err)
	}

	return appmodels.NewCertificate(organization, certParents, cert), nil
}

// NewCertificateController creates a new instance of CertificateController
//
// injecting the specified ICertificateService implementation. This setup
// facilitates the separation of business logic from data access layers,
// aligning with the principles of dependency injection.
//
//   - repository is appmodels.ICertificateService
//   - certManager is managers.ICertificateManager
//   - randomManager is  managers.IRandomManager
//   - expiration is time.Duration
func NewCertificateController(
	repository appmodels.ICertificateService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
	expiration time.Duration,
) *CertificateController {
	return &CertificateController{
		repository:    repository,
		expiration:    expiration,
		certManager:   certManager,
		randomManager: randomManager,
	}
}

var _ appmodels.ICertificateController = (*CertificateController)(nil)
