package main

import (
	"text/template"
	"net/http"
	"embed"
	"os"
	"io"

	"github.com/rs/zerolog/log"
)

//go:embed templates/*.gohtml
var templates embed.FS
var allTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		allTemplates.ExecuteTemplate(res, "upload.gohtml", nil)
	})

	mux.HandleFunc("/upload", func(res http.ResponseWriter, req *http.Request) {
		file, fileHeader, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}

		fileDst, err := os.Create("./resources/" + fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		io.Copy(fileDst, file)
		allTemplates.ExecuteTemplate(res, "upload-success.gohtml", map[string]string{
			"Name": req.PostFormValue("name"),
			"File": "/static/" + fileHeader.Filename,
		})
	})

	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
