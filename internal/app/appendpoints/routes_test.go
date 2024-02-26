// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints_test

import (
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/app/appendpoints"
)

func TestGetInfo(t *testing.T) {
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

	info := appendpoints.GetInfo()

	if info.Title != expectedInfo.Title ||
		info.Version != expectedInfo.Version ||
		info.License.Name != expectedInfo.License.Name ||
		info.License.URL != expectedInfo.License.URL ||
		info.Description != expectedInfo.Description ||
		info.Contact.Name != expectedInfo.Contact.Name ||
		info.Contact.URL != expectedInfo.Contact.URL ||
		info.Contact.Email != expectedInfo.Contact.Email {
		t.Errorf("GetInfo returned unexpected values")
	}
}

func TestGetRoutes(t *testing.T) {
	routes := appendpoints.GetRoutes()

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
