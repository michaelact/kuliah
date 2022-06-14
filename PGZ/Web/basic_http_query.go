package main

import (
	"net/http"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()

		// Example URL: http://localhost:3000/?message=halo&message=juga
		messages := query["message"]
		if len(messages) > 0 {
			io.WriteString(res, "<h1> " + messages[0] + messages[1] + " </h1>")
		} else {
			io.WriteString(res, "<h1> No Message. </h1>")
		}
	}

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handler,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
