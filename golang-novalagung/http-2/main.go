package main

import (
	"fmt"
	"log"
	"net/http"
)

const indexHTML = `
    <!DOCTYPE html>
    <html>
        <head>
            <title>Hello World</title>
            <script src="/static/app.js"></script>
            <link rel="stylesheet" href="/static/app.css"">
        </head>
        <body>
        Hello, gopher!<br>
        <img src="https://blog.golang.org/go-brand/logos.jpg" height="100">
        </body>
    </html>
`

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("."))))

	http.HandleFunc("/", index)

	log.Println("Server started at : 9000")
	err := http.ListenAndServeTLS(":9000", "server.crt", "server.key", nil)
	if err != nil {
		panic(err)
	}
}

func index(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(rw, r)
		return
	}

	if pusher, ok := rw.(http.Pusher); ok {
		if err := pusher.Push("/static/app.js", nil); err != nil {
			log.Printf("Failed to Push : %v", err)
		}

		if err := pusher.Push("/static/app.css", nil); err != nil {
			log.Printf("Failed to Push : %v", err)
		}
	}

	fmt.Fprintf(rw, indexHTML)
}
