// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

// DecodeOrganizationFromRequestBody parses organization DTO from request body
func (c *ApiController) DecodeOrganizationFromRequestBody(request apitypes.IRequest) (appdtos.OrganizationDTO, error) {

	if request == nil {
		return appdtos.OrganizationDTO{}, errors.New("request must be defined")
	}

	bodyIO := request.Body()

	// Decode the JSON body into the struct
	var body appdtos.OrganizationDTO
	err := json.NewDecoder(bodyIO).Decode(&body)
	if err != nil {
		return appdtos.OrganizationDTO{}, fmt.Errorf("request decoding failed: %s", err)
	}
	_ = bodyIO.Close()

	return body, nil
}

// DecodeCertificateRequestFromRequestBody parses organization DTO from request body
func (c *ApiController) DecodeCertificateRequestFromRequestBody(request apitypes.IRequest) (appdtos.CertificateRequestDTO, error) {

	if request == nil {
		return appdtos.CertificateRequestDTO{}, errors.New("request must be defined")
	}

	bodyIO := request.Body()

	// Decode the JSON body into the struct
	var body appdtos.CertificateRequestDTO
	err := json.NewDecoder(bodyIO).Decode(&body)
	if err != nil {
		return appdtos.CertificateRequestDTO{}, fmt.Errorf("request decoding failed: %s", err)
	}
	_ = bodyIO.Close()

	return body, nil
}
