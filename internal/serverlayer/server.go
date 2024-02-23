// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package serverlayer

import (
	"context"
	"errors"
	"fmt"
	swagger "github.com/davidebianchi/gswagger"
	"github.com/davidebianchi/gswagger/support/gorilla"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/hyperifyio/gocertcenter/internal/api"
	"github.com/hyperifyio/gocertcenter/internal/apierrors"
	"github.com/hyperifyio/gocertcenter/internal/apirequests"
	"github.com/hyperifyio/gocertcenter/internal/apiresponses"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"log"
	"net/http"
	"strings"
	"time"
)

const DefaultHostname = "localhost"
const DefaultProtocol = "http"

type Server struct {
	listen                         string
	server                         *http.Server
	RepositoryControllerCollection modelcontrollers.ControllerCollection
	router                         *mux.Router
	swaggerRouter                  *swagger.Router[gorilla.HandlerFunc, gorilla.Route]
	context                        context.Context
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection modelcontrollers.ControllerCollection,
) *Server {
	s := &Server{
		listen:                         listen,
		RepositoryControllerCollection: repositoryControllerCollection,
	}
	s.InitSetup()
	s.SetupNotFoundHandler(apierrors.NotFound)
	s.SetupMethodNotAllowedHandler(apierrors.MethodNotAllowed)
	return s
}

func (s *Server) IsStarted() bool {
	return s.server != nil
}

func (s *Server) GetAddress() string {
	return s.listen
}

func (s *Server) GetURL() string {
	url := s.listen
	if strings.HasPrefix(url, ":") {
		return fmt.Sprintf("%s://%s%s", DefaultProtocol, DefaultHostname, url)
	}
	return fmt.Sprintf("%s://%s", DefaultProtocol, url)
}

func (s *Server) InitSetup() {

	if s.context == nil {
		s.context = context.Background()
	}

	if s.router == nil {
		s.router = mux.NewRouter()
	}

	if s.swaggerRouter == nil {
		swaggerRouter, err := swagger.NewRouter(
			gorilla.NewRouter(s.router),
			swagger.Options{
				Context: s.context,
				Openapi: &openapi3.T{
					Info: api.GetOpenApiInfo(),
					Servers: []*openapi3.Server{
						{
							URL:         s.GetURL(),
							Description: "Server location",
						},
						// You can add more servers (URLs) if needed
					},
				},
			},
		)
		if err != nil {
			log.Fatalf("Failed to create swagger router: %v", err)
		}
		s.swaggerRouter = swaggerRouter
	}

}

func (s *Server) SetupRoutes(
	routes []apitypes.Route,
) {
	for _, route := range routes {
		s.SetupHandler(route.Method, route.Path, route.Handler, route.Definitions)
	}
}

func (s *Server) SetupNotFoundHandler(handler apitypes.RequestHandlerFunc) {
	s.router.NotFoundHandler = responseHandler(handler, s)
}

func (s *Server) SetupMethodNotAllowedHandler(handler apitypes.RequestHandlerFunc) {
	s.router.MethodNotAllowedHandler = responseHandler(handler, s)
}

func (s *Server) FinalizeSetup() error {
	if err := s.swaggerRouter.GenerateAndExposeOpenapi(); err != nil {
		return fmt.Errorf("Failed to initialize OpenAPI integration: %v", err)
	}
	return nil
}

func (s *Server) SetupHandler(
	method string,
	path string,
	apiHandler apitypes.RequestHandlerFunc,
	definitions swagger.Definitions,
) {
	if method == "" {
		method = http.MethodGet
	}
	if path == "" {
		path = "/"
	}
	handler := responseHandler(apiHandler, s)
	gorillaHandler := adaptToHttpHandlerFunc(handler)
	if _, err := s.swaggerRouter.AddRoute(method, path, gorillaHandler, definitions); err != nil {
		log.Fatalf("Failed to setup route %s %s: %v", method, path, err)
	}
}

func (s *Server) Start() error {

	if s.server != nil {
		return errors.New("Server already started")
	}

	address := s.listen

	if s.router == nil {
		s.InitSetup()
	}

	if err := s.FinalizeSetup(); err != nil {
		return fmt.Errorf("failed to finalize the setup: %v", err)
	}

	s.server = &http.Server{
		Addr:    address,
		Handler: s.router,
	}

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

func (s *Server) Stop() error {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			return fmt.Errorf("[server] Server shutdown error: %v", err)
		}
	}
	return nil
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

func adaptToHttpHandlerFunc(h http.HandlerFunc) gorilla.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
