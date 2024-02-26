// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// OrganizationController implements models.IOrganizationController to control
// operations for organization models.
//
// It utilizes models.IOrganizationService interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type OrganizationController struct {
	repository    appmodels.IOrganizationService
	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	// defaultExpiration - Expiration time for new root certificates
	defaultExpiration time.Duration

	// defaultKeyType - The default key type for root certificates
	defaultKeyType appmodels.KeyType
}

func (r *OrganizationController) GetOrganizationID() string {
	// TODO implement me
	panic("implement me")
}

func (r *OrganizationController) GetOrganizationModel() appmodels.IOrganization {
	// TODO implement me
	panic("implement me")
}

func (r *OrganizationController) GetApplicationController() appmodels.IApplicationController {
	// TODO implement me
	panic("implement me")
}

func (r *OrganizationController) GetCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	// TODO implement me
	panic("implement me")
}

func (r *OrganizationController) GetCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *OrganizationController) SetExpirationDuration(expiration time.Duration) {
	r.defaultExpiration = expiration
}

func (r *OrganizationController) NewRootCertificate(commonName string) (appmodels.ICertificate, error) {

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create serial number: %w", err)
	}

	privateKey, err := apputils.GeneratePrivateKey(
		r.GetOrganizationID(),
		[]appmodels.ISerialNumber{serialNumber},
		r.defaultKeyType,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
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
		return nil, fmt.Errorf("failed to create root certificate: %w", err)
	}
	return cert, nil
}

func (r *OrganizationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	return r.repository == service
}

func (r *OrganizationController) GetExistingOrganization(id string) (appmodels.IOrganization, error) {
	return r.repository.GetExistingOrganization(id)
}

func (r *OrganizationController) CreateOrganization(certificate appmodels.IOrganization) (appmodels.IOrganization, error) {
	return r.repository.CreateOrganization(certificate)
}

// NewOrganizationController creates a new instance of OrganizationController
//
//	injecting the specified models.IOrganizationService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewOrganizationController(
	service appmodels.IOrganizationService,
) *OrganizationController {
	return &OrganizationController{
		repository: service,
	}
}

var _ appmodels.IOrganizationController = (*OrganizationController)(nil)
