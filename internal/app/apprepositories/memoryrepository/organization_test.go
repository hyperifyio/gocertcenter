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
