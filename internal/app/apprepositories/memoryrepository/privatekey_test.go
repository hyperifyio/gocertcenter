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

// TestPrivateKeyRepository_CreateAndGetPrivateKey tests the CreatePrivateKey and GetExistingPrivateKey methods
func TestPrivateKeyRepository_CreateAndGetPrivateKey(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewPrivateKeyRepository()
	serialNumber := appmodels.NewSerialNumber(big.NewInt(123))

	// Setting up the mock to return the serial number when GetSerialNumber is called
	mockKey := new(appmocks.MockPrivateKey)
	mockKey.On("GetOrganizationID").Return(organization)
	mockKey.On("GetSerialNumber").Return(serialNumber)
	mockKey.On("GetParents").Return([]appmodels.ISerialNumber{})

	// Test CreatePrivateKey
	_, err := repo.CreatePrivateKey(mockKey)
	assert.NoError(t, err)

	// Test GetExistingPrivateKey success
	foundKey, err := repo.GetExistingPrivateKey(organization, []appmodels.ISerialNumber{serialNumber})
	assert.NoError(t, err)
	assert.NotNil(t, foundKey, "The key should be found")

	// Verify that the mock expectations were met
	mockKey.AssertExpectations(t)
}

// TestPrivateKeyRepository_GetExistingPrivateKeyNotFound tests retrieving a private key that does not exist
func TestPrivateKeyRepository_GetExistingPrivateKeyNotFound(t *testing.T) {
	organization := "testOrg"
	repo := memoryrepository.NewPrivateKeyRepository()
	serialNumber := appmodels.NewSerialNumber(big.NewInt(456))

	// Test GetExistingPrivateKey for a non-existent key
	_, err := repo.GetExistingPrivateKey(organization, []appmodels.ISerialNumber{serialNumber})
	assert.Error(t, err)
	assert.Contains(t, "key not found", err.Error())
}
