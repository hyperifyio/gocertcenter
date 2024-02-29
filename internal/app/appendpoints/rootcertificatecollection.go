// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetRootCertificateCollectionDefinitions returns OpenAPI definitions
func (c *HttpApiController) GetRootCertificateCollectionDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of root certificate entities",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.CertificateListDTO{}},
				},
			},
		},
	}
}

// GetRootCertificateCollection handles a request to get organization's certificates
func (c *HttpApiController) GetRootCertificateCollection(response apitypes.Response, request apitypes.Request) error {

	controller, err := c.getOrganizationController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	list, err := controller.GetCertificateCollection()
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	c.logf(request, "list len = %d", len(list))
	dto := apputils.ToCertificateListDTO(list)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).GetRootCertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).GetRootCertificateCollection
