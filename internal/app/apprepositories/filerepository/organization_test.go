// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
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
		appdtos.NewOrganizationDTO(orgID, "Test Org", []string{"Test Org"}),
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
	orgName := "Test Org"
	mockOrg := &appmocks.MockOrganization{}
	mockOrg.On("GetName").Return(orgName)
	mockOrg.On("GetNames").Return([]string{orgName})
	mockOrg.On("GetID").Return(orgID)
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
	expectedNames := []string{orgName}
	assert.Equal(t, expectedNames, savedOrg.AllNames, "The saved organization names should match the original names")

}

func TestOrganizationRepository_GetExistingOrganization_ReadFail(t *testing.T) {
	// Setup
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	filePath := tempDir
	orgID := "nonexistent_org"
	repo := filerepository.NewOrganizationRepository(filePath)

	// Attempt to get an organization with an ID that does not have a corresponding JSON file
	org, err := repo.GetExistingOrganization(orgID)

	// Test assertions
	assert.Error(t, err, "Expected an error due to failed file read")
	assert.Nil(t, org, "Organization should be nil when file reading fails")
	assert.Contains(t, err.Error(), "failed to read saved organization", "Error message should indicate failure to read organization")
}

func TestOrganizationRepository_CreateOrganization_SaveFail(t *testing.T) {

	// Assuming we have a way to simulate save failure, e.g., by using an invalid file path
	// This step is illustrative; you'd need to adjust based on your actual ability to induce a save failure
	filePath := "/invalid/path/that/does/not/exist"
	repo := filerepository.NewOrganizationRepository(filePath)

	orgId := "org123"
	orgName := "Test Org"

	// Mock organization to pass to CreateOrganization
	mockOrg := &appmocks.MockOrganization{}
	mockOrg.On("GetID").Return(orgId)
	mockOrg.On("GetName").Return(orgName)
	mockOrg.On("GetNames").Return([]string{orgName})
	mockOrg.On("GetID").Return(orgId)

	// Test
	org, err := repo.CreateOrganization(mockOrg)

	// Assertions
	assert.Error(t, err, "Expected an error due to failed organization save")
	assert.Nil(t, org, "Organization should be nil when saving fails")
	assert.Contains(t, err.Error(), "organization creation failed", "Error message should indicate failure in organization creation")
}
