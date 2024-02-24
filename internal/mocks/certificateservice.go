// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// MockCertificateService is a mock implementation of models.ICertificateService for testing purposes.
type MockCertificateService struct {
	GetExistingCertificateFunc func(
		orgId string,
		signedBy models.ISerialNumber,
		serialNumber models.ISerialNumber) (models.ICertificate, error)
	CreateCertificateFunc func(certificate models.ICertificate) (models.ICertificate, error)
}

var _ models.ICertificateService = (*MockCertificateService)(nil)

func (m *MockCertificateService) GetExistingCertificate(
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber,
) (models.ICertificate, error) {
	return m.GetExistingCertificateFunc(orgId, signedBy, serialNumber)
}

func (m *MockCertificateService) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	return m.CreateCertificateFunc(certificate)
}
