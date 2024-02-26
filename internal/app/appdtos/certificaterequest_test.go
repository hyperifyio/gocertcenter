// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"reflect"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewCertificateRequestDTO(t *testing.T) {
	tests := []struct {
		name            string
		certificateType appdtos.CertificateType
		commonName      string
		expiration      int
		dnsNames        []string
		want            appdtos.CertificateRequestDTO
	}{
		{
			name:            "Server certificate with multiple DNS names",
			certificateType: appdtos.ServerCertificate, // Assuming ServerCertificate is a valid CertificateType
			commonName:      "example.com",
			expiration:      1440, // 1 day
			dnsNames:        []string{"example.com", "www.example.com"},
			want: appdtos.CertificateRequestDTO{
				CertificateType: appdtos.ServerCertificate,
				CommonName:      "example.com",
				DnsNames:        []string{"example.com", "www.example.com"},
				Expiration:      1440,
			},
		},
		// Add more test cases for different scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appdtos.NewCertificateRequestDTO(tt.certificateType, tt.commonName, tt.expiration, tt.dnsNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCertificateRequestDTO() got = %v, want %v", got, tt.want)
			}
		})
	}
}
