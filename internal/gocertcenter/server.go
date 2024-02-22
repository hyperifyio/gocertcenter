// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package gocertcenter

import (
	"github.com/hyperifyio/gocertcenter/internal/storage/controllers"
)

type Server struct {
	listen                         string
	repositoryControllerCollection controllers.ControllerCollection
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection controllers.ControllerCollection,
) *Server {
	return &Server{
		listen,
		repositoryControllerCollection,
	}
}

func (*Server) Start() {

}

func (*Server) Stop() {

}
