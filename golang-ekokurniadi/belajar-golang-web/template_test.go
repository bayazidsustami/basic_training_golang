package belajargolangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello from simple html template")
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "hello from template")
}

func SimpleHTMLDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "hello from template")
}

//go:embed templates/*.gohtml
var templates embed.FS

func SimpleHTMLDirectoryEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "hello from template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDirectory(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLDirectoryEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDirectoryEmbed(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
