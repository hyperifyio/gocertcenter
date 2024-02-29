// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints_test

import (
	"net/http"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appendpoints"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestGetRoutes(t *testing.T) {

	mockServer := apimocks.NewMockServer()
	mockApp := new(appmocks.MockApplicationController)
	certManager := new(commonmocks.MockCertificateManager)

	controller := appendpoints.NewHttpApiController(mockServer, mockApp, certManager)

	routes := controller.GetRoutes()

	if len(routes) == 0 {
		t.Fatalf("GetRoutes returned no routes")
	}

	expectedPath := "/"
	found := false
	for _, route := range routes {
		if route.Path == expectedPath && route.Method == http.MethodGet {
			found = true
			// Further checks could be added here to verify the handler and definitions
			// For example, checking if the handler is indexapi.Index might require reflection or interface comparison
			break
		}
	}

	if !found {
		t.Errorf("Expected to find route for path %s, but did not", expectedPath)
	}
}
