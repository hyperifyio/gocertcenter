// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateRootCertificateDefinitions returns OpenAPI definitions
func (c *HttpApiController) CreateRootCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of organization entities",
		Description: "",
		RequestBody: &swagger.ContentValue{
			Description: "Certificate request data",
			Content: swagger.Content{
				"application/json": {
					Value: appdtos.CertificateRequestDTO{},
				},
			},
		},
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.CertificateDTO{}},
				},
			},
		},
	}
}

// CreateOrganizationRootCertificate handles a request
func (c *HttpApiController) CreateRootCertificate(response apitypes.Response, request apitypes.Request) error {

	body, err := c.DecodeCertificateRequestFromRequestBody(request)
	if err != nil {
		return c.sendBadRequest(response, request, "body invalid", err)
	}

	// Parse certificate type from body
	certificateType := body.CertificateType
	if certificateType == "" {
		certificateType = appdtos.RootCertificate
	}
	if certificateType != appdtos.RootCertificate {
		return c.sendBadRequest(response, request, "body type invalid", nil)
	}

	organizationController, err := c.getOrganizationController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	commonName := body.CommonName

	cert, err := organizationController.NewRootCertificate(commonName)
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	dto := apputils.ToCertificateDTO(cert)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).CreateRootCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).CreateRootCertificate
