package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func newChat(c echo.Context) error {
	return c.Render(http.StatusOK, "new-chat", "Sun")
}
