// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"context"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpServerManager struct {
	server *http.Server
}

func NewHttpServerManager(
	address string,
	handler *mux.Router,
) apitypes.IServerManager {
	return &HttpServerManager{server: &http.Server{
		Addr:    address,
		Handler: handler,
	}}
}

func (s *HttpServerManager) Serve(l net.Listener) error {
	return s.server.Serve(l)
}

func (s *HttpServerManager) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
