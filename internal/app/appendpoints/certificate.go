// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CertificateDefinitions returns OpenAPI definitions
func (c *HttpApiController) CertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a certificate entity owned by a root certificate",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationDTO{}},
				},
			},
		},
	}
}

// Certificate handles a request
func (c *HttpApiController) Certificate(response apitypes.Response, request apitypes.Request) error {

	// Fetch the certificate controller
	controller, err := c.innerCertificateController(request)
	if err != nil {
		return c.notFound(response, request, err)
	}

	model := controller.Certificate()
	dto := apputils.ToCertificateDTO(model)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).CertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).Certificate
