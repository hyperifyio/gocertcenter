// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestNewCollection(t *testing.T) {
	mockOrganizationService := &appmocks.MockOrganizationService{}
	mockCertificateService := &appmocks.MockCertificateService{}
	mockPrivateKeyService := &appmocks.MockPrivateKeyService{}

	collection := appmodels.NewCollection(mockOrganizationService, mockCertificateService, mockPrivateKeyService)

	if collection.Organization != mockOrganizationService {
		t.Errorf("Certificate service was not correctly assigned")
	}

	if collection.Certificate != mockCertificateService {
		t.Errorf("Certificate service was not correctly assigned")
	}

	if collection.PrivateKey != mockPrivateKeyService {
		t.Errorf("Private Key service was not correctly assigned")
	}
}
