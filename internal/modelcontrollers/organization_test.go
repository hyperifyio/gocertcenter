// Copyright (c) 2024. Heusala roup Oy <info@heusalagroup.fi>. All rights reserved.

package modelcontrollers_test

import (
	"errors"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

func TestNewOrganizationController(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	if !controller.UsesOrganizationService(mockService) {
		t.Fatalf("Expected the organization controller to use the mockService, got false")
	}

}

// Now, let's write a test for a successful operation in OrganizationController.
func TestOrganizationController_CreateOrganization_Success(t *testing.T) {
	// Setup
	mockService := &mocks.MockOrganizationService{
		CreateOrganizationFunc: func(certificate models.IOrganization) (models.IOrganization, error) {
			return certificate, nil // Simulate successful creation
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // You'd fill this with actual data
	createdOrg, err := controller.CreateOrganization(newOrg)

	// Verify
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdOrg != newOrg {
		t.Errorf("Expected created organization to be the same as the input")
	}
}

// Test for handling errors from the service
func TestOrganizationController_CreateOrganization_Error(t *testing.T) {
	// Setup
	expectedError := errors.New("failed to create organization")
	mockService := &mocks.MockOrganizationService{
		CreateOrganizationFunc: func(certificate models.IOrganization) (models.IOrganization, error) {
			return nil, expectedError // Simulate failure
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // Assuming this would be populated with test data
	_, err := controller.CreateOrganization(newOrg)

	// Verify
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestOrganizationController_GetExistingOrganization(t *testing.T) {

	orgId := "testorg"
	expectedModel := &models.Organization{ /* Initialized fields */ }

	mockService := &mocks.MockOrganizationService{
		GetExistingOrganizationFunc: func(id string) (models.IOrganization, error) {
			return expectedModel, nil
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)
	organization, err := controller.GetExistingOrganization(orgId)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if organization != expectedModel {
		t.Errorf("Expected to receive %v, got %v", expectedModel, organization)
	}
}
