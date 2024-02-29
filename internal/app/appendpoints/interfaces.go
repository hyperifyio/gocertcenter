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
	Info() *openapi3.Info
	Routes() []apitypes.Route

	IndexDefinitions() swagger.Definitions
	Index(response apitypes.Response, request apitypes.Request) error

	OrganizationCollection(response apitypes.Response, request apitypes.Request) error
	OrganizationCollectionDefinitions() swagger.Definitions

	RootCertificateCollection(response apitypes.Response, request apitypes.Request) error
	RootCertificateCollectionDefinitions() swagger.Definitions

	CreateRootCertificate(response apitypes.Response, request apitypes.Request) error
	CreateRootCertificateDefinitions() swagger.Definitions
}
