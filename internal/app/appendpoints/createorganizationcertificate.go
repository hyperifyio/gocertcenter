// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateOrganizationCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) CreateOrganizationCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Creates an certificate under a root certificate",
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

// CreateOrganizationCertificate handles a request
func (c *ApiController) CreateOrganizationCertificate(response apitypes.IResponse, request apitypes.IRequest) error {

	// Decode request body
	body, err := c.DecodeCertificateRequestFromRequestBody(request)
	if err != nil {
		log.Printf("[ApiController.CreateOrganizationCertificate]: Request body invalid: %v", err)
		response.SendError(400, "[ApiController.CreateOrganizationCertificate]: request body invalid")
		return nil
	}

	commonName := body.CommonName
	log.Printf("[ApiController.CreateOrganizationCertificate] commonName = %s", commonName)

	certificateType := body.CertificateType
	if certificateType == "" {
		certificateType = appdtos.ClientCertificate
	}
	log.Printf("[ApiController.CreateOrganizationCertificate] certificateType = %s", certificateType)

	// Check certificate type is not a root certificate
	if certificateType == appdtos.RootCertificate {
		response.SendError(400, "[ApiController.CreateOrganizationCertificate]: cannot create a root certificate")
		return nil
	}

	organization := request.GetVariable("organization")

	log.Printf("[GetOrganizationCertificate] organization = %s", organization)

	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[ApiController.CreateOrganizationCertificate] serialNumber = %s", serialNumber.String())

	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to find organization controller: %v", err)
	}

	certificateController, err := organizationController.GetCertificateController(serialNumber)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to find certificate controller: %v", err)
	}

	var cert appmodels.ICertificate
	var privateKey appmodels.IPrivateKey

	if certificateType == appdtos.ClientCertificate {
		cert, privateKey, err = certificateController.NewClientCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to create client certificate: %w", err)
		}
		log.Printf("[ApiController.CreateOrganizationCertificate] created client certificate")

	} else if certificateType == appdtos.ServerCertificate {
		cert, privateKey, err = certificateController.NewServerCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to create server certificate: %w", err)
		}
		log.Printf("[ApiController.CreateOrganizationCertificate] created server certificate")

	} else if certificateType == appdtos.IntermediateCertificate {
		cert, privateKey, err = certificateController.NewIntermediateCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to create intermediate certificate: %w", err)
		}
		log.Printf("[ApiController.CreateOrganizationCertificate] created intermediate certificate")

	} else {
		return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: Unsupported cert type: %s", certificateType)
	}

	data, err := apputils.ToCertificateCreatedDTO(c.certManager, cert, privateKey)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateOrganizationCertificate]: failed to create a DTO: %w", err)
	}

	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateOrganizationCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateOrganizationCertificate
