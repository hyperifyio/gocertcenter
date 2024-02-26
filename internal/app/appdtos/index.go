// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type IndexDTO struct {
	Name    string `json:"name" jsonschema:"title=The name of the project,required" jsonschema_extras:"example=gocertcenter"`
	Version string `json:"version" jsonschema:"title=The version of the package,required" jsonschema_extras:"example=0.0.1"`
}

func NewIndexDTO(
	name string,
	version string,
) IndexDTO {
	return IndexDTO{
		Name:    name,
		Version: version,
	}
}
