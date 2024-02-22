// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/mainutils"
	"github.com/hyperifyio/gocertcenter/internal/repositories/memoryRepository"
	"github.com/hyperifyio/gocertcenter/internal/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	listenPort = flag.String("port", mainutils.GetEnvOrDefault("PORT", "8080"), "port on which the server listens")
	certFile   = flag.String("cert", mainutils.GetEnvOrDefault("CERT_FILE", "cert.pem"), "server certificate as PEM file")
	keyFile    = flag.String("key", mainutils.GetEnvOrDefault("KEY_FILE", "key.pem"), "server private key as PEM file")
	caFile     = flag.String("ca", mainutils.GetEnvOrDefault("CA_FILE", "ca.pem"), "server CA as PEM file")
)

func main() {

	flag.Parse()

	var wg sync.WaitGroup

	listenAddr := fmt.Sprintf(":%s", *listenPort)

	repositoryCollection := memoryRepository.NewCollection()

	repositoryControllerCollection := modelcontrollers.NewControllerCollection(
		repositoryCollection.CertificateRepository,
		repositoryCollection.PrivateKeyRepository,
	)

	server := server.NewServer(listenAddr, *repositoryControllerCollection)

	shutdownHandler := func() error {
		server.Stop()
		return nil
	}

	// Setup signal handling for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("[main]: Starting server: %s", listenAddr)
	server.Start()

	<-shutdown
	log.Printf("[main]: Shutting down: %s", listenAddr)
	if err := shutdownHandler(); err != nil {
		log.Printf("[main]: Failed to close server: %v", err)
	}
	wg.Wait()

}
