// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"testing"
)

func TestCertificateType_IsClientCertificate(t *testing.T) {
	tests := []struct {
		certificateType dtos.CertificateType
		want            bool
	}{
		{dtos.ClientCertificate, true},
		{dtos.ServerCertificate, false},
		{dtos.IntermediateCertificate, false},
		{dtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsClientCertificate(); got != tt.want {
			t.Errorf("%v.IsClientCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsServerCertificate(t *testing.T) {
	tests := []struct {
		certificateType dtos.CertificateType
		want            bool
	}{
		{dtos.ClientCertificate, false},
		{dtos.ServerCertificate, true},
		{dtos.IntermediateCertificate, false},
		{dtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsServerCertificate(); got != tt.want {
			t.Errorf("%v.IsServerCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsIntermediateCertificate(t *testing.T) {
	tests := []struct {
		certificateType dtos.CertificateType
		want            bool
	}{
		{dtos.ClientCertificate, false},
		{dtos.ServerCertificate, false},
		{dtos.IntermediateCertificate, true},
		{dtos.RootCertificate, false},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsIntermediateCertificate(); got != tt.want {
			t.Errorf("%v.IsIntermediateCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsRootCertificate(t *testing.T) {
	tests := []struct {
		certificateType dtos.CertificateType
		want            bool
	}{
		{dtos.ClientCertificate, false},
		{dtos.ServerCertificate, false},
		{dtos.IntermediateCertificate, false},
		{dtos.RootCertificate, true},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsRootCertificate(); got != tt.want {
			t.Errorf("%v.IsRootCertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}

func TestCertificateType_IsCACertificate(t *testing.T) {
	tests := []struct {
		certificateType dtos.CertificateType
		want            bool
	}{
		{dtos.ClientCertificate, false},
		{dtos.ServerCertificate, false},
		{dtos.IntermediateCertificate, true},
		{dtos.RootCertificate, true},
	}

	for _, tt := range tests {
		if got := tt.certificateType.IsCACertificate(); got != tt.want {
			t.Errorf("%v.IsCACertificate() = %v, want %v", tt.certificateType, got, tt.want)
		}
	}
}
