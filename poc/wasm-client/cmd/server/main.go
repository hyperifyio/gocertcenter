package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/webview/webview_go"
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
	go func() {
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Open a webview window
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Your Application")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8080")
	w.Run()

}
