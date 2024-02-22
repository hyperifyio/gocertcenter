// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package server

type Server struct {
	listen                         string
	repositoryControllerCollection modelcontrollers.ControllerCollection
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection modelcontrollers.ControllerCollection,
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
