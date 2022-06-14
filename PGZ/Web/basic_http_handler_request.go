package main

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, req.Method + "\n")
		io.WriteString(res, req.RemoteAddr + "\n")
	}

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handler,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
