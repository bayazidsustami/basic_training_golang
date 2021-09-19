package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerDefault)
	startServer()
}

func startServer() {
	fmt.Println("server started at :9000")
	http.ListenAndServe(":9000", nil)
}

func handlerDefault(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rw.Write([]byte("post"))
	case "GET":
		rw.Write([]byte("get"))
	default:
		http.Error(rw, "", http.StatusBadRequest)
	}
}
