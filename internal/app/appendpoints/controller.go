// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// HttpApiController implements apitypes.AppController and GoCertCenterController
type HttpApiController struct {
	server        apitypes.Server
	appController appmodels.ApplicationController
	certManager   managers.CertificateManager
}

func NewHttpApiController(
	server apitypes.Server,
	appController appmodels.ApplicationController,
	certManager managers.CertificateManager,
) *HttpApiController {
	return &HttpApiController{
		server:        server,
		appController: appController,
		certManager:   certManager,
	}
}

// Note! Other methods are defined in adjacent files.

var _ apitypes.AppController = (*HttpApiController)(nil)
var _ GoCertCenterController = (*HttpApiController)(nil)
