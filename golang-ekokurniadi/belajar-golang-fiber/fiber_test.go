package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app *fiber.App = fiber.New()

func TestRoutingHelloWorld(t *testing.T) {
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

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")   //header
		last := c.Cookies("lastname") //cookies
		return c.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest(http.MethodGet, "/request", nil)
	request.Header.Set("firstname", "bay")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "bayazid"})
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello bay bayazid", string(byte))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " from user " + userId)
	})

	request := httptest.NewRequest(http.MethodGet, "/users/2/orders/1", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 1 from user 2", string(byte))

}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	bodyRequest := strings.NewReader("name=bay")
	request := httptest.NewRequest(http.MethodPost, "/hello", bodyRequest)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello bay", string(byte))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("upload success")
	})

	bodyRequest := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyRequest)
	file, _ := writer.CreateFormFile("file", "contoh.txt")
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/upload", bodyRequest)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "upload success", string(byte))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()
		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}
		return c.SendString("Hello " + request.Username)
	})

	bodyRequest := strings.NewReader(`{ "username": "bay", "password": "ini password"}`)
	request := httptest.NewRequest(http.MethodPost, "/login", bodyRequest)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello bay", string(byte))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := c.BodyParser(request)
		if err != nil {
			return err
		}
		return c.SendString("Register success " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	bodyRequest := strings.NewReader(`{ "username": "bay", "password": "ini password", "name": "baybay"}`)
	request := httptest.NewRequest(http.MethodPost, "/register", bodyRequest)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success bay", string(byte))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	bodyRequest := strings.NewReader(`username=bay&password=rahasia&name=baybay`)
	request := httptest.NewRequest(http.MethodPost, "/register", bodyRequest)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success bay", string(byte))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	bodyRequest := strings.NewReader(
		`<RegisterRequest>
			<username>bay</username>
			<password>rahasia</password>
			<name>baybay</name>
		</RegisterRequest>`)
	request := httptest.NewRequest(http.MethodPost, "/register", bodyRequest)
	request.Header.Set("Content-Type", "application/xml")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register success bay", string(byte))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "bay",
			"name":     "bay bay",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"bay bay","username":"bay"}`, string(byte))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/contoh.txt", "contoh.txt")
	})

	request := httptest.NewRequest(http.MethodGet, "/download", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "attachment; filename=\"contoh.txt\"", response.Header.Get("Content-Disposition"))

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "this is sample file", string(byte))
}

func TestRoutingGroup(t *testing.T) {
	helloWord := func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWord)
	api.Get("/world", helloWord)

	web := app.Group("/web")
	web.Get("/hello", helloWord)
	api.Get("/world", helloWord)

	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(byte))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	request := httptest.NewRequest(http.MethodGet, "/public/contoh.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "this is sample file", string(byte))
}
