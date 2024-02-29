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
	mockCertificateController.On("ApplicationController").Return(mockApplicationController)
	mockCertificateController.On("OrganizationID").Return(orgID)
	mockCertificateController.On("Organization").Return(mockOrganization)
	mockCertificateController.On("OrganizationController").Return(mockOrganizationController)
	mockCertificateController.On("Certificate").Return(mockCertificate)

	controller := appcontrollers.NewPrivateKeyController(
		mockPrivateKey,
		mockCertificateController,
		mockPrivateKeyService,
	)

	t.Run("ApplicationController", func(t *testing.T) {
		assert.Equal(t, mockApplicationController, controller.ApplicationController())
	})

	t.Run("OrganizationID", func(t *testing.T) {
		assert.Equal(t, orgID, controller.OrganizationID())
	})

	t.Run("Organization", func(t *testing.T) {
		assert.Equal(t, mockOrganization, controller.Organization())
	})

	t.Run("OrganizationController", func(t *testing.T) {
		assert.Equal(t, mockOrganizationController, controller.OrganizationController())
	})

	t.Run("Certificate", func(t *testing.T) {
		assert.Equal(t, mockCertificate, controller.Certificate())
	})

	t.Run("CertificateController", func(t *testing.T) {
		assert.Equal(t, mockCertificateController, controller.CertificateController())
	})
}

func TestPrivateKeyController_GetOrganizationController_WithNilParent(t *testing.T) {
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockPrivateKeyService := new(appmocks.MockPrivateKeyService)

	// Initialize the CertPrivateKeyController with a nil parent to simulate the error condition
	controller := appcontrollers.NewPrivateKeyController(
		mockPrivateKey,
		nil, // Pass nil to simulate the absence of a parent ICertificateController
		mockPrivateKeyService,
	)

	orgController := controller.OrganizationController()

	// Verify that the method returns nil when parent is nil, and no error is involved
	assert.Nil(t, orgController, "Expected OrganizationController to return nil when parent is nil")
}
