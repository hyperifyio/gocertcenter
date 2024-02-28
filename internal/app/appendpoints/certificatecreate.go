// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) CreateCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Creates another certificate under a root certificate",
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

// CreateCertificate handles a request
func (c *ApiController) CreateCertificate(response apitypes.IResponse, request apitypes.IRequest) error {

	// Decode request body
	body, err := c.DecodeCertificateRequestFromRequestBody(request)
	if err != nil {
		return c.sendBadRequest(response, request, "body invalid", err)
	}

	// Parse common name
	commonName := body.CommonName
	c.logf(request, "commonName = %s", commonName)

	// Parse certificate type
	certificateType := body.CertificateType
	if certificateType == "" {
		certificateType = appdtos.ClientCertificate
	}
	c.logf(request, "certificateType = %s", certificateType)

	// Check certificate type is not a root certificate
	if certificateType == appdtos.RootCertificate {
		return c.sendBadRequest(response, request, "body type invalid", err)
	}

	// Fetch root certificate controller
	rootCertificateController, err := c.getRootCertificateController(request)
	if rootCertificateController == nil {
		return c.sendNotFound(response, request, err)
	}

	var cert appmodels.ICertificate
	var privateKey appmodels.IPrivateKey

	if certificateType == appdtos.ClientCertificate {

		cert, privateKey, err = rootCertificateController.NewClientCertificate(commonName)
		if err != nil {
			return c.sendInternalServerError(response, request, err)
		}
		c.logf(request, "created client certificate: %s", cert.GetSerialNumber())

	} else if certificateType == appdtos.ServerCertificate {

		cert, privateKey, err = rootCertificateController.NewServerCertificate(commonName)
		if err != nil {
			return c.sendInternalServerError(response, request, err)
		}
		c.logf(request, "created server certificate: %s", cert.GetSerialNumber())

	} else if certificateType == appdtos.IntermediateCertificate {

		cert, privateKey, err = rootCertificateController.NewIntermediateCertificate(commonName)
		if err != nil {
			return c.sendInternalServerError(response, request, err)
		}
		c.logf(request, "created intermediate certificate: %s", cert.GetSerialNumber())

	} else {
		return c.sendBadRequest(response, request, fmt.Sprintf("unsupported cert type: %s", certificateType), err)
	}

	dto, err := apputils.ToCertificateCreatedDTO(c.certManager, cert, privateKey)
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateCertificate
