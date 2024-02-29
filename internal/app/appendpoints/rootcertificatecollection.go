// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetRootCertificateCollectionDefinitions returns OpenAPI definitions
func (c *HttpApiController) RootCertificateCollectionDefinitions() swagger.Definitions {
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
func (c *HttpApiController) RootCertificateCollection(response apitypes.Response, request apitypes.Request) error {

	controller, err := c.organizationController(request)
	if err != nil {
		return c.notFound(response, request, err)
	}

	list, err := controller.CertificateCollection()
	if err != nil {
		return c.internalServerError(response, request, err)
	}

	c.logf(request, "list len = %d", len(list))
	dto := apputils.ToCertificateListDTO(list)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).RootCertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).RootCertificateCollection
