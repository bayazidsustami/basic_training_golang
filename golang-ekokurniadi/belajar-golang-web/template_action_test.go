package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", map[string]any{
		"Name": "bay",
	})
}

func TemplateIfComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]any{
		"FinalValue": 50,
	})
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]any{
		"Hobbies": []string{
			"Makan", "Tidur", "Salto",
		},
	})
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(w, "with.gohtml", Page{
		Title: "Template With",
		Name:  "Bay",
		Address: Address{
			City:   "Makassar",
			Street: "Jalan Jalan",
		},
	})
}

func TestSimpleHTMLIfTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLIfComparatorTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIfComparator(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLRangeTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLWithTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
