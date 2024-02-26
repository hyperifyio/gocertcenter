// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

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
