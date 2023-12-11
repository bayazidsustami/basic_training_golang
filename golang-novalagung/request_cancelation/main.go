package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.ListenAndServe(":8080", nil)
}

func handlerIndex(rw http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(body))
		time.Sleep(10 * time.Second)

		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}
