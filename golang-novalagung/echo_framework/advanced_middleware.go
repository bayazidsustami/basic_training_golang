package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//middleware
	e.Use(middlewareOne)
	e.Use(middlewareTwo)
	e.Use(echo.WrapMiddleware(middlewareSomething))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/index", actionIndex)

	e.Logger.Fatal(e.Start(":9000"))
}

func actionIndex(ctx echo.Context) error {
	fmt.Println("threee!")
	return ctx.JSON(http.StatusOK, true)
}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware one")
		return next(c)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware two")
		return next(c)
	}
}

func middlewareSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("from middleware something")
		next.ServeHTTP(rw, r)
	})
}
