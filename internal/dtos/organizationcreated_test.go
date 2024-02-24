// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"reflect"
	"testing"
)

func TestNewOrganizationCreatedDTO(t *testing.T) {
	// Example data to use in test cases
	organization := dtos.OrganizationDTO{ID: "org1", Name: "Test Organization", AllNames: []string{"Test Organization"}}
	certificate := dtos.CertificateDTO{SerialNumber: "123", Organization: "Test Organization", IsCA: true}
	privateKey := dtos.PrivateKeyDTO{Certificate: "123", Type: "RSA", PrivateKey: "key payload"}

	tests := []struct {
		name         string
		organization dtos.OrganizationDTO
		certificates []dtos.CertificateDTO
		privateKeys  []dtos.PrivateKeyDTO
		want         dtos.OrganizationCreatedDTO
	}{
		{
			name:         "Organization with certificates and private keys",
			organization: organization,
			certificates: []dtos.CertificateDTO{certificate},
			privateKeys:  []dtos.PrivateKeyDTO{privateKey},
			want: dtos.OrganizationCreatedDTO{
				Organization: organization,
				Certificates: []dtos.CertificateDTO{certificate},
				PrivateKeys:  []dtos.PrivateKeyDTO{privateKey},
			},
		},
		// Additional test cases can be defined here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dtos.NewOrganizationCreatedDTO(tt.organization, tt.certificates, tt.privateKeys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrganizationCreatedDTO() = %+v,\n want %+v\n", got, tt.want)
			}
		})
	}
}
