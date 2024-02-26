// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hyperifyio/gocertcenter/internal/app/appendpoints"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apiserver"
	"github.com/hyperifyio/gocertcenter/internal/common/mainutils"
)

var (
	listenPort = flag.String("port", mainutils.GetEnvOrDefault("PORT", "8080"), "port on which the server listens")
	dataDir    = flag.String("data-dir", mainutils.GetEnvOrDefault("DATA_DIR", "./tmp/data"), "application data directory")
)

func main() {

	flag.Parse()

	var wg sync.WaitGroup

	listenAddr := fmt.Sprintf(":%s", *listenPort)

	// fileManager := managers.NewFileManager()
	// randomManager := managers.NewRandomManager()
	// certManager := managers.NewCertificateManager(randomManager)
	// repository := filerepository.NewCollection(certManager, fileManager, dataDir)
	// expiration := 24 * time.Hour
	//
	// repositoryControllerCollection := appcontrollers.NewCollection(
	// 	appcontrollers.NewOrganizationController(repository.Organization),
	// 	appcontrollers.NewCertificateController(repository.Certificate, certManager, randomManager, expiration),
	// 	appcontrollers.NewPrivateKeyController(repository.PrivateKey),
	// )

	server, err := apiserver.NewServer(listenAddr, nil)
	if err != nil {
		log.Fatalf("[main]: Failed to create the server: %v", err)
	}

	server.SetInfo(appendpoints.GetInfo())

	if err := server.SetupRoutes(appendpoints.GetRoutes()); err != nil {
		log.Fatalf("[main]: Failed to setup routes: %v", err)
	}

	shutdownHandler := func() error {
		if err := server.Stop(); err != nil {
			log.Printf("[main]: Failed to stop server: %v", err)
		}
		return nil
	}

	// Setup signal handling for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("[main]: Starting server: %s", server.GetAddress())
	if err := server.Start(); err != nil {
		log.Printf("[main]: Failed to start server: %v", err)
	}

	<-shutdown
	log.Printf("[main]: Shutting down: %s", server.GetAddress())
	if err := shutdownHandler(); err != nil {
		log.Printf("[main]: Failed to close server: %v", err)
	}
	wg.Wait()

}
