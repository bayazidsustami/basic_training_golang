package main

import (
	"github.com/labstack/echo/v4"
	"github.com/unrolled/secure"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/index", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.String(http.StatusOK, "hello")
	})

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:9000", "https://www.google.com"},
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))
}
