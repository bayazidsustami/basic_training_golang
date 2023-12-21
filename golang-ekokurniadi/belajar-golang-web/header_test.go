package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "PZN")
	fmt.Fprint(w, "Ok")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	headerPowered := response.Header.Get("X-Powered-By")

	fmt.Println(headerPowered)
}
