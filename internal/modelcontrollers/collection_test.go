// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers_test

import (
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"testing"
)

func TestNewControllerCollection(t *testing.T) {
	mockOrganizationService := &mocks.MockOrganizationService{}
	mockCertificateService := &mocks.MockCertificateService{}
	mockPrivateKeyService := &mocks.MockPrivateKeyService{}

	collection := modelcontrollers.NewControllerCollection(mockOrganizationService, mockCertificateService, mockPrivateKeyService)

	if collection.Certificate != mockCertificateService {
		t.Errorf("Certificate service was not correctly assigned")
	}

	if collection.PrivateKey != mockPrivateKeyService {
		t.Errorf("Private Key service was not correctly assigned")
	}
}
