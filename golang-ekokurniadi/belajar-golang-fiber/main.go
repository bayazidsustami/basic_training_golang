package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	if fiber.IsChild() {
		fmt.Println("this is chiled")
	} else {
		fmt.Println("this is parent")
	}

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
