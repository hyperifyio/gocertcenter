// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateRootCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) CreateRootCertificateDefinitions() swagger.Definitions {
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
func (c *ApiController) CreateRootCertificate(response apitypes.IResponse, request apitypes.IRequest) error {

	body, err := c.DecodeCertificateRequestFromRequestBody(request)
	if err != nil {
		log.Printf("[ApiController.CreateRootCertificate]: Request body invalid: %v", err)
		response.SendError(400, "[ApiController.CreateRootCertificate]: request body invalid")
		return nil
	}

	certificateType := body.CertificateType
	if certificateType == "" {
		certificateType = appdtos.RootCertificate
	}

	if certificateType != appdtos.RootCertificate {
		response.SendError(400, "[ApiController.CreateRootCertificate]: only root certificate type supported")
		return nil
	}

	organization := request.GetVariable("organization")

	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateRootCertificate]: failed to find organization controller: %w", err)
	}

	commonName := body.CommonName

	cert, err := organizationController.NewRootCertificate(commonName)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateRootCertificate]: [OrganizationController(%s).NewRootCertificate]: failed: %w", organization, err)
	}

	data := apputils.GetCertificateDTO(cert)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateRootCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateRootCertificate
