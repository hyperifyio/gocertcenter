// Copyright (c) 2024. Heusala roup Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewOrganizationController(t *testing.T) {

	organization := "testorg"
	model := &appmocks.MockOrganization{}
	mockOrganizationRepository := &appmocks.MockOrganizationService{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}
	certManager := commonmocks.NewMockCertificateManager()
	randomManager := commonmocks.NewMockRandomManager()

	controller := appcontrollers.NewOrganizationController(
		organization,
		model,
		mockOrganizationRepository,
		mockCertificateRepository,
		mockPrivateKeyRepository,
		certManager,
		randomManager,
	)

	if !controller.UsesOrganizationService(mockOrganizationRepository) {
		t.Fatalf("Expected the organization controller to use the mockService, got false")
	}

}
