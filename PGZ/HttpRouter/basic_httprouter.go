package main

import (
	"net/http"
	"embed"
	"io/fs"
	"fmt"
	"io"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

//go:embed resources/*.txt
var resources embed.FS

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("x-powered-by", "michaelact")
	middleware.Handler.ServeHTTP(res, req)
	fmt.Println("Executed!")
}

func main() {
	router := httprouter.New()

	// Handle panic handler
	router.PanicHandler = func(res http.ResponseWriter, req *http.Request, error interface{}) {
		fmt.Fprintf(res, "Server sedang sibuk. \n%s", error)
	}

	// Handler not found path
	router.NotFound = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Gak ketemu gan: \n%s", req.URL.Path)
	})

	// Handler restrict HTTP method
	router.MethodNotAllowed = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "HTTP method tidak di izinkan: \n%s", req.Method)
	})
	router.POST("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		res.Header().Add("x-powered-by", "michaelact")
		io.WriteString(res, "HTTP method di izinkan")
	})

	// Named Parameter
	router.GET("/class/:classId", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		res.Header().Add("x-powered-by", "michaelact")
		io.WriteString(res, "Class: " + params.ByName("classId"))
	})
	router.GET("/class/:classId/chapter/:chapterId", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		res.Header().Add("x-powered-by", "michaelact")

		classId := params.ByName("classId")
		chapterId := params.ByName("chapterId")
		fmt.Fprintf(res, "Class: %s \nChapter: %s", classId, chapterId)
	})

	// Catch All Parameter
	router.GET("/images/*imagePath", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		res.Header().Add("x-powered-by", "michaelact")

		imagePath := params.ByName("imagePath")
		fmt.Fprintf(res, "Here is you are: %s", imagePath)
	})

	// Serve File
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/static/*filepath", http.FS(directory))

	// Panicccccc!
	router.GET("/panic", func(res http.ResponseWriter, req *http.Request, params httprouter.Params){
		panic("Gagallll!")
	})

	middleware := new(LogMiddleware)
	middleware.Handler = router

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: middleware,
	}

	// Log error if failed to start
	log.Fatal().Err(server.ListenAndServe())
}
