package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var filePath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filePath)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "learning golang web",
			"name":  "bay bay",
		}

		err = tmpl.Execute(rw, data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)

		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
