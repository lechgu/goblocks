package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed html/*
var html embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(html)
	contentStatic, _ := fs.Sub(fsys, "html")
	return http.FileServer(http.FS(contentStatic))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
