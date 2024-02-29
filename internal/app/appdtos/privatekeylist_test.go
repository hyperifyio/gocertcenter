// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewPrivateKeyListDTO(t *testing.T) {
	// Prepare a sample payload of PrivateKeyDTO instances
	payload := []appdtos.PrivateKeyDTO{
		{
			Certificate: "cert1",
			Type:        "RSA",
			PrivateKey:  "privateKeyData1",
		},
		{
			Certificate: "cert2",
			Type:        "ECDSA",
			PrivateKey:  "privateKeyData2",
		},
	}

	// Create a PrivateKeyListDTO instance using the constructor
	privateKeyListDTO := appdtos.NewPrivateKeyListDTO(payload)

	// Assert that the payload is correctly assigned
	assert.Equal(t, payload, privateKeyListDTO.Payload, "Payload should exactly match the input payload")
}
