// glint - An AI chat web UI
package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template is a custom HTML template renderer for Echo framework.
type Template struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface.
func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	e.GET("/", newChat)

	e.Logger.Fatal(e.Start(":1323"))
}
