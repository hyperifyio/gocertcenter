// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// CertificateController implements models.ICertificateController to control
// operations for certificate models.
//
// It utilizes models.ICertificateService implementation to abstract
// and inject the underlying storage mechanism (e.g., database, memory). This
// design promotes separation of concerns by decoupling the business logic from
// the specific details of data persistence.
type CertificateController struct {
	repository models.ICertificateService
}

var _ models.ICertificateController = (*CertificateController)(nil)

func (r *CertificateController) UsesCertificateService(service models.ICertificateService) bool {
	return r.repository == service
}

func (r *CertificateController) GetExistingCertificate(
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber) (models.ICertificate, error) {
	return r.repository.GetExistingCertificate(orgId, signedBy, serialNumber)
}

func (r *CertificateController) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	return r.repository.CreateCertificate(certificate)
}

// NewCertificateController creates a new instance of CertificateController
//
//	injecting the specified ICertificateService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewCertificateController(repository models.ICertificateService) *CertificateController {
	return &CertificateController{
		repository: repository,
	}
}
