// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
)

func TestNewPrivateKeyController(t *testing.T) {

	// organization := "testorg"
	// mockOrganizationRepository := &appmocks.MockOrganizationService{}
	// mockCertificateRepository := &appmocks.MockCertificateService{}
	// certManager := commonmocks.NewMockCertificateManager()
	// randomManager := commonmocks.NewMockRandomManager()

	model := &appmocks.MockPrivateKey{}
	mockPrivateKeyService := &appmocks.MockPrivateKeyService{}
	mockCertificateController := &appmocks.MockCertificateController{}

	controller := appcontrollers.NewPrivateKeyController(
		model,
		mockCertificateController,
		mockPrivateKeyService,
	)

	if !controller.UsesPrivateKeyService(mockPrivateKeyService) {
		t.Errorf("Expected the private key controller to use the mockPrivateKeyService, got false")
	}

}

func TestPrivateKeyController_Delegations(t *testing.T) {
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockPrivateKeyService := new(appmocks.MockPrivateKeyService)
	mockCertificateController := new(appmocks.MockCertificateController)
	mockApplicationController := new(appmocks.MockApplicationController)
	mockOrganizationController := new(appmocks.MockOrganizationController)
	mockOrganization := new(appmocks.MockOrganization)
	mockCertificate := new(appmocks.MockCertificate)

	// Organization ID and other return values for the mocked calls
	orgID := "org123"
	mockCertificateController.On("GetApplicationController").Return(mockApplicationController)
	mockCertificateController.On("GetOrganizationID").Return(orgID)
	mockCertificateController.On("GetOrganizationModel").Return(mockOrganization)
	mockCertificateController.On("GetOrganizationController").Return(mockOrganizationController)
	mockCertificateController.On("GetCertificateModel").Return(mockCertificate)

	controller := appcontrollers.NewPrivateKeyController(
		mockPrivateKey,
		mockCertificateController,
		mockPrivateKeyService,
	)

	t.Run("GetApplicationController", func(t *testing.T) {
		assert.Equal(t, mockApplicationController, controller.GetApplicationController())
	})

	t.Run("GetOrganizationID", func(t *testing.T) {
		assert.Equal(t, orgID, controller.GetOrganizationID())
	})

	t.Run("GetOrganizationModel", func(t *testing.T) {
		assert.Equal(t, mockOrganization, controller.GetOrganizationModel())
	})

	t.Run("GetOrganizationController", func(t *testing.T) {
		assert.Equal(t, mockOrganizationController, controller.GetOrganizationController())
	})

	t.Run("GetCertificateModel", func(t *testing.T) {
		assert.Equal(t, mockCertificate, controller.GetCertificateModel())
	})

	t.Run("GetCertificateController", func(t *testing.T) {
		assert.Equal(t, mockCertificateController, controller.GetCertificateController())
	})
}

func TestPrivateKeyController_GetOrganizationController_WithNilParent(t *testing.T) {
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockPrivateKeyService := new(appmocks.MockPrivateKeyService)

	// Initialize the PrivateKeyController with a nil parent to simulate the error condition
	controller := appcontrollers.NewPrivateKeyController(
		mockPrivateKey,
		nil, // Pass nil to simulate the absence of a parent ICertificateController
		mockPrivateKeyService,
	)

	orgController := controller.GetOrganizationController()

	// Verify that the method returns nil when parent is nil, and no error is involved
	assert.Nil(t, orgController, "Expected GetOrganizationController to return nil when parent is nil")
}
