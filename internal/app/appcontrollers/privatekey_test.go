// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
)

func TestNewPrivateKeyController(t *testing.T) {
	mockService := &appmocks.MockPrivateKeyService{}

	controller := appcontrollers.NewPrivateKeyController(mockService)
	if !controller.UsesPrivateKeyService(mockService) {
		t.Errorf("Expected the private key controller to use the mockService, got false")
	}
}

func TestPrivateKeyController_GetExistingPrivateKey(t *testing.T) {

	randomManager := managers.NewRandomManager()

	serialNumber, err := apputils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}

	organization := "testOrg"
	certificates := []appmodels.ISerialNumber{serialNumber}

	expectedKey := &appmodels.PrivateKey{ /* Initialized fields */ }

	mockService := &appmocks.MockPrivateKeyService{
		GetExistingPrivateKeyFunc: func(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
			return expectedKey, nil
		},
	}

	controller := appcontrollers.NewPrivateKeyController(mockService)
	key, err := controller.GetExistingPrivateKey(organization, certificates)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if key != expectedKey {
		t.Errorf("Expected to receive %v, got %v", expectedKey, key)
	}
}

// Now, let's write a test for a successful operation in PrivateKeyController.
func TestPrivateKeyController_CreatePrivateKey_Success(t *testing.T) {
	// Setup
	mockService := &appmocks.MockPrivateKeyService{
		CreatePrivateKeyFunc: func(certificate appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
			return certificate, nil // Simulate successful creation
		},
	}

	controller := appcontrollers.NewPrivateKeyController(mockService)

	// Execute
	newOrg := &appmodels.PrivateKey{} // You'd fill this with actual data
	createdOrg, err := controller.CreatePrivateKey(newOrg)

	// Verify
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdOrg != newOrg {
		t.Errorf("Expected created organization to be the same as the input")
	}
}
