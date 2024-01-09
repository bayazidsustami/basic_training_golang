package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoutingHelloWorld(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", string(byte))

}

func TestFiberCtx(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")

		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest(http.MethodGet, "/hello?name=bay", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello bay", string(byte))

}

func TestFiberCtxDefaultQuery(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")

		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(byte))

}
