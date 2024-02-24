// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package serverlayer

import (
	"context"
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/hashutils"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/apierrors"
	"github.com/hyperifyio/gocertcenter/internal/apirequests"
	"github.com/hyperifyio/gocertcenter/internal/apiresponses"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
)

const DefaultHostname = "localhost"
const DefaultProtocol = "http"
const DefaultOpenApiTitle = "Example API"
const DefaultOpenApiVersion = "0.0.0"

// DefaultStartDuration This is a time to wait for any possible errors while starting the server
const DefaultStartDuration = 500 * time.Millisecond

type Server struct {
	listen                         string
	server                         apitypes.IServerManager
	repositoryControllerCollection modelcontrollers.ControllerCollection
	router                         *mux.Router
	listener                       *net.Listener
	context                        *context.Context
	swaggerManager                 apitypes.ISwaggerManager
	swaggerFactory                 apitypes.NewSwaggerManagerFunc
	serverFactory                  apitypes.NewServerManagerFunc
	info                           *openapi3.Info
	hashFactory                    apitypes.Hash64FactoryFunc
}

// NewServer ..
func NewServer(
	listen string,
	repositoryControllerCollection modelcontrollers.ControllerCollection,
	swaggerFactory apitypes.NewSwaggerManagerFunc,
) (*Server, error) {
	s := &Server{
		listen:                         listen,
		repositoryControllerCollection: repositoryControllerCollection,
		swaggerFactory:                 swaggerFactory,
		serverFactory:                  nil,
		swaggerManager:                 nil,
	}
	if err := s.InitSetup(); err != nil {
		return nil, fmt.Errorf("[server] server initialization failed: %v", err)
	}
	s.SetupNotFoundHandler(apierrors.NotFound)
	s.SetupMethodNotAllowedHandler(apierrors.MethodNotAllowed)
	return s, nil
}

func (s *Server) IsStarted() bool {
	return s.server != nil
}

func (s *Server) GetAddress() string {
	return s.listen
}

func (s *Server) SetInfo(info *openapi3.Info) {
	s.info = info
}

func (s *Server) GetInfo() *openapi3.Info {
	if s.info == nil {
		s.info = &openapi3.Info{
			Title:   DefaultOpenApiTitle,
			Version: DefaultOpenApiVersion,
		}
	}
	return s.info
}

func (s *Server) GetURL() string {
	url := s.listen
	if strings.HasPrefix(url, ":") {
		return fmt.Sprintf("%s://%s%s", DefaultProtocol, DefaultHostname, url)
	}
	return fmt.Sprintf("%s://%s", DefaultProtocol, url)
}

func (s *Server) SetSwaggerFactory(factory apitypes.NewSwaggerManagerFunc) {
	s.swaggerFactory = factory
}

func (s *Server) SetServerFactory(factory apitypes.NewServerManagerFunc) {
	s.serverFactory = factory
}

func (s *Server) InitSetup() error {

	if s.context == nil {
		ctx := context.Background()
		s.context = &ctx
	}

	if s.router == nil {
		s.router = mux.NewRouter()
	}

	if s.swaggerFactory == nil {
		s.swaggerFactory = managers.NewSwaggerManager
	}

	if s.swaggerManager == nil {
		swaggerManager, err := s.swaggerFactory(
			s.router,
			s.context,
			s.GetURL(),
			"Server location",
			s.GetInfo(),
		)
		if err != nil {
			return fmt.Errorf("[server] failed to create swagger router: %v", err)
		}
		s.swaggerManager = swaggerManager
	}

	return nil

}

