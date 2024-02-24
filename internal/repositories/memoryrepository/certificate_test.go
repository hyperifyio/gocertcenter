// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"

	"github.com/hyperifyio/gocertcenter/internal/repositories/memoryrepository"
)

func TestCertificateRepository_CreateAndGetCertificate(t *testing.T) {
	repo := memoryrepository.NewCertificateRepository()
	mockCert := new(mocks.MockCertificate)
	signedBy := models.NewSerialNumber(big.NewInt(345))
	serialNumber := models.NewSerialNumber(big.NewInt(123))

	// Setting up expectations
	mockCert.On("GetSerialNumber").Return(serialNumber)

	// Test CreateCertificate
	_, err := repo.CreateCertificate(mockCert)
	assert.NoError(t, err)

	// Test GetExistingCertificate success
	foundCert, err := repo.GetExistingCertificate("testOrg", signedBy, serialNumber)
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockCert.AssertExpectations(t)
}

func TestCertificateRepository_GetExistingCertificateNotFound(t *testing.T) {
	repo := memoryrepository.NewCertificateRepository()
	signedBy := models.NewSerialNumber(big.NewInt(111))
	serialNumber := models.NewSerialNumber(big.NewInt(999))

	// Test GetExistingCertificate for a non-existent certificate
	_, err := repo.GetExistingCertificate("testorg", signedBy, serialNumber)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "certificate not found")
}