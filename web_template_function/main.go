package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const view string = `<html>
	<head>
		<title>learning golang</title>
	</head>
	<body>
		<h1>Hello from string golang</h1>
	</body>
</html>`

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

	var funcMap = template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"avg": func(n ...int) int {
			var total = 0
			for _, each := range n {
				total += each
			}
			return total / len(n)
		},
	}

	http.HandleFunc("/custom_function", func(rw http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("view_custom_function.html").
			Funcs(funcMap).
			ParseFiles("view_custom_function.html"))
		if err := tmpl.Execute(rw, nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/other_index", handlerOtherIndex)
	http.HandleFunc("/html_string", handlerHtmlString)
	http.HandleFunc("/other_html/redirect", handlerRedirect)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handlerIndex(rw http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("index").ParseFiles("view_render_spesific_template.html"))
	if err := tmpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handlerOtherIndex(rw http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("other-index").ParseFiles("view_render_spesific_template.html"))
	if err := tmpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handlerHtmlString(rw http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("main-template").Parse(view))
	if err := tmpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handlerRedirect(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/html_string", http.StatusTemporaryRedirect)
}
