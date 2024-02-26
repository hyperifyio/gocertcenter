// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type PrivateKeyDTO struct {

	// Certificate is the serial number of the certificate this key belongs to
	Certificate string `json:"certificate"`

	// Type is the type of the private key
	Type string `json:"Type"`

	// Private key in PEM format
	PrivateKey string `json:"privateKey"`
}

func NewPrivateKeyDTO(
	certificate string,
	_type string,
	privateKey string,
) PrivateKeyDTO {
	return PrivateKeyDTO{
		Certificate: certificate,
		Type:        _type,
		PrivateKey:  privateKey,
	}
}
