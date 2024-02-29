// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

type CertApplicationController struct {
	certManager   managers.CertificateManager
	randomManager managers.RandomManager

	organizationRepository appmodels.OrganizationRepository
	certificateRepository  appmodels.CertificateRepository
	privateKeyRepository   appmodels.PrivateKeyRepository

	// defaultExpiration - Expiration time for new root certificates
	defaultExpiration time.Duration
}

func (a *CertApplicationController) UsesOrganizationService(service appmodels.OrganizationRepository) bool {
	return service == a.organizationRepository
}

func (a *CertApplicationController) UsesCertificateService(service appmodels.CertificateRepository) bool {
	return service == a.certificateRepository
}

func (a *CertApplicationController) UsesPrivateKeyService(service appmodels.PrivateKeyRepository) bool {
	return service == a.privateKeyRepository
}

func (a *CertApplicationController) Organization(organization string) (appmodels.Organization, error) {
	model, err := a.organizationRepository.FindById(organization)
	if err != nil {
		return nil, fmt.Errorf("[Organization]: failed: '%s': %w", organization, err)
	}
	return model, nil
}

func (a *CertApplicationController) OrganizationController(organization string) (appmodels.OrganizationController, error) {
	model, err := a.Organization(organization)
	if err != nil {
		return nil, fmt.Errorf("[OrganizationController:%s]: not found: %w", organization, err)
	}
	return NewOrganizationController(
		organization,
		model,
		a.organizationRepository,
		a.certificateRepository,
		a.privateKeyRepository,
		a.certManager,
		a.randomManager,
		a.defaultExpiration,
		a,
	), nil
}

func (a *CertApplicationController) NewOrganization(model appmodels.Organization) (appmodels.Organization, error) {

	if err := apputils.ValidateOrganizationModel(model); err != nil {
		return nil, fmt.Errorf("CertApplicationController.NewOrganization: organization model invalid: %w", err)
	}

	organization := model.ID()

	_, err := a.organizationRepository.FindById(organization)
	if err == nil {
		return nil, fmt.Errorf("CertApplicationController.NewOrganization: organization exist already: %s", organization)
	}

	savedModel, err := a.organizationRepository.Save(model)
	if err != nil {
		return nil, fmt.Errorf("CertApplicationController.NewOrganization: could not create organization: %s: %w", organization, err)
	}
	return savedModel, nil
}

func (a *CertApplicationController) OrganizationCollection() ([]appmodels.Organization, error) {
	list, err := a.organizationRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("[OrganizationCollection]: failed: %w", err)
	}
	return list, nil
}

// NewApplicationController implements appmodels.ApplicationController
//   - organizationRepository appmodels.OrganizationRepository
//   - certificateRepository appmodels.CertificateRepository
//   - privateKeyRepository appmodels.PrivateKeyRepository
//   - certManager managers.CertificateManager
//   - randomManager managers.RandomManager
//   - defaultExpiration time.Duration,
func NewApplicationController(
	organizationRepository appmodels.OrganizationRepository,
	certificateRepository appmodels.CertificateRepository,
	privateKeyRepository appmodels.PrivateKeyRepository,
	certManager managers.CertificateManager,
	randomManager managers.RandomManager,
	defaultExpiration time.Duration,
) *CertApplicationController {
	return &CertApplicationController{
		organizationRepository: organizationRepository,
		certificateRepository:  certificateRepository,
		privateKeyRepository:   privateKeyRepository,
		certManager:            certManager,
		randomManager:          randomManager,
	}
}

var _ appmodels.ApplicationController = (*CertApplicationController)(nil)
