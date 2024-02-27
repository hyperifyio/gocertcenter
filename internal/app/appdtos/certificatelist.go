// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type CertificateListDTO struct {
	Payload []CertificateDTO `json:"payload" jsonschema:"title=Certificate Payload DTOs,required"`
}

func NewCertificateListDTO(
	payload []CertificateDTO,
) CertificateListDTO {
	return CertificateListDTO{
		Payload: payload,
	}
}
