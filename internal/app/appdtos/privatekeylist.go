// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type PrivateKeyListDTO struct {
	Payload []PrivateKeyDTO `json:"payload" jsonschema:"title=Private Key Payload DTOs,required"`
}

func NewPrivateKeyListDTO(
	payload []PrivateKeyDTO,
) PrivateKeyListDTO {
	return PrivateKeyListDTO{
		Payload: payload,
	}
}
