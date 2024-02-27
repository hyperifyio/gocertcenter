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
	list, err := r.certificateRepository.FindAllByOrganization(organization)
	if err != nil {
		return nil, fmt.Errorf("OrganizationController(%s).GetCertificateCollection: failed: %w", organization, err)
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
		return nil, fmt.Errorf("OrganizationController('%s').GetCertificateController('%s'): could not find: %w", r.id, serialNumber, err)
	}
	return NewCertificateController(
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
		return nil, fmt.Errorf("OrganizationController('%s').GetCertificateModel('%s'): failed to fetch: %w", r.id, serialNumber.String(), err)
	}
	return model, nil
}

func (r *OrganizationController) SetExpirationDuration(expiration time.Duration) {
	r.defaultExpiration = expiration
}

func (r *OrganizationController) NewRootCertificate(commonName string) (appmodels.ICertificate, error) {

	organization := r.GetOrganizationID()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: failed to create serial number: %w", organization, err)
	}

	_, err = r.certificateRepository.FindByOrganizationAndSerialNumbers(organization, []appmodels.ISerialNumber{serialNumber})
	if err == nil {
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: serial number exist already: %s", organization, serialNumber.String())
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
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: failed to generate private key: %w", organization, err)
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
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: failed to create certificate: %w", organization, err)
	}

	_, err = r.privateKeyRepository.Save(privateKey)
	if err != nil {
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: could not save private key: %w", organization, err)
	}

	savedModel, err := r.certificateRepository.Save(cert)
	if err != nil {
		return nil, fmt.Errorf("[OrganizationController(%s).NewRootCertificate]: could not save certificate: %w", organization, err)
	}

	return savedModel, nil
}

func (r *OrganizationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	return r.organizationRepository == service
}

func (r *OrganizationController) UsesApplicationController(service appmodels.IApplicationController) bool {
	return r.parent == service
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
	}
}

var _ appmodels.IOrganizationController = (*OrganizationController)(nil)
