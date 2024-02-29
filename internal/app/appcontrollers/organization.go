// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

const (
	DefaultRootKeyType = appmodels.ECDSA_P384
)

// OrganizationController implements models.IOrganizationController to control
// operations for organization models.
//
// It utilizes models.IOrganizationService interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type OrganizationController struct {

	// id is the organization ID this controller controls
	id string

	// model is the latest model of the organization
	model appmodels.IOrganization

	parent appmodels.IApplicationController

	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	organizationRepository appmodels.IOrganizationService
	certificateRepository  appmodels.ICertificateService
	privateKeyRepository   appmodels.IPrivateKeyService

	// defaultExpiration - Expiration time for new root certificates
	defaultExpiration time.Duration

	// defaultKeyType - The default key type for root certificates
	defaultKeyType appmodels.KeyType
}

func (r *OrganizationController) GetCertificateCollection() ([]appmodels.ICertificate, error) {
	organization := r.GetOrganizationID()
	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s:GetCertificateCollection]: no certificate repository", organization)
	}
	list, err := r.certificateRepository.FindAllByOrganizationAndSerialNumbers(organization, []appmodels.ISerialNumber{})
	if err != nil {
		return nil, fmt.Errorf("[%s:GetCertificateCollection]: failed: %w", organization, err)
	}
	return list, nil
}

func (r *OrganizationController) GetOrganizationID() string {
	return r.id
}

func (r *OrganizationController) GetOrganizationModel() appmodels.IOrganization {
	return r.model
}

func (r *OrganizationController) GetApplicationController() appmodels.IApplicationController {
	return r.parent
}

func (r *OrganizationController) GetCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	model, err := r.GetCertificateModel(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s:GetCertificateController:%s]: failed: %w", r.id, serialNumber, err)
	}
	return NewCertificateController(
		r,
		nil,
		serialNumber,
		model,
		r.certificateRepository,
		r.privateKeyRepository,
		r.certManager,
		r.randomManager,
		r.defaultExpiration,
	), nil
}

func (r *OrganizationController) GetCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	model, err := r.certificateRepository.FindByOrganizationAndSerialNumbers(r.id, []appmodels.ISerialNumber{serialNumber})
	if err != nil {
		return nil, fmt.Errorf("[%s:GetCertificateModel:%s]: failed: %w", r.id, serialNumber.String(), err)
	}
	return model, nil
}

func (r *OrganizationController) SetExpirationDuration(expiration time.Duration) {
	r.defaultExpiration = expiration
}

func (r *OrganizationController) ExpirationDuration() time.Duration {
	return r.defaultExpiration
}

func (r *OrganizationController) NewRootCertificate(commonName string) (appmodels.ICertificate, error) {

	organization := r.GetOrganizationID()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: failed to create serial number: %w", organization, commonName, err)
	}

	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: no certificate repository", organization, commonName)
	}

	if r.privateKeyRepository == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: no certificate repository", organization, commonName)
	}

	_, err = r.certificateRepository.FindByOrganizationAndSerialNumbers(organization, []appmodels.ISerialNumber{serialNumber})
	if err == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: serial number exists already: %s", organization, commonName, serialNumber.String())
	}

	keyType := r.defaultKeyType
	if keyType == appmodels.NIL_KEY_TYPE {
		keyType = DefaultRootKeyType
	}

	privateKey, err := apputils.GeneratePrivateKey(
		r.GetOrganizationID(),
		[]appmodels.ISerialNumber{serialNumber},
		keyType,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: failed to generate private key: %w", organization, commonName, err)
	}

	cert, err := apputils.NewRootCertificate(
		r.certManager,
		serialNumber,
		r.GetOrganizationModel(),
		r.defaultExpiration,
		privateKey,
		commonName,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: failed to create certificate: %w", organization, commonName, err)
	}

	_, err = r.privateKeyRepository.Save(privateKey)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: could not save private key: %w", organization, commonName, err)
	}

	savedModel, err := r.certificateRepository.Save(cert)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: could not save certificate: %w", organization, commonName, err)
	}

	return savedModel, nil
}

func (r *OrganizationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	return r.organizationRepository == service
}

func (r *OrganizationController) UsesApplicationController(service appmodels.IApplicationController) bool {
	return r.parent == service
}

func (r *OrganizationController) RevokeCertificate(certificate appmodels.ICertificate) (appmodels.IRevokedCertificate, error) {
	// TODO implement me
	panic("implement me")
}

// NewOrganizationController creates a new instance of OrganizationController
// implementing appmodels.IOrganizationService interface.
//
//   - organization string
//   - model appmodels.IOrganization
//   - organizationRepository appmodels.IOrganizationService
//   - certificateRepository appmodels.ICertificateService
//   - privateKeyRepository appmodels.IPrivateKeyService
//   - certManager managers.ICertificateManager
//   - randomManager managers.IRandomManager
//   - defaultExpiration time.Duration
//
// Returns *OrganizationController
func NewOrganizationController(
	organization string,
	model appmodels.IOrganization,
	organizationRepository appmodels.IOrganizationService,
	certificateRepository appmodels.ICertificateService,
	privateKeyRepository appmodels.IPrivateKeyService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
	defaultExpiration time.Duration,
	parent appmodels.IApplicationController,
) *OrganizationController {
	return &OrganizationController{
		id:                     organization,
		model:                  model,
		organizationRepository: organizationRepository,
		certificateRepository:  certificateRepository,
		privateKeyRepository:   privateKeyRepository,
		certManager:            certManager,
		randomManager:          randomManager,
		defaultExpiration:      defaultExpiration,
		parent:                 parent,
	}
}

var _ appmodels.IOrganizationController = (*OrganizationController)(nil)
