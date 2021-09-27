package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"status_code": 200, "message": "success"}
		return ctx.JSON(http.StatusOK, data)
	})

	//parsing request
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s i have message for you : %s", name, message)
		dataJson := M{
			"respose_code": 200,
			"name":         name,
			"message":      data,
		}
		return ctx.JSON(http.StatusOK, dataJson)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf("Hello %s i have message for you %s", name, strings.Replace(message, "/", "", 1))
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}
