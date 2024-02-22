// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package gocertcenter

import (
	"crypto/tls"
	"github.com/hyperifyio/gocertcenter/internal/storage/controllers"
)

type Server struct {
	listen                         string
	repositoryControllerCollection controllers.ControllerCollection
	tlsConfig                      *tls.Config
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection controllers.ControllerCollection,
	tlsConfig *tls.Config,
) *Server {
	return &Server{
		listen,
		repositoryControllerCollection,
		tlsConfig,
	}
}

func (*Server) Start() {

}

func (*Server) Stop() {

}
