package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

func (i Info) GetAffiliationInfo() string {
	return "have 31 division"
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Boy boy",
			Gender:  "male",
			Hobbies: []string{"reading book", "gaming", "coding"},
			Info:    Info{"wayne enterprise", "makassar city"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(rw, person); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
