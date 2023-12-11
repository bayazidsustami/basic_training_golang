package main

import (
	"fmt"
	"log"
	"net/http"
	config "simplecofig/conf"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if config.Configuration().Log.Verbose {
		log.Println("incoming request from", r.Host, "accesing", r.URL.String())
	}

	c.ServeMux.ServeHTTP(rw, r)
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello bosqu"))
	})

	router.HandleFunc("/greeting", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello helloooo"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = config.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = config.Configuration().Server.WriteTimeout * time.Second

	server.Addr = fmt.Sprintf(":%d", config.Configuration().Server.Port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
