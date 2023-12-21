package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello aja")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func SayHelloMultiple(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	if firstName == "" || lastName == "" {
		fmt.Fprint(w, "hello aja")
	} else {
		fmt.Fprintf(w, "Hello %s+%s", firstName, lastName)
	}
}

func MultipleValuesQueryParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var names []string = query["name"]
	fmt.Fprint(w, strings.Join(names, ","))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/hello?name=bays", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/hello?first_name=bays&last_name=bay", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultiple(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func TestMultipleValuesQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/hello?name=bays&name=bay", nil)
	recorder := httptest.NewRecorder()

	MultipleValuesQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
