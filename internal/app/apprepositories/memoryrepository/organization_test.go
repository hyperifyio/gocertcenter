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
	mockOrg.On("GetID").Return(id)

	// Test CreateOrganization
	_, err := repo.CreateOrganization(mockOrg)
	assert.NoError(t, err)

	// Test GetExistingOrganization success
	foundCert, err := repo.GetExistingOrganization(id)
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockOrg.AssertExpectations(t)
}

func TestOrganizationRepository_GetExistingOrganizationNotFound(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()
	id := "testOrg"

	// Test GetExistingOrganization for a non-existent organization
	_, err := repo.GetExistingOrganization(id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization not found")
}
