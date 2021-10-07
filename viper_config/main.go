package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/index", func(ctx echo.Context) (err error) {
		return ctx.JSON(http.StatusOK, true)
	})

	e.Logger.Print("starting", viper.GetString("appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}
