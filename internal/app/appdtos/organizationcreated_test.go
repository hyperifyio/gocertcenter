// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"reflect"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewOrganizationCreatedDTO(t *testing.T) {
	// Example data to use in test cases
	organization := appdtos.OrganizationDTO{ID: "org1", Name: "Test Organization", AllNames: []string{"Test Organization"}}
	certificate := appdtos.CertificateDTO{SerialNumber: "123", Organization: "Test Organization", IsCA: true}
	privateKey := appdtos.PrivateKeyDTO{Certificate: "123", Type: "RSA", PrivateKey: "key payload"}

	tests := []struct {
		name         string
		organization appdtos.OrganizationDTO
		certificates []appdtos.CertificateDTO
		privateKeys  []appdtos.PrivateKeyDTO
		want         appdtos.OrganizationCreatedDTO
	}{
		{
			name:         "Organization with certificates and private keys",
			organization: organization,
			certificates: []appdtos.CertificateDTO{certificate},
			privateKeys:  []appdtos.PrivateKeyDTO{privateKey},
			want: appdtos.OrganizationCreatedDTO{
				Organization: organization,
				Certificates: []appdtos.CertificateDTO{certificate},
				PrivateKeys:  []appdtos.PrivateKeyDTO{privateKey},
			},
		},
		// Additional test cases can be defined here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appdtos.NewOrganizationCreatedDTO(tt.organization, tt.certificates, tt.privateKeys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrganizationCreatedDTO() = %+v,\n want %+v\n", got, tt.want)
			}
		})
	}
}
