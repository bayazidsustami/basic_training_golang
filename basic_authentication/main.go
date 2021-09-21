package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(rw http.ResponseWriter, r *http.Request) {
	if !Auth(rw, r) {
		return
	}

	if !AllowOnlyGet(rw, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJson(rw, SelectStudent(id))
		return
	}

	OutputJson(rw, GetStudents())
}

func OutputJson(rw http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}
