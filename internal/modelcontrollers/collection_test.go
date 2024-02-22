// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"testing"
)

func TestNewControllerCollection(t *testing.T) {
	mockCertificateService := &mocks.MockCertificateService{}
	mockPrivateKeyService := &mocks.MockPrivateKeyService{}

	collection := NewControllerCollection(mockCertificateService, mockPrivateKeyService)

	if collection.Certificate != mockCertificateService {
		t.Errorf("Certificate service was not correctly assigned")
	}

	if collection.PrivateKey != mockPrivateKeyService {
		t.Errorf("Private Key service was not correctly assigned")
	}
}
