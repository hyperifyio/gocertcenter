// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package main

import (
	"flag"
	"github.com/hyperifyio/gocertcenter/internal/mainutils"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	listenPort = flag.String("port", mainutils.GetEnvOrDefault("PORT", "8080"), "port on which the server listens")
	certFile   = flag.String("cert", mainutils.GetEnvOrDefault("GOCERTCENTER_CERT_FILE", "cert.pem"), "server certificate as PEM file")
	keyFile    = flag.String("key", mainutils.GetEnvOrDefault("GOCERTCENTER_KEY_FILE", "key.pem"), "server private key as PEM file")
	caFile     = flag.String("ca", mainutils.GetEnvOrDefault("GOCERTCENTER_CA_FILE", "ca.pem"), "server CA as PEM file")
)

func main() {

	flag.Parse()

	var wg sync.WaitGroup

	//listenAddr := fmt.Sprintf(":%s", *listenPort)
	//listenerTlsConfig := tlsutils.LoadTLSConfig(*certFile, *keyFile, *caFile)

	shutdownHandler := func() error {

		return nil
	}

	// Setup signal handling for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	log.Println("Shutting down...")
	if err := shutdownHandler(); err != nil {
		log.Printf("[main]: Failed to close server: %v", err)
	}
	wg.Wait()

}
