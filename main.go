// glint - An AI chat web UI
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<!doctype html><html><head></head><body><h1>Hello World!</h1></body></html>")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
