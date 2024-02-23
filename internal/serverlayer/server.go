// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package serverlayer

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hyperifyio/gocertcenter/internal/api/indexapi"
	"github.com/hyperifyio/gocertcenter/internal/apierrors"
	"github.com/hyperifyio/gocertcenter/internal/apirequests"
	"github.com/hyperifyio/gocertcenter/internal/apiresponses"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"log"
	"net/http"
	"time"
)

type Server struct {
	listen                         string
	server                         *http.Server
	RepositoryControllerCollection modelcontrollers.ControllerCollection
	router                         *mux.Router
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection modelcontrollers.ControllerCollection,
) *Server {
	return &Server{
		listen:                         listen,
		RepositoryControllerCollection: repositoryControllerCollection,
		router:                         nil,
	}
}

func (s *Server) GetAddress() string {
	return s.listen
}

func (s *Server) SetupRoutes() {
	if s.router == nil {
		s.router = mux.NewRouter()
		//s.router.HandleFunc("/organizations/{organization:[a-z0-9_]+}", responseHandler(organizations.Organization, s)).Methods("GET")
		s.router.HandleFunc("/", responseHandler(indexapi.Index, s)).Methods("GET")
		s.router.NotFoundHandler = responseHandler(apierrors.NotFound, s)
		s.router.MethodNotAllowedHandler = responseHandler(apierrors.MethodNotAllowed, s)
	}
}

func (s *Server) Start() error {

	if s.router == nil {
		s.SetupRoutes()
	}

	if s.server == nil {
		s.server = &http.Server{
			Addr:    s.listen,
			Handler: s.router,
		}
	} else {
		return errors.New("Server already started")
	}

	address := s.listen

	go func() {
		log.Printf("[server] Starting server at %s", address)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Could not start server at %s: %v", address, err)
		}
		log.Printf("[server] Stopped server at %s", address)
		s.server = nil
	}()

	return nil
}

func (s *Server) Stop() {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			log.Printf("[server] Server shutdown error: %v", err)
		}
	}
}

// responseHandler wraps a handler function to inject dependencies.
func responseHandler(
	handler apitypes.RequestHandlerFunc,
	server apitypes.IServer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := apiresponses.NewJSONResponse(w)
		request := apirequests.NewRequest(r)
		handler(response, request, server)
	}
}
