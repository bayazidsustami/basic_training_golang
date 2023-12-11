package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func main() {
	/*
		http.HandleFunc("/index", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Access-Control-Allow-Origin", "*") //allow all origin
			rw.Header().Set("Access-Control-Allow-Method", "OPTIONS, GET, POST, PUT")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

			if r.Method == "OPTIONS" {
				rw.Write([]byte("allowed"))
				return
			}

			rw.Write([]byte("hello"))
		})
	*/
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://www.google.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/index", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":8000"))

}
