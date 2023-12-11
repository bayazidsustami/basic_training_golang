package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routSubmitPost)
	http.HandleFunc("/input_file", routeInputFile)
	http.HandleFunc("/process_file", routeProcessInputFile)

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

func routeInputFile(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view_form_file.html"))
	var err = tmpl.Execute(rw, nil)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func routeProcessInputFile(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias")

	uploadFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := handler.Filename
	if alias != "" {
		fileName = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", fileName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadFile); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("done"))
}
