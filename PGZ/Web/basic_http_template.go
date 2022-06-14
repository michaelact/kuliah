package main

import (
	"text/template"
	"net/http"
	"embed"

	"github.com/rs/zerolog/log"
)

//go:embed templates/*.gohtml
var templates embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/string", func(res http.ResponseWriter, req *http.Request) {
		templateHtml := `<html><title>{{.}}</title><h1>{{.}}</h1></html>`
		t, err := template.New("SIMPLE").Parse(templateHtml)
		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(res, "SIMPLE", "Michael Act Disini | By String")
	})

	mux.HandleFunc("/single", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseGlob("./templates/simple.gohtml")
		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(res, "simple.gohtml", "Michael Act Disini | By Single File")
	})

	mux.HandleFunc("/glob", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(res, "simple.gohtml", "Michael Act Disini | By GLOB")
	})

	mux.HandleFunc("/embed", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFS(templates, "templates/*.gohtml")
		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(res, "simple.gohtml", "Michael Act Disini | By Embed")
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
