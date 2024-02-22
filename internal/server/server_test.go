// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package server

import (
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/repositories/memoryRepository"
	"net/http"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	listen := "localhost:8080"
	mockControllerCollection := modelcontrollers.ControllerCollection{} // Assume this is a mock or a valid instance

	server := NewServer(listen, mockControllerCollection)
	if server.listen != listen {
		t.Errorf("NewServer listen = %s; want %s", server.listen, listen)
	}

	if server.repositoryControllerCollection != mockControllerCollection {
		t.Errorf("NewServer repositoryControllerCollection does not match the provided controller collection")
	}
}

// Assuming Start and Stop methods are updated to start/stop an HTTP server
func TestServer_StartStop(t *testing.T) {

	t.Skip("Skipping this test for now since not implemented.")

	listenAddr := "localhost:8080"

	repositoryCollection := memoryRepository.NewCollection()

	repositoryControllerCollection := modelcontrollers.NewControllerCollection(
		repositoryCollection.CertificateRepository,
		repositoryCollection.PrivateKeyRepository,
	)

	server := NewServer(listenAddr, *repositoryControllerCollection)

	server.Start()

	// Wait for the server to start
	time.Sleep(time.Second)

	// Attempt to connect to the server
	resp, err := http.Get("http://" + listenAddr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	resp.Body.Close() // Close the body to avoid resource leaks

	// Now test stopping the server
	server.Stop()

	// Attempt to connect to the server after stopping it
	_, err = http.Get("http://" + listenAddr)
	if err == nil {
		t.Fatal("Server should not be reachable after stopping")
	}
}
