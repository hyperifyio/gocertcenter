// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appendpoints_test

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/app/appendpoints"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apimocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestApiController_GetInfo(t *testing.T) {

	expectedInfo := &openapi3.Info{
		Title:   gocertcenter.Name,
		Version: gocertcenter.Version,
		License: &openapi3.License{
			Name: gocertcenter.LicenseName,
			URL:  gocertcenter.LicenseURL,
		},
		Description: gocertcenter.Description,
		Contact: &openapi3.Contact{
			Name:  gocertcenter.SupportName,
			URL:   gocertcenter.SupportURL,
			Email: gocertcenter.SupportEmail,
		},
	}

	mockServer := apimocks.NewMockServer()
	certManager := new(commonmocks.MockCertificateManager)
	mockApp := new(appmocks.MockApplicationController)

	controller := appendpoints.NewHttpApiController(mockServer, mockApp, certManager)

	info := controller.Info()

	if info.Title != expectedInfo.Title ||
		info.Version != expectedInfo.Version ||
		info.License.Name != expectedInfo.License.Name ||
		info.License.URL != expectedInfo.License.URL ||
		info.Description != expectedInfo.Description ||
		info.Contact.Name != expectedInfo.Contact.Name ||
		info.Contact.URL != expectedInfo.Contact.URL ||
		info.Contact.Email != expectedInfo.Contact.Email {
		t.Errorf("Info returned unexpected values")
	}
}
