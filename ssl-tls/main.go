package main

import (
	"log"
	"net/http"
)

func StartNotTLSServer() {
	mux := new(http.ServeMux)
	mux.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("redirecting to https://localhost")
		http.Redirect(writer, request, "https://localhost", http.StatusTemporaryRedirect)
	}))

	log.Fatal(http.ListenAndServe(":80", mux))
}

func main() {
	go StartNotTLSServer()

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	log.Println("Server started at : 443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}
}
