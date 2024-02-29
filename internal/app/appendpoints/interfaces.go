// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GoCertCenterController defines our own methods specific to our HTTP
// end-points, while apitypes.AppController defines common non-application
// specific methods.
type GoCertCenterController interface {
	GetInfo() *openapi3.Info
	GetRoutes() []apitypes.Route

	GetIndexDefinitions() swagger.Definitions
	GetIndex(response apitypes.Response, request apitypes.Request) error

	GetOrganizationCollection(response apitypes.Response, request apitypes.Request) error
	GetOrganizationCollectionDefinitions() swagger.Definitions

	GetRootCertificateCollection(response apitypes.Response, request apitypes.Request) error
	GetRootCertificateCollectionDefinitions() swagger.Definitions

	CreateRootCertificate(response apitypes.Response, request apitypes.Request) error
	CreateRootCertificateDefinitions() swagger.Definitions
}
