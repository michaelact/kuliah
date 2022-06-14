package main

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		cookie := http.Cookie{
			Name:  "_SecureToken", 
			Value: req.URL.Query().Get("name"), 
			Path:  "/", 
		}

		http.SetCookie(res, &cookie)
		io.WriteString(res, "<h1> Successfully create cookie. </h1>")
	})

	mux.HandleFunc("/dashboard", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("_SecureToken")

		if err != nil {
			io.WriteString(res, "<h1> Cookie not found. Please generate via /login path. </h1>")
		} else {
			io.WriteString(res, "<h1> " + cookie.Value + " </h1>")
		}
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
