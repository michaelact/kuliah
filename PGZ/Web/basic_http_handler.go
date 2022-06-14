package main

import (
	"net/http"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "<h1> Michael Act </h1>")
	}

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handler,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
