// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apitypes

import (
	"github.com/davidebianchi/gswagger/support/gorilla"
	"github.com/gorilla/mux"
)

func FromGorillaRoute(route gorilla.Route) *mux.Route {
	return route
}

func ToGorillaRoute(route *mux.Route) gorilla.Route {
	return route
}
