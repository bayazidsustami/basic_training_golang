package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/save", handlerSave)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Printf("server started at localhost :9000")
	http.ListenAndServe(":9000", nil)
}

func handlerIndex(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handlerSave(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf(
			"hello my name is %s. i'm %d year old %s",
			payload.Name,
			payload.Age,
			payload.Gender,
		)

		rw.Write([]byte(message))
		return
	}

	http.Error(rw, "only accept post request", http.StatusBadRequest)
}
