// Copyright (c) 2024. Heusala roup Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
)

func TestNewOrganizationController(t *testing.T) {

	mockService := &appmocks.MockOrganizationService{}
	controller := appcontrollers.NewOrganizationController(mockService)

	if !controller.UsesOrganizationService(mockService) {
		t.Fatalf("Expected the organization controller to use the mockService, got false")
	}

}