func (s *Server) SetupRoutes(
	routes []apitypes.Route,
) error {
	for _, route := range routes {
		err := s.SetupHandler(route.Method, route.Path, route.Handler, route.Definitions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) SetupNotFoundHandler(handler apitypes.RequestHandlerFunc) {
	s.router.NotFoundHandler = responseHandler(handler, s)
}

func (s *Server) SetupMethodNotAllowedHandler(handler apitypes.RequestHandlerFunc) {
	s.router.MethodNotAllowedHandler = responseHandler(handler, s)
}

func (s *Server) FinalizeSetup() error {
	if err := s.swaggerManager.GenerateAndExposeOpenapi(); err != nil {
		return fmt.Errorf("failed to initialize OpenAPI integration: %v", err)
	}
	return nil
}

func (s *Server) SetupHandler(
	method string,
	path string,
	apiHandler apitypes.RequestHandlerFunc,
	definitions swagger.Definitions,
) error {
	if method == "" {
		method = http.MethodGet
	}
	if path == "" {
		path = "/"
	}
	handler := responseHandler(apiHandler, s)
	if _, err := s.swaggerManager.AddRoute(method, path, handler, definitions); err != nil {
		return fmt.Errorf("failed to setup route %s %s: %w", method, path, err)
	}
	return nil
}

func (s *Server) Start() error {

	if s.server != nil {
		return errors.New("[server] Already started")
	}

	if s.listener != nil {
		return errors.New("[server] Already listening")
	}

	address := s.listen

	if s.router == nil {
		s.InitSetup()
	}

	if err := s.FinalizeSetup(); err != nil {
		return fmt.Errorf("[server] failed to finalize the setup: %v", err)
	}

	if address == "" {
		address = ":http"
	}

	if s.serverFactory == nil {
		s.serverFactory = managers.NewHttpServerManager
	}

	s.server = s.serverFactory(address, s.router)
	if s.server == nil {
		return fmt.Errorf("[server] failed to initialize server for %s", address)
	}

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("[server] failed to listen the address %s: %v", address, err)
	}
	s.listener = &ln

	// Channel to signal the server start errors
	startErrors := make(chan error, 1)

	var logErrorHandler = func(err error) {
		log.Printf("[server] Could not start serving at %s: %v", address, err)
	}

	var channelErrorHandler = func(err error) {
		startErrors <- err
	}

	var startErrorHandler = channelErrorHandler

	go func() {
		if s.server != nil {
			log.Printf("[server] Starting server at %s", address)
			if err := s.server.Serve(ln); err != http.ErrServerClosed {
				startErrorHandler(err)
			}
		}
		log.Printf("[server] Stopped server at %s", address)
		s.server = nil
		startErrorHandler = nil
	}()

	// Wait for a short duration for server to start or fail
	select {
	case err := <-startErrors:
		startErrorHandler = nil
		if err != nil {
			return fmt.Errorf("[server] failed to start server: %v", err)
		}
	case <-time.After(DefaultStartDuration):
		// Continue if the server has not encountered immediate errors
		startErrorHandler = logErrorHandler
	}

	return nil
}

func (s *Server) Stop() error {
	if s.server != nil {
		if err := s.server.Shutdown(); err != nil {
			return fmt.Errorf("[server] Server shutdown error: %v", err)
		}
	}
	return nil
}

func (s *Server) SetListener(ln *net.Listener) error {
	if s.listener != nil {
		return fmt.Errorf("[server] Cannot set listener: Listening already")
	}
	s.listener = ln
	return nil
}

func (s *Server) UnSetup() error {

	if s.server != nil {
		return errors.New("[server] Cannot revert route setup, server still running")
	}

	if s.router != nil {
		s.router = nil
	}

	if s.swaggerManager != nil {
		s.swaggerManager = nil
	}

	return nil
}

// GetInternalHash returns a hash value representing the current internal state of the server
func (s *Server) GetInternalHash() (uint64, error) {
	if s.hashFactory == nil {
		s.hashFactory = fnv.New64a
	}
	return hashutils.ToUint64(
		fmt.Sprintf(
			"%s%p%p%p%p%p%p%p",
			s.listen, // string
			s.server, // pointer
			//s.repositoryControllerCollection, // Does not change
			s.router,          // pointer
			s.listener,        // pointer
			&s.swaggerManager, // pointer
			s.context,         // pointer
			&s.swaggerFactory, // pointer
			&s.serverFactory,  // pointer
		),
		s.hashFactory(),
	)
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
