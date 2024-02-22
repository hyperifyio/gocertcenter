// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package controllers

import (
	models2 "github.com/hyperifyio/gocertcenter/internal/models"
)

// ICertificateService defines the interface for certificate storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type ICertificateService interface {
	GetExistingCertificate(serialNumber models2.SerialNumber) (*models2.Certificate, error)
	CreateCertificate(certificate *models2.Certificate) (*models2.Certificate, error)
}

// CertificateController manages certificate operations.
//
//	It utilizes the ICertificateService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type CertificateController struct {
	service ICertificateService
}

// NewCertificateController creates a new instance of CertificateController
//
//	injecting the specified ICertificateService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewCertificateController(repository ICertificateService) *CertificateController {
	return &CertificateController{
		service: repository,
	}
}
