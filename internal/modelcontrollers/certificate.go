// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// CertificateController manages certificate operations.
//
//	It utilizes the ICertificateService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type CertificateController struct {
	service models.ICertificateService
}

// NewCertificateController creates a new instance of CertificateController
//
//	injecting the specified ICertificateService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewCertificateController(repository models.ICertificateService) *CertificateController {
	return &CertificateController{
		service: repository,
	}
}
