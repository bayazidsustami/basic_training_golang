package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mytemplates = template.Must(template.New("post.gothml").ParseFiles("./templates/post.gohtml"))

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {

	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body":  "<p>Ini adalah body</body>",
	})
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {

	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body":  template.HTML("<p>Ini adalah body</body>"),
	})
}

func TemplateAutoEscapeXss(w http.ResponseWriter, r *http.Request) {

	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestSimpleHTMLAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
func TestSimpleHTMLXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeXss(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
