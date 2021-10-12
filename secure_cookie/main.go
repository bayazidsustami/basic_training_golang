package main

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	echo "github.com/labstack/echo/v4"
)

type M map[string]interface{}

var sc = securecookie.New([]byte("very-secret"), []byte("a-lot-secret-yay"))

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":9000"))

}

func setCookie(ctx echo.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(ctx.Response(), cookie)

	return nil
}
