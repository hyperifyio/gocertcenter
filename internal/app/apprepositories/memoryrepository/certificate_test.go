// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/memoryrepository"
)

func TestCertificateRepository_CreateAndGetCertificate(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewCertificateRepository()
	mockCert := new(appmocks.MockCertificate)
	serialNumber := appmodels.NewSerialNumber(123)

	// Setting up expectations
	mockCert.On("SerialNumber").Return(serialNumber)
	mockCert.On("OrganizationID").Return(organization)

	// Test Save
	_, err := repo.Save(mockCert)
	assert.NoError(t, err)

	// Test FindByOrganizationAndSerialNumber success
	foundCert, err := repo.FindByOrganizationAndSerialNumber(organization, serialNumber)
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockCert.AssertExpectations(t)
}

func TestCertificateRepository_GetExistingCertificateNotFound(t *testing.T) {
	repo := memoryrepository.NewCertificateRepository()
	serialNumber := appmodels.NewSerialNumber(999)

	// Test FindByOrganizationAndSerialNumbers for a non-existent certificate
	_, err := repo.FindByOrganizationAndSerialNumber("testorg", serialNumber)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ": not found:")
}

func TestCertificateRepository_FindAllByOrganizationAndSignedBy(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewCertificateRepository()
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	signedBy1 := appmodels.NewSerialNumber(1)
	signedBy2 := appmodels.NewSerialNumber(2)
	serialNumber1 := appmodels.NewSerialNumber(123)
	serialNumber2 := appmodels.NewSerialNumber(456)

	// Setting up expectations
	mockCert1.On("SerialNumber").Return(serialNumber1)
	mockCert1.On("OrganizationID").Return(organization)
	mockCert1.On("SignedBy").Return(signedBy1)

	mockCert2.On("SerialNumber").Return(serialNumber2)
	mockCert2.On("OrganizationID").Return(organization)
	mockCert2.On("SignedBy").Return(signedBy2)

	// Test Save
	_, err1 := repo.Save(mockCert1)
	assert.NoError(t, err1)

	_, err2 := repo.Save(mockCert2)
	assert.NoError(t, err2)

	// Test FindAllByOrganizationAndSerialNumbers
	foundCerts, err := repo.FindAllByOrganizationAndSignedBy(organization, signedBy1)
	assert.NoError(t, err)
	assert.Len(t, foundCerts, 1, "Expected to find 1 certificates")

	// Verify expectations were met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
}

func TestCertificateRepository_FindAllByOrganization(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewCertificateRepository()
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	serialNumber1 := appmodels.NewSerialNumber(123)
	serialNumber2 := appmodels.NewSerialNumber(456)

	// Setting up expectations
	mockCert1.On("SerialNumber").Return(serialNumber1)
	mockCert1.On("OrganizationID").Return(organization)

	mockCert2.On("SerialNumber").Return(serialNumber2)
	mockCert2.On("OrganizationID").Return(organization)

	// Test Save
	_, err1 := repo.Save(mockCert1)
	assert.NoError(t, err1)

	_, err2 := repo.Save(mockCert2)
	assert.NoError(t, err2)

	// Test FindAllByOrganization
	foundCerts, err := repo.FindAllByOrganization(organization)
	assert.NoError(t, err)
	assert.Len(t, foundCerts, 2, "Expected to find 2 certificates")

	// Verify expectations were met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
}

func TestCertificateRepository_FindByOrganizationAndSerialNumber_WithNoCertificates(t *testing.T) {
	organization := "testOrg"

	// Indirectly setting certificates to nil to test the specific case
	repo := memoryrepository.MemoryCertificateRepository{}

	// Mock serial numbers array to pass as parameter
	serialNumber := appmodels.NewSerialNumber(456)

	// Call the method with the repository having nil certificates
	foundCerts, err := repo.FindAllByOrganizationAndSignedBy(organization, serialNumber)

	// Validate that no error is returned and the result is an empty slice
	assert.NoError(t, err, "Expected no error when certificates map is nil")
	assert.Nil(t, foundCerts, "Expected result to be nil")
	assert.Empty(t, foundCerts, "Expected result to be an empty slice when certificates map is nil")
}

func TestCertificateRepository_FindAllByOrganization_WithNilCertificates(t *testing.T) {

	// Manually setting the certificates map to nil to simulate the uninitialized scenario
	repo := memoryrepository.MemoryCertificateRepository{}

	organization := "testOrg"

	// Attempt to retrieve certificates by organization
	_, err := repo.FindAllByOrganization(organization)

	// Assert that the function returns the expected error
	assert.Error(t, err, "[Certificate:FindAllByOrganization]: not initialized")
	assert.Contains(t, err.Error(), "[Certificate:FindAllByOrganization]: not initialized", "Error message should indicate that the repository is not initialized")
}
