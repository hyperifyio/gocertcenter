// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

type ApplicationController struct {
	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	organizationRepository appmodels.IOrganizationService
	certificateRepository  appmodels.ICertificateService
	privateKeyRepository   appmodels.IPrivateKeyService
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
	model, err := a.organizationRepository.GetExistingOrganization(organization)
	if err != nil {
		return nil, fmt.Errorf("ApplicationController.GetOrganizationModel: failed to fetch '%s': %w", organization, err)
	}
	return model, nil
}

func (a *ApplicationController) GetOrganizationController(organization string) (appmodels.IOrganizationController, error) {
	model, err := a.GetOrganizationModel(organization)
	if err != nil {
		return nil, fmt.Errorf("ApplicationController.GetOrganizationController: could not find organization '%s': %w", organization, err)
	}
	return NewOrganizationController(organization, model, a.organizationRepository, a.certificateRepository, a.privateKeyRepository, a.certManager, a.randomManager), nil
}

func (a *ApplicationController) NewOrganization(model appmodels.IOrganization) (appmodels.IOrganization, error) {
	organization := model.GetID()
	savedModel, err := a.organizationRepository.CreateOrganization(model)
	if err != nil {
		return nil, fmt.Errorf("ApplicationController.NewOrganization: could not create organization: %s: %w", organization, err)
	}
	return savedModel, nil
}

func NewApplicationController(
	organizationRepository appmodels.IOrganizationService,
	certificateRepository appmodels.ICertificateService,
	privateKeyRepository appmodels.IPrivateKeyService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
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
