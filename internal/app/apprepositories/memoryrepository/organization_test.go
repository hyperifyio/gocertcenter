// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/memoryrepository"
)

func TestOrganizationRepository_CreateAndGetOrganization(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()
	mockOrg := new(appmocks.MockOrganization)
	id := "orgId"

	// Setting up expectations
	mockOrg.On("ID").Return(id)

	// Test Save
	_, err := repo.Save(mockOrg)
	assert.NoError(t, err)

	// Test FindById success
	foundCert, err := repo.FindById(id)
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockOrg.AssertExpectations(t)
}

func TestOrganizationRepository_GetExistingOrganizationNotFound(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()
	id := "testOrg"

	// Test FindById for a non-existent organization
	_, err := repo.FindById(id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ": not found:")
}

func TestOrganizationRepository_FindAll(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()

	// Mock organizations
	mockOrg1 := new(appmocks.MockOrganization)
	mockOrg2 := new(appmocks.MockOrganization)

	id1 := "orgId1"
	id2 := "orgId2"

	// Setting up expectations for the mock organizations
	mockOrg1.On("ID").Return(id1)
	mockOrg2.On("ID").Return(id2)

	// Save mock organizations to the repository
	_, err1 := repo.Save(mockOrg1)
	assert.NoError(t, err1)

	_, err2 := repo.Save(mockOrg2)
	assert.NoError(t, err2)

	// Test FindAll
	organizations, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, organizations, 2, "Expected to find 2 organizations")

	// Verify that returned organizations match saved ones
	// Since map iteration order is not guaranteed, use a map to verify existence
	foundIds := make(map[string]bool)
	for _, org := range organizations {
		foundIds[org.ID()] = true
	}

	assert.True(t, foundIds[id1], "Expected to find organization with ID %s", id1)
	assert.True(t, foundIds[id2], "Expected to find organization with ID %s", id2)

	// Verify expectations were met for the mock organizations
	mockOrg1.AssertExpectations(t)
	mockOrg2.AssertExpectations(t)
}
