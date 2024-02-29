// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// RevokeCertificateDefinitions returns OpenAPI definitions
func (c *HttpApiController) RevokeCertificateDefinitions() swagger.Definitions {
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

// RevokeCertificate handles a request
func (c *HttpApiController) RevokeCertificate(response apitypes.Response, request apitypes.Request) error {

	// Fetch the certificate controller
	controller, err := c.getInnerCertificateController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	cert := controller.Certificate()
	model, err := controller.OrganizationController().RevokeCertificate(cert)
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	dto := apputils.ToCertificateRevokedDTO(model)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).RevokeCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).RevokeCertificate
