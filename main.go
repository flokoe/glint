// Clairvoyance - A web based AI chat UI
// Copyright (C) 2025 Florian Köhler
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see
// <https://www.gnu.org/licenses/>.
//
// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("clairvoyance.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database:", err)
	}
	MigrateAndSeed(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(DBMiddleware(db))

	e.Static("/", "public")

	t := &templateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", newChat)
	e.GET("/admin", admin)

	e.POST("/api/admin/adapters", addAdapter)

	e.Logger.Fatal(e.Start(":1323"))
}
