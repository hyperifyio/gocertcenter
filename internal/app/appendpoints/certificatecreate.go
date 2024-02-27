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
		log.Printf("[ApiController.CreateCertificate]: Request body invalid: %v", err)
		response.SendError(400, "[ApiController.CreateCertificate]: request body invalid")
		return nil
	}

	// Parse common name
	commonName := body.CommonName
	log.Printf("[ApiController.CreateCertificate] commonName = %s", commonName)

	// Parse certificate type
	certificateType := body.CertificateType
	if certificateType == "" {
		certificateType = appdtos.ClientCertificate
	}
	log.Printf("[ApiController.CreateCertificate] certificateType = %s", certificateType)

	// Check certificate type is not a root certificate
	if certificateType == appdtos.RootCertificate {
		response.SendError(400, "[ApiController.CreateCertificate]: cannot create a root certificate")
		return nil
	}

	// Get organization ID
	organization := request.GetVariable("organization")
	log.Printf("[GetRootCertificate] organization = %s", organization)

	// Parse serial number
	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateCertificate]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[ApiController.CreateCertificate] serialNumber = %s", serialNumber.String())

	// Fetch organization controller
	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateCertificate]: failed to find organization controller: %v", err)
	}

	// Fetch certificate controller
	certificateController, err := organizationController.GetCertificateController(serialNumber)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateCertificate]: failed to find certificate controller: %v", err)
	}

	var cert appmodels.ICertificate
	var privateKey appmodels.IPrivateKey

	if certificateType == appdtos.ClientCertificate {
		cert, privateKey, err = certificateController.NewClientCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateCertificate]: failed to create client certificate: %w", err)
		}
		log.Printf("[ApiController.CreateCertificate] created client certificate")

	} else if certificateType == appdtos.ServerCertificate {
		cert, privateKey, err = certificateController.NewServerCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateCertificate]: failed to create server certificate: %w", err)
		}
		log.Printf("[ApiController.CreateCertificate] created server certificate")

	} else if certificateType == appdtos.IntermediateCertificate {
		cert, privateKey, err = certificateController.NewIntermediateCertificate(commonName)
		if err != nil {
			return fmt.Errorf("[ApiController.CreateCertificate]: failed to create intermediate certificate: %w", err)
		}
		log.Printf("[ApiController.CreateCertificate] created intermediate certificate")

	} else {
		return fmt.Errorf("[ApiController.CreateCertificate]: Unsupported cert type: %s", certificateType)
	}

	data, err := apputils.ToCertificateCreatedDTO(c.certManager, cert, privateKey)
	if err != nil {
		return fmt.Errorf("[ApiController.CreateCertificate]: failed to create a DTO: %w", err)
	}

	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateCertificate
