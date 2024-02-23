// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// OrganizationController manages certificate operations.
//
//	It utilizes the IOrganizationService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type OrganizationController struct {
	Service models.IOrganizationService
}

// NewOrganizationController creates a new instance of OrganizationController
//
//	injecting the specified IOrganizationService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewOrganizationController(repository models.IOrganizationService) *OrganizationController {
	return &OrganizationController{
		Service: repository,
	}
}
