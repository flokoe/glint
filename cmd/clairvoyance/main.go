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
	"database/sql"
	_ "embed"
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"

	"github.com/flokoe/clairvoyance/internal/handler"
	"github.com/flokoe/clairvoyance/internal/llm"
)

type templateRegistry struct {
	templates *template.Template
}

// Implement the echo.Renderer interface.
func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var ddl string

func main() {
	db, err := sql.Open("sqlite", "clairvoyance.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	LLMProvider, err := llm.NewProvider("openai-compatible", map[string]string{
		"base_url": "http://127.0.0.1:8080",
	})
	if err != nil {
		log.Fatal("Failed to create LLM provider:", err)
	}

	chatHandler := handler.NewChatHandler(db, LLMProvider)
	adminHandler := handler.NewAdminHandler(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "static")

	t := &templateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", chatHandler.ChatView)
	e.GET("/admin", adminHandler.AdminView)

	e.POST("/api/completion", chatHandler.ApiCompletion)

	e.POST("/api/admin/providers", adminHandler.ApiAddProvider)

	e.Logger.Fatal(e.Start(":1323"))
}
