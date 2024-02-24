// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers_test

import (
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"testing"
)

func TestNewPrivateKeyController(t *testing.T) {
	mockService := &mocks.MockPrivateKeyService{}

	controller := modelcontrollers.NewPrivateKeyController(mockService)
	if !controller.UsesPrivateKeyService(mockService) {
		t.Errorf("Expected the private key controller to use the mockService, got false")
	}
}

func TestPrivateKeyController_GetExistingPrivateKey(t *testing.T) {

	randomManager := managers.NewRandomManager()

	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}

	expectedKey := &models.PrivateKey{ /* Initialized fields */ }

	mockService := &mocks.MockPrivateKeyService{
		GetExistingPrivateKeyFunc: func(serialNumber models.SerialNumber) (models.IPrivateKey, error) {
			return expectedKey, nil
		},
	}

	controller := modelcontrollers.NewPrivateKeyController(mockService)
	key, err := controller.GetExistingPrivateKey(serialNumber)
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
	mockService := &mocks.MockPrivateKeyService{
		CreatePrivateKeyFunc: func(certificate models.IPrivateKey) (models.IPrivateKey, error) {
			return certificate, nil // Simulate successful creation
		},
	}

	controller := modelcontrollers.NewPrivateKeyController(mockService)

	// Execute
	newOrg := &models.PrivateKey{} // You'd fill this with actual data
	createdOrg, err := controller.CreatePrivateKey(newOrg)

	// Verify
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdOrg != newOrg {
		t.Errorf("Expected created organization to be the same as the input")
	}
}
