// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"reflect"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewCertificateDTO(t *testing.T) {
	tests := []struct {
		commonName                string
		serialNumber              string
		signedBy                  string
		parents                   []string
		organization              string
		isCA                      bool
		isRootCertificate         bool
		isIntermediateCertificate bool
		isServerCertificate       bool
		isClientCertificate       bool
		certificate               string
		want                      appdtos.CertificateDTO
	}{
		{
			commonName:                "Root CA certificate",
			serialNumber:              "123456789",
			signedBy:                  "Self",
			parents:                   []string{"Self"},
			organization:              "Test Org",
			isCA:                      true,
			isRootCertificate:         true,
			isIntermediateCertificate: false,
			isServerCertificate:       false,
			isClientCertificate:       false,
			certificate:               "cert-data-root",
			want: appdtos.CertificateDTO{
				CommonName:                "Root CA certificate",
				SerialNumber:              "123456789",
				SignedBy:                  "Self",
				Parents:                   []string{"Self"},
				Organization:              "Test Org",
				IsCA:                      true,
				IsRootCertificate:         true,
				IsIntermediateCertificate: false,
				IsServerCertificate:       false,
				IsClientCertificate:       false,
				Certificate:               "cert-data-root",
			},
		},
		{
			commonName:                "Intermediate CA certificate",
			serialNumber:              "987654321",
			signedBy:                  "Root CA",
			parents:                   []string{"Root CA"},
			organization:              "Test Org",
			isCA:                      true,
			isRootCertificate:         false,
			isIntermediateCertificate: true,
			isServerCertificate:       false,
			isClientCertificate:       false,
			certificate:               "cert-data-intermediate",
			want: appdtos.CertificateDTO{
				CommonName:                "Intermediate CA certificate",
				SerialNumber:              "987654321",
				SignedBy:                  "Root CA",
				Parents:                   []string{"Root CA"},
				Organization:              "Test Org",
				IsCA:                      true,
				IsRootCertificate:         false,
				IsIntermediateCertificate: true,
				IsServerCertificate:       false,
				IsClientCertificate:       false,
				Certificate:               "cert-data-intermediate",
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.commonName, func(t *testing.T) {
			got := appdtos.NewCertificateDTO(
				tt.commonName,
				tt.serialNumber,
				tt.parents,
				tt.signedBy,
				tt.organization,
				tt.isCA,
				tt.isRootCertificate,
				tt.isIntermediateCertificate,
				tt.isServerCertificate,
				tt.isClientCertificate,
				tt.certificate,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCertificateDTO() = %v,\n want %v\n", got, tt.want)
			}
		})
	}
}
