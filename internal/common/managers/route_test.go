// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers_test

import (
	"testing"

	"github.com/davidebianchi/gswagger/support/gorilla"
	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestFromGorillaRoute(t *testing.T) {
	router := mux.NewRouter()
	muxRoute := router.NewRoute()
	var testRoute gorilla.Route = muxRoute

	// Assuming gorilla.Route is somehow compatible or an alias for *mux.Route
	convertedRoute := managers.FromGorillaRoute(testRoute)

	if convertedRoute != testRoute {
		t.Errorf("Expected converted route to be the same as the input route")
	}
}

func TestToGorillaRoute(t *testing.T) {
	router := mux.NewRouter()
	testRoute := router.NewRoute()

	// Again, assuming the types are directly compatible or aliased
	convertedRoute := managers.ToGorillaRoute(testRoute)

	if convertedRoute != testRoute {
		t.Errorf("Expected converted route to be the same as the input route")
	}
}
