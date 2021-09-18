package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server startted at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}
