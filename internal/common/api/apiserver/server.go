// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package apiserver

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/hashutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apierrors"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

const DefaultHostname = "localhost"
const DefaultProtocol = "http"
const DefaultOpenApiTitle = "Example API"
const DefaultOpenApiVersion = "0.0.0"

// DefaultStartDuration This is a time to wait for any possible errors while starting the server
const DefaultStartDuration = 500 * time.Millisecond

// ApplicationServer implements apitypes.Server
type ApplicationServer struct {
	listen         string
	server         managers.ServerManager
	router         *mux.Router
	listener       *net.Listener
	context        *context.Context
	swaggerManager managers.SwaggerManager
	swaggerFactory apitypes.NewSwaggerManagerFunc
	serverFactory  apitypes.NewServerManagerFunc
	info           *openapi3.Info
	hashFactory    apitypes.Hash64FactoryFunc
}

func (s *ApplicationServer) IsStarted() bool {
	return s.server != nil
}

func (s *ApplicationServer) GetAddress() string {
	return s.listen
}

func (s *ApplicationServer) SetInfo(info *openapi3.Info) {
	s.info = info
}

func (s *ApplicationServer) GetInfo() *openapi3.Info {
	if s.info == nil {
		s.info = &openapi3.Info{
			Title:   DefaultOpenApiTitle,
			Version: DefaultOpenApiVersion,
		}
	}
	return s.info
}

func (s *ApplicationServer) GetURL() string {
	url := s.listen
	if strings.HasPrefix(url, ":") {
		return fmt.Sprintf("%s://%s%s", DefaultProtocol, DefaultHostname, url)
	}
	return fmt.Sprintf("%s://%s", DefaultProtocol, url)
}

func (s *ApplicationServer) SetSwaggerFactory(factory apitypes.NewSwaggerManagerFunc) {
	s.swaggerFactory = factory
}

func (s *ApplicationServer) SetServerFactory(factory apitypes.NewServerManagerFunc) {
	s.serverFactory = factory
}

func (s *ApplicationServer) InitSetup() error {

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
			"ApplicationServer location",
			s.GetInfo(),
		)
		if err != nil {
			return fmt.Errorf("[server] failed to create swagger router: %v", err)
		}
		s.swaggerManager = swaggerManager
	}

	return nil

}

func (s *ApplicationServer) SetupRoutes(
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

func (s *ApplicationServer) SetupNotFoundHandler(handler apitypes.RequestHandlerFunc) {
	s.router.NotFoundHandler = ResponseHandler(handler)
}

func (s *ApplicationServer) SetupMethodNotAllowedHandler(handler apitypes.RequestHandlerFunc) {
	s.router.MethodNotAllowedHandler = ResponseHandler(handler)
}

func (s *ApplicationServer) FinalizeSetup() error {
	if err := s.swaggerManager.GenerateAndExposeOpenapi(); err != nil {
		return fmt.Errorf("failed to initialize OpenAPI integration: %v", err)
	}
	return nil
}

func (s *ApplicationServer) SetupHandler(
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
	handler := ResponseHandler(apiHandler)
	if _, err := s.swaggerManager.AddRoute(method, path, handler, definitions); err != nil {
		return fmt.Errorf("failed to setup route %s %s: %w", method, path, err)
	}
	return nil
}

func (s *ApplicationServer) Start() error {

	if s.server != nil {
		return errors.New("[server] Already started")
	}

	if s.listener != nil {
		return errors.New("[server] Already listening")
	}

	address := s.listen

	if s.router == nil {
		if err := s.InitSetup(); err != nil {
			return fmt.Errorf("[server] failed to initialize: %v", err)
		}
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

func (s *ApplicationServer) Stop() error {
	if s.server != nil {
		if err := s.server.Shutdown(); err != nil {
			return fmt.Errorf("[server] ApplicationServer shutdown error: %v", err)
		}
	}
	return nil
}

func (s *ApplicationServer) SetListener(ln *net.Listener) error {
	if s.listener != nil {
		return fmt.Errorf("[server] Cannot set listener: Listening already")
	}
	s.listener = ln
	return nil
}

func (s *ApplicationServer) UnSetup() error {

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
func (s *ApplicationServer) GetInternalHash() (uint64, error) {
	if s.hashFactory == nil {
		s.hashFactory = fnv.New64a
	}
	return hashutils.ToUint64(
		fmt.Sprintf(
			"%s%p%p%p%p%p%p%p",
			s.listen, // string
			s.server, // pointer
			// s.repositoryControllerCollection, // Does not change
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

// NewServer ..
func NewServer(
	listen string,
	swaggerFactory apitypes.NewSwaggerManagerFunc,
) (*ApplicationServer, error) {
	s := NewUninitializedServer(listen, swaggerFactory)
	if err := s.InitSetup(); err != nil {
		return nil, fmt.Errorf("[server] server initialization failed: %v", err)
	}
	s.SetupNotFoundHandler(apierrors.NotFound)
	s.SetupMethodNotAllowedHandler(apierrors.MethodNotAllowed)
	return s, nil
}

func NewUninitializedServer(
	listen string,
	swaggerFactory apitypes.NewSwaggerManagerFunc,
) *ApplicationServer {
	s := &ApplicationServer{
		listen:         listen,
		swaggerFactory: swaggerFactory,
		serverFactory:  nil,
		swaggerManager: nil,
	}
	return s
}

var _ apitypes.Server = (*ApplicationServer)(nil)
