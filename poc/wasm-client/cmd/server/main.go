package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var webContent embed.FS

func main() {

	// Create an http.FileSystem from the embedded files.
	// The "web" subdirectory becomes the root of this file system.
	contentFS, err := fs.Sub(webContent, "web")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(contentFS)))

	log.Println("Listening on http://localhost:8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
