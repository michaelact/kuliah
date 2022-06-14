package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:3000",
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
