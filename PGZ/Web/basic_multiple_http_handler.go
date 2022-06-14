package main

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	// Priority: 0
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "<h1> Michael Act </h1>")
	})

	// Priority: 1
	mux.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "<h1> Ini Testing </h1>")
	})

	// Priority: 1
	mux.HandleFunc("/all/thumb/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "<h1> Semua kesini | Thumb </h1>")
	})

	// Priority: 2
	mux.HandleFunc("/all/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "<h1> Semua kesini </h1>")
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
