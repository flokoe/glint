/*
Clairvoyance - A feature-rich web based AI chat UI
Copyright (C) 2025 Florian Köhler

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

SPDX-License-Identifier: AGPL-3.0-only
*/

// Clairvoyance is a feature-rich web based AI chat UI.
// It provides a user-friendly interface for interacting with AI models.
package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/flokoe/clairvoyance/internal/database"
	"github.com/flokoe/clairvoyance/internal/route"
)

// DBMiddleware make DB available in the context so we can use it in the handlers.
func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

type templateRegistry struct {
	templates *template.Template
}

// Implement the echo.Renderer interface.
func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("clairvoyance.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database:", err)
	}
	database.MigrateAndSeed(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(DBMiddleware(db))

	e.Static("/", "public")

	t := &templateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", route.NewChatView)
	e.GET("/admin", route.AdminView)

	e.POST("/api/admin/providers", route.ApiAdminAddProvider)

	e.Logger.Fatal(e.Start(":1323"))
}
