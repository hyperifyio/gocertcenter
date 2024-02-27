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
	serialNumber appmodels.ISerialNumber
	model        appmodels.ICertificate

	// parentOrganizationController is the parent organization controller
	parentOrganizationController appmodels.IOrganizationController

	// parentCertificateController  is an optional parent certificate controller
	parentCertificateController appmodels.ICertificateController

	privateKeyController appmodels.IPrivateKeyController

	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	certificateRepository appmodels.ICertificateService
	privateKeyRepository  appmodels.IPrivateKeyService

	expiration time.Duration
}

func (r *CertificateController) GetOrganizationController() appmodels.IOrganizationController {
	return r.parentOrganizationController
}

func (r *CertificateController) GetApplicationController() appmodels.IApplicationController {
	return r.parentOrganizationController.GetApplicationController()
}

func (r *CertificateController) GetOrganizationID() string {
	return r.parentOrganizationController.GetOrganizationID()
}

func (r *CertificateController) GetOrganizationModel() appmodels.IOrganization {
	return r.parentOrganizationController.GetOrganizationModel()
}

func (r *CertificateController) GetCertificateModel() appmodels.ICertificate {
	return r.model
}

func (r *CertificateController) GetChildCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	model, err := r.certificateRepository.GetExistingCertificate(
		r.GetOrganizationID(),
		append(r.model.GetParents(), r.serialNumber, serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("CertificateController('%s').GetChildCertificateModel('%s'): failed to fetch: %w", r.serialNumber, serialNumber.String(), err)
	}
	return model, nil
}

func (r *CertificateController) GetChildCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	model, err := r.GetChildCertificateModel(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("CertificateController('%s').GetChildCertificateController('%s'): could not find: %w", r.serialNumber, serialNumber, err)
	}
	return NewCertificateController(
		serialNumber,
		model,
		r.certificateRepository,
		r.privateKeyRepository,
		r.certManager,
		r.randomManager,
		r.expiration,
	), nil
}

func (r *CertificateController) GetParentCertificateModel() appmodels.ICertificate {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController.GetCertificateModel()
}

func (r *CertificateController) GetParentCertificateController() appmodels.ICertificateController {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController
}

func (r *CertificateController) GetPrivateKeyModel() (appmodels.IPrivateKey, error) {
	model, err := r.privateKeyRepository.GetExistingPrivateKey(
		r.GetOrganizationID(),
		append(r.model.GetParents(), r.serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("CertificateController('%s').GetPrivateKeyModel(): failed to fetch: %w", r.serialNumber, err)
	}
	return model, nil
}

func (r *CertificateController) GetPrivateKeyController() (appmodels.IPrivateKeyController, error) {
	model, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, fmt.Errorf("CertificateController('%s').GetPrivateKeyController(): could not find: %w", r.serialNumber, err)
	}
	return NewPrivateKeyController(
		model,
		r,
		r.privateKeyRepository,
	), nil
}

func (r *CertificateController) SetExpirationDuration(expiration time.Duration) {
	r.expiration = expiration
}

func (r *CertificateController) NewIntermediateCertificate(commonName string) (appmodels.ICertificate, error) {

	privateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: failed to fetch private key: %w", err)
	}

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: failed to create serial number: %w", err)
	}

	cert, err := apputils.NewIntermediateCertificate(
		r.certManager,
		serialNumber,
		r.GetOrganizationModel(),
		r.expiration,
		r.GetCertificateModel(),
		privateKey,
		commonName,
	)
	if err != nil {
		return nil, fmt.Errorf("NewIntermediateCertificate: failed to create intermediate certificate: %w", err)
	}

	return cert, nil
}

func (r *CertificateController) NewServerCertificate(dnsNames ...string) (appmodels.ICertificate, error) {
	if len(dnsNames) <= 0 {
		return nil, errors.New("server certificate must have at least one dns name")
	}

	privateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, fmt.Errorf("NewServerCertificate: failed to fetch private key: %w", err)
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
		privateKey,
		dnsNames[0],
		dnsNames...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create intermediate certificate: %w", err)
	}
	return cert, nil
}

func (r *CertificateController) NewClientCertificate(commonName string) (appmodels.ICertificate, error) {
	privateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, fmt.Errorf("NewServerCertificate: failed to fetch private key: %w", err)
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
		privateKey,
		commonName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create intermediate certificate: %w", err)
	}
	return cert, nil
}

func (r *CertificateController) UsesCertificateService(service appmodels.ICertificateService) bool {
	return r.certificateRepository == service
}

func (r *CertificateController) GetExistingCertificate(organization string, certificates []appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	return r.certificateRepository.GetExistingCertificate(organization, certificates)
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
//   - serialNumber appmodels.ISerialNumber
//   - model appmodels.ICertificate
//   - certificateRepository is appmodels.ICertificateService
//   - privateKeyRepository is appmodels.IPrivateKeyService
//   - certManager is managers.ICertificateManager
//   - randomManager is  managers.IRandomManager
//   - expiration time.Duration is
func NewCertificateController(
	serialNumber appmodels.ISerialNumber,
	model appmodels.ICertificate,
	certificateRepository appmodels.ICertificateService,
	privateKeyRepository appmodels.IPrivateKeyService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
	expiration time.Duration,
) *CertificateController {
	return &CertificateController{
		serialNumber:          serialNumber,
		model:                 model,
		certificateRepository: certificateRepository,
		privateKeyRepository:  privateKeyRepository,
		expiration:            expiration,
		certManager:           certManager,
		randomManager:         randomManager,
	}
}

var _ appmodels.ICertificateController = (*CertificateController)(nil)
