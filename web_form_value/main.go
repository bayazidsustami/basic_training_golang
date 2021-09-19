package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routSubmitPost)

	fmt.Println("server started at :9000")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(rw, nil)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(rw, "", http.StatusBadRequest)
}

func routSubmitPost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var message = r.Form.Get("message")

		var data = map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(rw, data); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(rw, "", http.StatusBadRequest)
}
