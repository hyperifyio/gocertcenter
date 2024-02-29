// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// CertificateCollectionDefinitions returns OpenAPI definitions
func (c *HttpApiController) CertificateCollectionDefinitions() swagger.Definitions {
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

// CertificateCollection handles a request to get organization's certificates
func (c *HttpApiController) CertificateCollection(response apitypes.Response, request apitypes.Request) error {

	// certificateType is server, client, root or intermediate
	certificateType := request.QueryParam("type")
	if !(certificateType == "" || certificateType == "server" || certificateType == "client" || certificateType == "root") {
		return c.badRequest(response, request, "query param invalid: type", nil)
	}

	// Fetch root certificate controller
	controller, err := c.rootCertificateController(request)
	if err != nil {
		return c.notFound(response, request, err)
	}

	// Get certificate list
	list, err := controller.ChildCertificateCollection(certificateType)
	if err != nil {
		return c.internalServerError(response, request, err)
	}

	c.logf(request, "list len = %d", len(list))
	dto := apputils.ToCertificateListDTO(list)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).CertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).CertificateCollection
