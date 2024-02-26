// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
)

func TestNewCollection(t *testing.T) {
	mockOrganizationController := new(appmocks.MockOrganizationController)
	mockCertificateController := new(appmocks.MockCertificateController)
	mockPrivateKeyController := new(appmocks.MockPrivateKeyController)

	collection := appcontrollers.NewCollection(mockOrganizationController, mockCertificateController, mockPrivateKeyController)

	// Use assert.Equal to check if the services were correctly assigned.
	// It's more concise and automatically handles the error message.
	assert.Equal(t, mockOrganizationController, collection.Organization, "Organization service was not correctly assigned")
	assert.Equal(t, mockCertificateController, collection.Certificate, "Certificate service was not correctly assigned")
	assert.Equal(t, mockPrivateKeyController, collection.PrivateKey, "PrivateKey service was not correctly assigned")
}
