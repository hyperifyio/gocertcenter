// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetOrganizationCollectionDefinitions returns OpenAPI definitions
func (c *HttpApiController) GetOrganizationCollectionDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a specific root certificate",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationListDTO{}},
				},
			},
		},
	}
}

// GetOrganizationCollection handles a request
func (c *HttpApiController) GetOrganizationCollection(response apitypes.Response, request apitypes.Request) error {
	list, err := c.appController.GetOrganizationCollection()
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}
	c.logf(request, "list len = %d", len(list))
	dto := apputils.ToOrganizationListDTO(list)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).GetOrganizationCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).GetOrganizationCollection
