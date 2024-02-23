// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type IndexDTO struct {
	Version string `json:"version"`
}

func NewIndexDTO(version string) IndexDTO {
	return IndexDTO{
		Version: version,
	}
}
