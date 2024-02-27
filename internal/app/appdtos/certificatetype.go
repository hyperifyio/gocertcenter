// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type CertificateType string

// Define enum values
const (
	RootCertificate         CertificateType = "root"
	IntermediateCertificate CertificateType = "intermediate"
	ServerCertificate       CertificateType = "server"
	ClientCertificate       CertificateType = "client"
)

func (f CertificateType) IsClientCertificate() bool {
	switch f {
	case ClientCertificate:
		return true
	default:
		return false
	}
}

func (f CertificateType) IsServerCertificate() bool {
	switch f {
	case ServerCertificate:
		return true
	default:
		return false
	}
}

func (f CertificateType) IsIntermediateCertificate() bool {
	switch f {
	case IntermediateCertificate:
		return true
	default:
		return false
	}
}

func (f CertificateType) IsRootCertificate() bool {
	switch f {
	case RootCertificate:
		return true
	default:
		return false
	}
}

func (f CertificateType) IsCACertificate() bool {
	switch f {
	case IntermediateCertificate:
		return true
	case RootCertificate:
		return true
	default:
		return false
	}
}
