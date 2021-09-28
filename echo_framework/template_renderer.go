package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplate()

	return tpl
}

func (t *Renderer) ReloadTemplate() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	ctx echo.Context,
) error {
	if t.debug {
		t.ReloadTemplate()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	router := echo.New()

	router.Renderer = NewRenderer("./*.html", true)

	router.GET("/index", func(ctx echo.Context) error {
		data := M{"message": "Hello World"}
		return ctx.Render(http.StatusOK, "index.html", data)
	})

	router.Logger.Fatal(router.Start(":9000"))
}
