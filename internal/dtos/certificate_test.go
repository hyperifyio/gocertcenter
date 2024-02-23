// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"reflect"
	"testing"
)

func TestNewCertificateDTO(t *testing.T) {
	tests := []struct {
		commonName                string
		serialNumber              string
		signedBy                  string
		organization              string
		isCA                      bool
		isRootCertificate         bool
		isIntermediateCertificate bool
		isServerCertificate       bool
		isClientCertificate       bool
		certificate               string
		want                      dtos.CertificateDTO
	}{
		{
			commonName:                "Root CA certificate",
			serialNumber:              "123456789",
			signedBy:                  "Self",
			organization:              "Test Org",
			isCA:                      true,
			isRootCertificate:         true,
			isIntermediateCertificate: false,
			isServerCertificate:       false,
			isClientCertificate:       false,
			certificate:               "cert-data-root",
			want: dtos.CertificateDTO{
				CommonName:                "Root CA certificate",
				SerialNumber:              "123456789",
				SignedBy:                  "Self",
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
			organization:              "Test Org",
			isCA:                      true,
			isRootCertificate:         false,
			isIntermediateCertificate: true,
			isServerCertificate:       false,
			isClientCertificate:       false,
			certificate:               "cert-data-intermediate",
			want: dtos.CertificateDTO{
				CommonName:                "Intermediate CA certificate",
				SerialNumber:              "987654321",
				SignedBy:                  "Root CA",
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
			got := dtos.NewCertificateDTO(
				tt.commonName,
				tt.serialNumber,
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
