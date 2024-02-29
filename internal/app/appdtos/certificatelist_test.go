// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewCertificateListDTO(t *testing.T) {
	// Prepare a sample payload of CertificateDTO instances
	payload := []appdtos.CertificateDTO{
		{
			CommonName:                "www.example.com",
			SerialNumber:              "12345",
			Parents:                   []string{"Parent1", "Parent2"},
			SignedBy:                  "CA",
			Organization:              "Example Org",
			IsCA:                      false,
			IsRootCertificate:         false,
			IsIntermediateCertificate: true,
			IsServerCertificate:       true,
			IsClientCertificate:       false,
			Certificate:               "CERT_DATA",
		},
		{
			CommonName:                "mail.example.com",
			SerialNumber:              "67890",
			Parents:                   []string{"Parent3", "Parent4"},
			SignedBy:                  "CA",
			Organization:              "Example Org",
			IsCA:                      false,
			IsRootCertificate:         false,
			IsIntermediateCertificate: true,
			IsServerCertificate:       true,
			IsClientCertificate:       false,
			Certificate:               "CERT_DATA_2",
		},
	}

	// Create a CertificateListDTO instance using the constructor
	certificateListDTO := appdtos.NewCertificateListDTO(payload)

	// Assert that the payload is correctly assigned
	assert.Equal(t, payload, certificateListDTO.Payload, "Payload should match the input payload")
}
