// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func (c *HttpApiController) IndexDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns information about the running server",
		Description: "This includes the software name and a version",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.IndexDTO{}},
				},
			},
		},
	}
}

// GetIndex handles the GET requests at the root MockURL.
func (c *HttpApiController) Index(response apitypes.Response, request apitypes.Request) error {
	c.log(request, "IndexController")
	dto := appdtos.NewIndexDTO(gocertcenter.Name, gocertcenter.Version)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).IndexDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).Index
