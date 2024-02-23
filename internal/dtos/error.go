// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type ErrorDTO struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func NewErrorDTO(code int, error string) ErrorDTO {
	return ErrorDTO{
		Code:  code,
		Error: error,
	}
}
