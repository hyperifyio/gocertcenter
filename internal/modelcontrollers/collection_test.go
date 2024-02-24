// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
)

func TestNewCollection(t *testing.T) {
	mockOrganizationController := &mocks.MockOrganizationController{}
	mockCertificateController := &mocks.MockCertificateController{}
	mockPrivateKeyController := &mocks.MockPrivateKeyController{}

	collection := modelcontrollers.NewCollection(mockOrganizationController, mockCertificateController, mockPrivateKeyController)

	if collection.Organization != mockOrganizationController {
		t.Errorf("Organization service was not correctly assigned")
	}

	if collection.Certificate != mockCertificateController {
		t.Errorf("Certificate service was not correctly assigned")
	}

	if collection.PrivateKey != mockPrivateKeyController {
		t.Errorf("PrivateKey service was not correctly assigned")
	}
}
