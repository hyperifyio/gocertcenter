// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"reflect"
	"testing"
)

func TestNewCertificateRequestDTO(t *testing.T) {
	tests := []struct {
		name            string
		certificateType dtos.CertificateType
		commonName      string
		expiration      int
		dnsNames        []string
		want            dtos.CertificateRequestDTO
	}{
		{
			name:            "Server certificate with multiple DNS names",
			certificateType: dtos.ServerCertificate, // Assuming ServerCertificate is a valid CertificateType
			commonName:      "example.com",
			expiration:      1440, // 1 day
			dnsNames:        []string{"example.com", "www.example.com"},
			want: dtos.CertificateRequestDTO{
				CertificateType: dtos.ServerCertificate,
				CommonName:      "example.com",
				DnsNames:        []string{"example.com", "www.example.com"},
				Expiration:      1440,
			},
		},
		// Add more test cases for different scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dtos.NewCertificateRequestDTO(tt.certificateType, tt.commonName, tt.expiration, tt.dnsNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCertificateRequestDTO() got = %v, want %v", got, tt.want)
			}
		})
	}
}
