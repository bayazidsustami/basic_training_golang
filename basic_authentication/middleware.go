package main

import "net/http"

const USERNAME = "batmam"
const PASSWORD = "secret"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			rw.Write([]byte(`something wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			rw.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(rw, r)
	})
}

func MiddlewareOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			rw.Write([]byte("only get is allowed"))
			return
		}

		next.ServeHTTP(rw, r)
	})
}
