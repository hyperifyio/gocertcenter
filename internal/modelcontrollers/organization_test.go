// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package modelcontrollers_test

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"testing"
)

// Now, let's write a test for a successful operation in OrganizationController.
func TestOrganizationController_CreateOrganization_Success(t *testing.T) {
	// Setup
	mockService := &mocks.MockOrganizationService{
		CreateOrganizationFunc: func(certificate *models.Organization) (*models.Organization, error) {
			return certificate, nil // Simulate successful creation
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // You'd fill this with actual data
	createdOrg, err := controller.Service.CreateOrganization(newOrg)

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
		CreateOrganizationFunc: func(certificate *models.Organization) (*models.Organization, error) {
			return nil, expectedError // Simulate failure
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // Assuming this would be populated with test data
	_, err := controller.Service.CreateOrganization(newOrg)

	// Verify
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}
