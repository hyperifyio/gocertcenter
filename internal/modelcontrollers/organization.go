// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

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

// NewOrganizationController creates a new instance of OrganizationController
//
//	injecting the specified IOrganizationService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewOrganizationController(
	service models.IOrganizationService,
) *OrganizationController {
	return &OrganizationController{
		repository: service,
	}
}
