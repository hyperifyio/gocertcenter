// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import "github.com/hyperifyio/gocertcenter/internal/models"

// MockCertificateService is a mock implementation of ICertificateService for testing purposes.
type MockCertificateService struct {
	GetExistingCertificateFunc func(serialNumber models.SerialNumber) (*models.Certificate, error)
	CreateCertificateFunc      func(certificate *models.Certificate) (*models.Certificate, error)
}

func (m *MockCertificateService) GetExistingCertificate(serialNumber models.SerialNumber) (*models.Certificate, error) {
	return m.GetExistingCertificateFunc(serialNumber)
}

func (m *MockCertificateService) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	return m.CreateCertificateFunc(certificate)
}
