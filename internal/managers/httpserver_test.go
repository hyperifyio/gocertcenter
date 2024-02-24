// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers_test

import (
	"github.com/gorilla/mux"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"net"
	"net/http"
	"testing"
)

func TestHttpServerManager_Serve(t *testing.T) {
	router := mux.NewRouter()
	manager := managers.NewHttpServerManager("localhost:0", router)

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	go func() {
		err := manager.Serve(listener)
		if err != http.ErrServerClosed {
			t.Errorf("Expected server to close with http.ErrServerClosed, got %v", err)
		}
	}()

	// Shutdown the server to end the test
	if err := manager.Shutdown(); err != nil {
		t.Fatalf("Failed to shutdown server: %v", err)
	}

}

func TestHttpServerManager_Shutdown(t *testing.T) {
	router := mux.NewRouter()
	manager := managers.NewHttpServerManager("localhost:0", router)

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	go func() {
		if err := manager.Serve(listener); err != http.ErrServerClosed {
			t.Errorf("Expected server to close with http.ErrServerClosed, got %v", err)
		}
	}()

	if err := manager.Shutdown(); err != nil {
		t.Fatalf("Failed to shutdown server: %v", err)
	}
}
