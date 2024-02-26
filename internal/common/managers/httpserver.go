// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HttpServerManager implements apitypes.IServerManager
type HttpServerManager struct {
	server *http.Server
}

func (s *HttpServerManager) Serve(l net.Listener) error {
	return s.server.Serve(l)
}

func (s *HttpServerManager) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func NewHttpServerManager(
	address string,
	handler *mux.Router,
) IServerManager {
	return &HttpServerManager{server: &http.Server{
		Addr:    address,
		Handler: handler,
	}}
}

var _ IServerManager = (*HttpServerManager)(nil)
