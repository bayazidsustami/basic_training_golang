package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said : \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Latern"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(rw, person); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:900")
	http.ListenAndServe(":9000", nil)
}
