// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetOrganizationCollectionDefinitions returns OpenAPI definitions
func (c *ApiController) GetOrganizationCollectionDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a specific root certificate",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationListDTO{}},
				},
			},
		},
	}
}

// GetOrganizationCollection handles a request
func (c *ApiController) GetOrganizationCollection(response apitypes.IResponse, request apitypes.IRequest) error {
	list, err := c.appController.GetOrganizationCollection()
	if err != nil {
		return fmt.Errorf("[ApiController.GetOrganization]: failed to find a collection: %v", err)
	}
	log.Printf("[GetOrganizationCollection] Request: list = %d", len(list))
	data := apputils.ToOrganizationListDTO(list)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetOrganizationCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetOrganizationCollection
