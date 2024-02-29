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

// CertOrganizationController implements models.OrganizationController to control
// operations for organization models.
//
// It utilizes models.OrganizationRepository interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type CertOrganizationController struct {

	// id is the organization ID this controller controls
	id string

	// model is the latest model of the organization
	model appmodels.Organization

	parent appmodels.ApplicationController

	certManager   managers.CertificateManager
	randomManager managers.RandomManager

	organizationRepository appmodels.OrganizationRepository
	certificateRepository  appmodels.CertificateRepository
	privateKeyRepository   appmodels.PrivateKeyRepository

	// defaultExpiration - Expiration time for new root certificates
	defaultExpiration time.Duration

	// defaultKeyType - The default key type for root certificates
	defaultKeyType appmodels.KeyType
}

func (r *CertOrganizationController) CertificateCollection() ([]appmodels.Certificate, error) {
	organization := r.OrganizationID()
	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s:CertificateCollection]: no certificate repository", organization)
	}
	list, err := r.certificateRepository.FindAllByOrganizationAndSerialNumbers(organization, []appmodels.SerialNumber{})
	if err != nil {
		return nil, fmt.Errorf("[%s:CertificateCollection]: failed: %w", organization, err)
	}
	return list, nil
}

func (r *CertOrganizationController) OrganizationID() string {
	return r.id
}

func (r *CertOrganizationController) Organization() appmodels.Organization {
	return r.model
}

func (r *CertOrganizationController) ApplicationController() appmodels.ApplicationController {
	return r.parent
}

func (r *CertOrganizationController) CertificateController(serialNumber appmodels.SerialNumber) (appmodels.CertificateController, error) {
	model, err := r.Certificate(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s:CertificateController:%s]: failed: %w", r.id, serialNumber, err)
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

func (r *CertOrganizationController) Certificate(serialNumber appmodels.SerialNumber) (appmodels.Certificate, error) {
	model, err := r.certificateRepository.FindByOrganizationAndSerialNumbers(r.id, []appmodels.SerialNumber{serialNumber})
	if err != nil {
		return nil, fmt.Errorf("[%s:Certificate:%s]: failed: %w", r.id, serialNumber.String(), err)
	}
	return model, nil
}

func (r *CertOrganizationController) SetExpirationDuration(expiration time.Duration) {
	r.defaultExpiration = expiration
}

func (r *CertOrganizationController) ExpirationDuration() time.Duration {
	return r.defaultExpiration
}

func (r *CertOrganizationController) NewRootCertificate(commonName string) (appmodels.Certificate, error) {

	organization := r.OrganizationID()

	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: no certificate repository", organization, commonName)
	}

	if r.privateKeyRepository == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: no certificate repository", organization, commonName)
	}

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: failed to create serial number: %w", organization, commonName, err)
	}

	_, err = r.certificateRepository.FindByOrganizationAndSerialNumbers(organization, []appmodels.SerialNumber{serialNumber})
	if err == nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: serial number exists already: %s", organization, commonName, serialNumber.String())
	}

	keyType := r.defaultKeyType
	if keyType == appmodels.NIL_KEY_TYPE {
		keyType = DefaultRootKeyType
	}

	privateKey, err := apputils.GeneratePrivateKey(
		organization,
		[]appmodels.SerialNumber{serialNumber},
		keyType,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s:NewRootCertificate:%s]: failed to generate private key: %w", organization, commonName, err)
	}

	cert, err := apputils.NewRootCertificate(
		r.certManager,
		serialNumber,
		r.Organization(),
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

func (r *CertOrganizationController) UsesOrganizationService(service appmodels.OrganizationRepository) bool {
	return r.organizationRepository == service
}

func (r *CertOrganizationController) UsesApplicationController(service appmodels.ApplicationController) bool {
	return r.parent == service
}

func (r *CertOrganizationController) RevokeCertificate(certificate appmodels.Certificate) (appmodels.RevokedCertificate, error) {
	// TODO implement me
	panic("implement me")
}

// NewOrganizationController creates a new instance of CertOrganizationController
// implementing appmodels.OrganizationRepository interface.
//
//   - organization string
//   - model appmodels.Organization
//   - organizationRepository appmodels.OrganizationRepository
//   - certificateRepository appmodels.CertificateRepository
//   - privateKeyRepository appmodels.PrivateKeyRepository
//   - certManager managers.CertificateManager
//   - randomManager managers.RandomManager
//   - defaultExpiration time.Duration
//
// Returns *CertOrganizationController
func NewOrganizationController(
	organization string,
	model appmodels.Organization,
	organizationRepository appmodels.OrganizationRepository,
	certificateRepository appmodels.CertificateRepository,
	privateKeyRepository appmodels.PrivateKeyRepository,
	certManager managers.CertificateManager,
	randomManager managers.RandomManager,
	defaultExpiration time.Duration,
	parent appmodels.ApplicationController,
) *CertOrganizationController {
	return &CertOrganizationController{
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

var _ appmodels.OrganizationController = (*CertOrganizationController)(nil)
