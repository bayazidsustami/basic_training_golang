package main

import (
	"fmt"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var cookieName = "CookieData"

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(rw http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"
	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(rw, c)
	}

	rw.Write([]byte(c.Value))
}

func ActionDelete(rw http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(rw, c)

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}
