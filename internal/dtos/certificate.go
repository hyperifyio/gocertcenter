// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type CertificateDTO struct {
	CommonName                string `json:"commonName"`
	SerialNumber              string `json:"serialNumber"`
	SignedBy                  string `json:"signedBy"`
	Organization              string `json:"organization"`
	IsCA                      bool   `json:"isCA"`
	IsRootCertificate         bool   `json:"isRootCertificate"`
	IsIntermediateCertificate bool   `json:"isIntermediateCertificate"`
	IsServerCertificate       bool   `json:"isServerCertificate"`
	IsClientCertificate       bool   `json:"isClientCertificate"`
	Certificate               string `json:"certificate"`
}

func NewCertificateDTO(
	commonName string,
	serialNumber string,
	signedBy string,
	organization string,
	isCA bool,
	isRootCertificate bool,
	isIntermediateCertificate bool,
	isServerCertificate bool,
	isClientCertificate bool,
	certificate string,
) CertificateDTO {
	return CertificateDTO{
		CommonName:                commonName,
		SerialNumber:              serialNumber,
		SignedBy:                  signedBy,
		Organization:              organization,
		IsCA:                      isCA,
		IsRootCertificate:         isRootCertificate,
		IsIntermediateCertificate: isIntermediateCertificate,
		IsServerCertificate:       isServerCertificate,
		IsClientCertificate:       isClientCertificate,
		Certificate:               certificate,
	}
}
