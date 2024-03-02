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

// TestPrivateKeyRepository_CreateAndGetPrivateKey tests the Save and FindByOrganizationAndSerialNumbers methods
func TestPrivateKeyRepository_CreateAndGetPrivateKey(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewPrivateKeyRepository()
	serialNumber := appmodels.NewSerialNumber(123)

	// Setting up the mock to return the serial number when SerialNumber is called
	mockKey := new(appmocks.MockPrivateKey)
	mockKey.On("OrganizationID").Return(organization)
	mockKey.On("SerialNumber").Return(serialNumber)
	mockKey.On("Parents").Return([]*big.Int{})

	// Test Save
	_, err := repo.Save(mockKey)
	assert.NoError(t, err)

	// Test FindByOrganizationAndSerialNumbers success
	foundKey, err := repo.FindByOrganizationAndSerialNumbers(organization, []*big.Int{serialNumber})
	assert.NoError(t, err)
	assert.NotNil(t, foundKey, "The key should be found")

	// Verify that the mock expectations were met
	mockKey.AssertExpectations(t)
}

// TestPrivateKeyRepository_GetExistingPrivateKeyNotFound tests retrieving a private key that does not exist
func TestPrivateKeyRepository_GetExistingPrivateKeyNotFound(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewPrivateKeyRepository()
	serialNumber := appmodels.NewSerialNumber(456)

	// Test FindByOrganizationAndSerialNumbers for a non-existent key
	_, err := repo.FindByOrganizationAndSerialNumbers(organization, []*big.Int{serialNumber})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ": not found:")
}
