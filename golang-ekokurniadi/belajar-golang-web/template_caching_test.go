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

//go:embed templates/*.gohtml
var templateToCache embed.FS

// auto caching parsing template
var mytemplate = template.Must(template.ParseFS(templateToCache, "templates/*.gohtml"))

func SimpleHTMLTemplateCaching(w http.ResponseWriter, r *http.Request) {
	mytemplate.ExecuteTemplate(w, "simple.gohtml", "hello from template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLTemplateCaching(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
