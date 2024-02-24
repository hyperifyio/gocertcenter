// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// MockCertificateController is a mock implementation of models.ICertificateController for testing purposes.
type MockCertificateController struct {
	UsesCertificateServiceFunc func(service models.ICertificateService) bool
	GetExistingCertificateFunc func(
		orgId string,
		signedBy models.ISerialNumber,
		serialNumber models.ISerialNumber) (models.ICertificate, error)
	CreateCertificateFunc func(certificate models.ICertificate) (models.ICertificate, error)
}

var _ models.ICertificateController = (*MockCertificateController)(nil)

func (m *MockCertificateController) UsesCertificateService(service models.ICertificateService) bool {
	return m.UsesCertificateServiceFunc(service)
}

func (m *MockCertificateController) GetExistingCertificate(
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber) (models.ICertificate, error) {
	return m.GetExistingCertificateFunc(orgId, signedBy, serialNumber)
}

func (m *MockCertificateController) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	return m.CreateCertificateFunc(certificate)
}
