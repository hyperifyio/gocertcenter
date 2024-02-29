// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// RootCertificateDefinitions returns OpenAPI definitions
func (c *HttpApiController) RootCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns an root certificate",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.CertificateDTO{}},
				},
			},
		},
	}
}

// RootCertificate handles a request
func (c *HttpApiController) RootCertificate(response apitypes.Response, request apitypes.Request) error {

	controller, err := c.rootCertificateController(request)
	if err != nil {
		return c.notFound(response, request, err)
	}

	model := controller.Certificate()
	c.logf(request, "model = %v", model)

	dto := apputils.ToCertificateDTO(model)
	return c.ok(response, dto)

}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).RootCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).RootCertificate
