package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type M map[string]interface{}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9000"

	log.Println("Starting server at :9000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

}

func ActionData(rw http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request with method", r.Method)

	if r.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusBadRequest)
		return
	}

	payload := make(M)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := payload["name"]; !ok {
		http.Error(rw, "payload name is required", http.StatusBadRequest)
		return
	}

	data := M{
		"Message": fmt.Sprintf("Hello %s", payload["name"]),
		"Status":  true,
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
