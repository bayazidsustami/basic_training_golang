package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type M map[string]interface{}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/list-files", handlerListFile)
	http.HandleFunc("/download", handlerDownload)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handlerIndex(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handlerListFile(rw http.ResponseWriter, r *http.Request) {
	files := []M{}
	basePath, _ := os.Getwd()
	fileLocation := filepath.Join(basePath, "files")

	err := filepath.Walk(fileLocation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, M{"filename": info.Name(), "path": path})
		return nil
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(files)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}

func handlerDownload(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	rw.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(rw, f); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
