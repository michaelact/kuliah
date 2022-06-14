package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	directory := http.Dir("./study")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/resources/", http.StripPrefix("/resources", fileserver))

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
