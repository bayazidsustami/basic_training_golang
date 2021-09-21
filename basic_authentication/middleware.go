package main

import "net/http"

const USERNAME = "batmam"
const PASSWORD = "secret"

func Auth(rw http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		rw.Write([]byte(`something wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		rw.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGet(rw http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		rw.Write([]byte("only get is allowed"))
		return false
	}

	return true
}
