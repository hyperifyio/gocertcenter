// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appendpoints"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewApiController(t *testing.T) {
	mockServer := new(apimocks.MockServer)
	mockApp := new(appmocks.MockApplicationController)
	certManager := new(commonmocks.MockCertificateManager)
	controller := appendpoints.NewHttpApiController(mockServer, mockApp, certManager)
	assert.NotNil(t, controller)
}
