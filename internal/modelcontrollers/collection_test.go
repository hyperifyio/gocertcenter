// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
)

func TestNewCollection(t *testing.T) {
	mockOrganizationController := new(mocks.MockOrganizationController)
	mockCertificateController := new(mocks.MockCertificateController)
	mockPrivateKeyController := new(mocks.MockPrivateKeyController)

	collection := modelcontrollers.NewCollection(mockOrganizationController, mockCertificateController, mockPrivateKeyController)

	// Use assert.Equal to check if the services were correctly assigned.
	// It's more concise and automatically handles the error message.
	assert.Equal(t, mockOrganizationController, collection.Organization, "Organization service was not correctly assigned")
	assert.Equal(t, mockCertificateController, collection.Certificate, "Certificate service was not correctly assigned")
	assert.Equal(t, mockPrivateKeyController, collection.PrivateKey, "PrivateKey service was not correctly assigned")
}
