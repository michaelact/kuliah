package golang_web

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"io"

	"github.com/stretchr/testify/assert"
)

func MyHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1> Michael Act </h1>")
}

func TestMyHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	MyHandler(recorder, request)

	response := recorder.Result()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "<h1> Michael Act </h1>", string(data))
}
