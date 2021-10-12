package main

import echo "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":9000"))

}
