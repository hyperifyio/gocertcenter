// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiserver

import (
	"log"
	"net/http"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apirequests"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apiresponses"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// ResponseHandler wraps a handler function to inject dependencies.
func ResponseHandler(handler apitypes.RequestHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := apiresponses.NewJSONResponse(w)
		request := apirequests.NewRequest(r)
		err := handler(response, request)
		if err != nil {
			log.Printf("[server] Request handler failed: %v", err)
			response.SendError(500, "Internal ApplicationServer Error")
		}
	}
}
