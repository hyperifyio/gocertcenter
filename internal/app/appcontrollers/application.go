// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

type ApplicationController struct {
	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	organizationRepository appmodels.IOrganizationService
	certificateRepository  appmodels.ICertificateService
	privateKeyRepository   appmodels.IPrivateKeyService

	// defaultExpiration - Expiration time for new root certificates
	defaultExpiration time.Duration
}

func (a *ApplicationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	return service == a.organizationRepository
}

func (a *ApplicationController) UsesCertificateService(service appmodels.ICertificateService) bool {
	return service == a.certificateRepository
}

func (a *ApplicationController) UsesPrivateKeyService(service appmodels.IPrivateKeyService) bool {
	return service == a.privateKeyRepository
}

func (a *ApplicationController) GetOrganizationModel(organization string) (appmodels.IOrganization, error) {
	model, err := a.organizationRepository.FindById(organization)
	if err != nil {
		return nil, fmt.Errorf("[GetOrganizationModel]: failed: '%s': %w", organization, err)
	}
	return model, nil
}

func (a *ApplicationController) GetOrganizationController(organization string) (appmodels.IOrganizationController, error) {
	model, err := a.GetOrganizationModel(organization)
	if err != nil {
		return nil, fmt.Errorf("[GetOrganizationController:%s]: not found: %w", organization, err)
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
	), nil
}

func (a *ApplicationController) NewOrganization(model appmodels.IOrganization) (appmodels.IOrganization, error) {
	organization := model.GetID()

	_, err := a.organizationRepository.FindById(organization)
	if err == nil {
		return nil, fmt.Errorf("ApplicationController.NewOrganization: organization exist already: %s", organization)
	}

	savedModel, err := a.organizationRepository.Save(model)
	if err != nil {
		return nil, fmt.Errorf("ApplicationController.NewOrganization: could not create organization: %s: %w", organization, err)
	}
	return savedModel, nil
}

func (a *ApplicationController) GetOrganizationCollection() ([]appmodels.IOrganization, error) {
	list, err := a.organizationRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("[GetOrganizationCollection]: failed: %w", err)
	}
	return list, nil
}

// NewApplicationController implements appmodels.IApplicationController
//   - organizationRepository appmodels.IOrganizationService
//   - certificateRepository appmodels.ICertificateService
//   - privateKeyRepository appmodels.IPrivateKeyService
//   - certManager managers.ICertificateManager
//   - randomManager managers.IRandomManager
//   - defaultExpiration time.Duration,
func NewApplicationController(
	organizationRepository appmodels.IOrganizationService,
	certificateRepository appmodels.ICertificateService,
	privateKeyRepository appmodels.IPrivateKeyService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
	defaultExpiration time.Duration,
) *ApplicationController {
	return &ApplicationController{
		organizationRepository: organizationRepository,
		certificateRepository:  certificateRepository,
		privateKeyRepository:   privateKeyRepository,
		certManager:            certManager,
		randomManager:          randomManager,
	}
}

var _ appmodels.IApplicationController = (*ApplicationController)(nil)
