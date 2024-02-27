// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

// CertificateRequestDTO is the body for creating certificates
type CertificateRequestDTO struct {

	// CertificateType is the type of the certificate to create
	CertificateType CertificateType `json:"type"`

	// CommonName of the certificate. This is also added to the DnsNames for server certificates.
	CommonName string `json:"commonName"`

	// DnsNames of the server certificate
	DnsNames []string `json:"dnsNames"`

	// Expiration in minutes
	Expiration int `json:"expiration"`
}

func NewCertificateRequestDTO(
	certificateType CertificateType,
	commonName string,
	expiration int,
	dnsNames []string,
) CertificateRequestDTO {
	return CertificateRequestDTO{
		CertificateType: certificateType,
		CommonName:      commonName,
		DnsNames:        dnsNames,
		Expiration:      expiration,
	}
}
