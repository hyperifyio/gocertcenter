// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"reflect"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewCertificateCreatedDTO(t *testing.T) {
	// Example CertificateDTO and PrivateKeyDTO to use in test cases
	certificateDTO := appdtos.CertificateDTO{
		SerialNumber:              "12345",
		SignedBy:                  "CA",
		Organization:              "Test Org",
		IsCA:                      true,
		IsRootCertificate:         false,
		IsIntermediateCertificate: true,
		IsServerCertificate:       false,
		IsClientCertificate:       false,
		Certificate:               "cert data",
	}
	privateKeyDTO := appdtos.PrivateKeyDTO{
		Certificate: "12345",
		Type:        "RSA",
		PrivateKey:  "key data",
	}

	tests := []struct {
		name        string
		certificate appdtos.CertificateDTO
		privateKey  appdtos.PrivateKeyDTO
		want        appdtos.CertificateCreatedDTO
	}{
		{
			name:        "Valid certificate and private key",
			certificate: certificateDTO,
			privateKey:  privateKeyDTO,
			want: appdtos.CertificateCreatedDTO{
				Certificate: certificateDTO,
				PrivateKey:  privateKeyDTO,
			},
		},
		// Additional test cases could be added here to test various scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appdtos.NewCertificateCreatedDTO(tt.certificate, tt.privateKey)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCertificateCreatedDTO() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
