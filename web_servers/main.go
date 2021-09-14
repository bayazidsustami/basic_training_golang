package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Halo")
	})

	http.HandleFunc("/index", index)
	http.HandleFunc("/greeting", showGreetingWithName)

	fmt.Println("starting web server as http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar")
}

func showGreetingWithName(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "bay bay",
		"Message": "hello all",
	}

	var t, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}
