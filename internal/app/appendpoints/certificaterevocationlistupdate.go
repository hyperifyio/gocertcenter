// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

// UpdateCertificateRevocationListDefinitions returns OpenAPI definitions
func (c *HttpApiController) UpdateCertificateRevocationListDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns the certificate revocation list",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/pkix-crl": {Value: appdtos.CertificateListDTO{}},
				},
			},
		},
	}
}

// UpdateCertificateRevocationList handles a request to get organization's certificates
func (c *HttpApiController) UpdateCertificateRevocationList(response apitypes.Response, request apitypes.Request) error {

	// Fetch root certificate controller
	rootCertificateController, err := c.rootCertificateController(request)
	if rootCertificateController == nil {
		return c.notFound(response, request, err)
	}

	// Get a list of revoked certificates

	// Sign a CRL

	// Save a CRL

	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).UpdateCertificateRevocationListDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).UpdateCertificateRevocationList
