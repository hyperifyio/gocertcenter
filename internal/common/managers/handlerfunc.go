// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"net/http"

	"github.com/davidebianchi/gswagger/support/gorilla"
)

// ToGorillaHandlerFunc converts http.HandlerFunc to gorilla.HandlerFunc. These
// types are identical but adaptation required because of different package.
func ToGorillaHandlerFunc(h http.HandlerFunc) gorilla.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
