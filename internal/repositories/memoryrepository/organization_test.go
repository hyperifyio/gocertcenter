// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package memoryrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/repositories/memoryrepository"
)

func TestOrganizationRepository_CreateAndGetOrganization(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()
	mockCert := new(mocks.MockOrganization)
	id := "orgId"

	// Setting up expectations
	mockCert.On("GetID").Return(id)

	// Test CreateOrganization
	_, err := repo.CreateOrganization(mockCert)
	assert.NoError(t, err)

	// Test GetExistingOrganization success
	foundCert, err := repo.GetExistingOrganization(id)
	assert.NoError(t, err)
	assert.NotNil(t, foundCert)

	// Verify expectations were met
	mockCert.AssertExpectations(t)
}

func TestOrganizationRepository_GetExistingOrganizationNotFound(t *testing.T) {
	repo := memoryrepository.NewOrganizationRepository()
	id := "testOrg"

	// Test GetExistingOrganization for a non-existent organization
	_, err := repo.GetExistingOrganization(id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization not found")
}
