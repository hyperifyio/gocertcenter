// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type CertificateCreatedDTO struct {

	// Certificate is the generated certificate
	Certificate CertificateDTO `json:"certificate"`

	// PrivateKey may be used to return the backend generated private key
	PrivateKey PrivateKeyDTO `json:"privateKey"`
}

func NewCertificateCreatedDTO(
	certificate CertificateDTO,
	privateKey PrivateKeyDTO,
) CertificateCreatedDTO {
	return CertificateCreatedDTO{
		Certificate: certificate,
		PrivateKey:  privateKey,
	}
}
