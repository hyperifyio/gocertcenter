// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"testing"
)

func TestNewCollection(t *testing.T) {
	mockOrganizationService := &mocks.MockOrganizationService{}
	mockCertificateService := &mocks.MockCertificateService{}
	mockPrivateKeyService := &mocks.MockPrivateKeyService{}

	collection := models.NewCollection(mockOrganizationService, mockCertificateService, mockPrivateKeyService)

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
