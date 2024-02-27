// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository_test

import (
	"math/big"
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
	signedBy := appmodels.NewSerialNumber(big.NewInt(345))
	serialNumber := appmodels.NewSerialNumber(big.NewInt(123))

	// Setting up expectations
	mockCert.On("GetSerialNumber").Return(serialNumber)
	mockCert.On("GetOrganizationID").Return(organization)
	mockCert.On("GetParents").Return([]appmodels.ISerialNumber{signedBy})

	// Test Save
	_, err := repo.Save(mockCert)
	assert.NoError(t, err)

	// Test FindByOrganizationAndSerialNumbers success
	foundCert, err := repo.FindByOrganizationAndSerialNumbers(organization, []appmodels.ISerialNumber{signedBy, serialNumber})
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockCert.AssertExpectations(t)
}

func TestCertificateRepository_GetExistingCertificateNotFound(t *testing.T) {
	repo := memoryrepository.NewCertificateRepository()
	signedBy := appmodels.NewSerialNumber(big.NewInt(111))
	serialNumber := appmodels.NewSerialNumber(big.NewInt(999))

	// Test FindByOrganizationAndSerialNumbers for a non-existent certificate
	_, err := repo.FindByOrganizationAndSerialNumbers("testorg", []appmodels.ISerialNumber{signedBy, serialNumber})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ": not found:")
}
