// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"testing"
)

func TestNewPrivateKeyController(t *testing.T) {
	mockService := &mocks.MockPrivateKeyService{}

	controller := NewPrivateKeyController(mockService)
	if controller.service != mockService {
		t.Errorf("Expected service to be set to the provided mockService, but it was not")
	}
}

func TestPrivateKeyController_GetPrivateKey(t *testing.T) {

	randomManager := managers.NewRandomManager()

	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}

	expectedKey := &models.PrivateKey{ /* Initialized fields */ }

	mockService := &mocks.MockPrivateKeyService{
		GetExistingPrivateKeyFunc: func(serialNumber models.SerialNumber) (*models.PrivateKey, error) {
			return expectedKey, nil
		},
	}

	controller := NewPrivateKeyController(mockService)
	key, err := controller.service.GetExistingPrivateKey(serialNumber)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if key != expectedKey {
		t.Errorf("Expected to receive %v, got %v", expectedKey, key)
	}
}
