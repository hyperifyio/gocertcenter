// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestCertificateType_IsClientCertificate(t *testing.T) {
	tests := []struct {
		certificateType appdtos.CertificateType
		want            bool
	}{
		{appdtos.ClientCertificate, true},
		{appdtos.ServerCertificate, false},
		{appdtos.IntermediateCertificate, false},
		{appdtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsClientCertificate(); got != tt.want {
			t.Errorf("%v.IsClientCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsServerCertificate(t *testing.T) {
	tests := []struct {
		certificateType appdtos.CertificateType
		want            bool
	}{
		{appdtos.ClientCertificate, false},
		{appdtos.ServerCertificate, true},
		{appdtos.IntermediateCertificate, false},
		{appdtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsServerCertificate(); got != tt.want {
			t.Errorf("%v.IsServerCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsIntermediateCertificate(t *testing.T) {
	tests := []struct {
		certificateType appdtos.CertificateType
		want            bool
	}{
		{appdtos.ClientCertificate, false},
		{appdtos.ServerCertificate, false},
		{appdtos.IntermediateCertificate, true},
		{appdtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsIntermediateCertificate(); got != tt.want {
			t.Errorf("%v.IsIntermediateCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsRootCertificate(t *testing.T) {
	tests := []struct {
		certificateType appdtos.CertificateType
		want            bool
	}{
		{appdtos.ClientCertificate, false},
		{appdtos.ServerCertificate, false},
		{appdtos.IntermediateCertificate, false},
		{appdtos.RootCertificate, true},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsRootCertificate(); got != tt.want {
			t.Errorf("%v.IsRootCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsCACertificate(t *testing.T) {
	tests := []struct {
		certificateType appdtos.CertificateType
		want            bool
	}{
		{appdtos.ClientCertificate, false},
		{appdtos.ServerCertificate, false},
		{appdtos.IntermediateCertificate, true},
		{appdtos.RootCertificate, true},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsCACertificate(); got != tt.want {
			t.Errorf("%v.IsCACertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}
