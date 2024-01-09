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

	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("middleware api before")
		err := ctx.Next()
		fmt.Println("middleware api after")
		return err
	})

	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Println("middleware before")
		err := ctx.Next()
		fmt.Println("middleware after")
		return err
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello API's")
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
