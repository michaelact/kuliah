package main

import (
	"net/http"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		firstname := req.PostFormValue("firstname")
		lastname := req.PostFormValue("lastname")

		if (firstname != "") && (lastname != "") {
			res.WriteHeader(http.StatusOK)
			io.WriteString(res, "<h1>" + firstname + " " + lastname + "</h1>")
		} else {
			res.WriteHeader(http.StatusBadRequest)
			io.WriteString(res, "Can't understand.")
		}
	}

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handler,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
