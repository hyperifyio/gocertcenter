// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestApplicationController_UsesOrganizationService(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	assert.True(t, controller.UsesOrganizationService(mockOrgService), "should return true when the service matches")
	assert.False(t, controller.UsesOrganizationService(new(appmocks.MockOrganizationService)), "should return false when the service does not match")
}

func TestApplicationController_GetOrganizationModel(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	mockOrg := new(appmocks.MockOrganization)
	orgID := "org123"

	// Setting up expectations
	mockOrgService.On("FindById", orgID).Return(mockOrg, nil)

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	org, err := controller.Organization(orgID)
	assert.NoError(t, err)
	assert.Equal(t, mockOrg, org, "should return the organization model")

	// Test for not found
	mockOrgService.On("FindById", "nonExistingOrg").Return(nil, fmt.Errorf("not found"))
	_, err = controller.Organization("nonExistingOrg")
	assert.Error(t, err, "should return an error for non-existing organization")
}

func TestApplicationController_NewOrganization(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	mockOrg := new(appmocks.MockOrganization)
	orgID := "neworg"

	// Setting up expectations for new organization creation
	mockOrg.On("ID").Return(orgID)
	mockOrg.On("Name").Return(orgID)
	mockOrg.On("Names").Return([]string{orgID})
	mockOrgService.On("FindById", orgID).Return(nil, fmt.Errorf("not found"))
	mockOrgService.On("Save", mock.Anything).Return(mockOrg, nil)

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	savedOrg, err := controller.NewOrganization(mockOrg)
	assert.NoError(t, err)
	assert.Equal(t, mockOrg, savedOrg, "should successfully create and return the new organization")

	mockOrgService.Mock = mock.Mock{}

	// Test for already existing organization
	mockOrgService.On("FindById", orgID).Return(mockOrg, nil)
	_, err = controller.NewOrganization(mockOrg)
	assert.Error(t, err, "should return an error if the organization already exists")
}

func TestApplicationController_UsesCertificateService(t *testing.T) {
	mockCertService := new(appmocks.MockCertificateService)
	controller := appcontrollers.NewApplicationController(
		nil, mockCertService, nil, nil, nil, 0,
	)

	assert.True(t, controller.UsesCertificateService(mockCertService), "should return true when the service matches")
	assert.False(t, controller.UsesCertificateService(new(appmocks.MockCertificateService)), "should return false when the service does not match")
}

func TestApplicationController_UsesPrivateKeyService(t *testing.T) {
	mockPrivateKeyService := new(appmocks.MockPrivateKeyService)
	controller := appcontrollers.NewApplicationController(
		nil, nil, mockPrivateKeyService, nil, nil, 0,
	)

	assert.True(t, controller.UsesPrivateKeyService(mockPrivateKeyService), "should return true when the service matches")
	assert.False(t, controller.UsesPrivateKeyService(new(appmocks.MockPrivateKeyService)), "should return false when the service does not match")
}

func TestApplicationController_GetOrganizationController(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	mockOrg := new(appmocks.MockOrganization)
	orgID := "org123"

	mockOrgService.On("FindById", orgID).Return(mockOrg, nil)

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	orgController, err := controller.OrganizationController(orgID)
	assert.NoError(t, err)
	assert.NotNil(t, orgController, "should return an CertOrganizationController instance")

	// Test for not found
	mockOrgService.On("FindById", "nonExistingOrg").Return(nil, fmt.Errorf("not found"))
	_, err = controller.OrganizationController("nonExistingOrg")
	assert.Error(t, err, "should return an error for non-existing organization")
}

func TestApplicationController_GetOrganizationCollection(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	mockOrg1 := new(appmocks.MockOrganization)
	mockOrg2 := new(appmocks.MockOrganization)

	// Setting up expectations
	mockOrgService.On("FindAll").Return([]appmodels.Organization{mockOrg1, mockOrg2}, nil)

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	orgs, err := controller.OrganizationCollection()
	assert.NoError(t, err)
	assert.Len(t, orgs, 2, "should return a collection of organizations")

	mockOrgService.Mock = mock.Mock{}

	// Test for failure
	mockOrgService.On("FindAll").Return(nil, fmt.Errorf("failed to fetch organizations"))
	_, err = controller.OrganizationCollection()
	assert.Error(t, err, "should return an error when fetching organizations fails")
}

func TestApplicationController_NewOrganization_ValidationFails(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	invalidMockOrg := new(appmocks.MockOrganization)
	orgID := "invalidOrgID"

	// Assuming ValidateOrganizationModel will fail if ID returns "invalidOrgID"
	invalidMockOrg.On("ID").Return(orgID)

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	_, err := controller.NewOrganization(invalidMockOrg)
	assert.Error(t, err, "should return an error if the organization model is invalid")
	assert.Contains(t, err.Error(), "organization model invalid", "error message should indicate validation failure")
}

func TestApplicationController_NewOrganization_SaveFails(t *testing.T) {
	mockOrgService := new(appmocks.MockOrganizationService)
	mockOrg := new(appmocks.MockOrganization)
	orgID := "org123"

	// Setting up expectations for the successful path up to the save operation
	mockOrg.On("ID").Return(orgID)
	mockOrg.On("Name").Return(orgID)
	mockOrg.On("Names").Return([]string{orgID})
	mockOrgService.On("FindById", orgID).Return(nil, fmt.Errorf("not found"))      // Ensuring FindById indicates org does not exist
	mockOrgService.On("Save", mock.Anything).Return(nil, fmt.Errorf("save error")) // Simulating failure on save

	controller := appcontrollers.NewApplicationController(
		mockOrgService, nil, nil, nil, nil, 0,
	)

	_, err := controller.NewOrganization(mockOrg)
	assert.Error(t, err, "should return an error if the save operation fails")
	assert.Contains(t, err.Error(), "could not create organization", "error message should indicate failure in saving the model")
}
