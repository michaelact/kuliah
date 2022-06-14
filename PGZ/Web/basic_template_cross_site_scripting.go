package main

import (
	"html/template"
	"net/http"
	"embed"

	"github.com/rs/zerolog/log"
)

//go:embed templates/*.gohtml
var templates embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/embed", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFS(templates, "templates/*.gohtml")
		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(res, "simple.gohtml", req.URL.Query().Get("text"))
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
