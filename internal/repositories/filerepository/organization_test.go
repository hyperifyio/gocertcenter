// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package filerepository_test

import (
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"github.com/hyperifyio/gocertcenter/internal/repositories/filerepository"
)

func TestOrganizationRepository_GetExistingOrganization(t *testing.T) {
	// Setup
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	filePath := tempDir
	orgID := "org123"
	repo := filerepository.NewOrganizationRepository(filePath)

	orgJsonPath := filerepository.GetOrganizationJsonPath(filePath, orgID)

	err := filerepository.SaveOrganizationJsonFile(
		orgJsonPath,
		dtos.NewOrganizationDTO(orgID, "Test Org", []string{"Test Org"}),
	)
	assert.NoError(t, err)

	// Test
	org, err := repo.GetExistingOrganization(orgID)
	assert.NoError(t, err)
	assert.NotNil(t, org)

	// Perform more assertions based on the expected DTO to be returned
	assert.Equal(t, orgID, org.GetID(), "The organization ID should match the requested ID")
	assert.Equal(t, "Test Org", org.GetName(), "The organization name should match the saved name")
	expectedNames := []string{"Test Org"} // This should match the names used in SaveOrganizationJsonFile
	assert.Equal(t, expectedNames, org.GetNames(), "The organization names should match the saved names")

}

func TestOrganizationRepository_CreateOrganization(t *testing.T) {

	// Setup
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	filePath := tempDir
	orgID := "org123"
	mockOrg := &mocks.MockOrganization{}
	mockOrg.On("GetID").Return(orgID)
	mockOrg.On("GetDTO").Return(dtos.OrganizationDTO{ID: orgID, AllNames: []string{"Test Org"}})
	repo := filerepository.NewOrganizationRepository(filePath)

	// Test
	org, err := repo.CreateOrganization(mockOrg)
	assert.NoError(t, err)
	assert.NotNil(t, org)
	// Perform more assertions based on the mockOrg and the expected behaviors

	orgJsonPath := filerepository.GetOrganizationJsonPath(filePath, orgID)

	savedOrg, err := filerepository.ReadOrganizationJsonFile(orgJsonPath)
	assert.NoError(t, err, "Should be able to retrieve the newly created organization without error")
	assert.NotNil(t, savedOrg, "The saved organization should not be nil")
	assert.Equal(t, orgID, savedOrg.ID, "The saved organization ID should match the original ID")
	expectedNames := []string{"Test Org"} // Match this with what's returned by mockOrg.GetDTO()
	assert.Equal(t, expectedNames, savedOrg.AllNames, "The saved organization names should match the original names")

}

func setupTempDir(t *testing.T) (string, func()) {
	t.Helper()
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "orgRepoTest")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Return the path and a cleanup function
	cleanup := func() {
		os.RemoveAll(tempDir) // clean up the temp directory
	}

	return tempDir, cleanup
}
