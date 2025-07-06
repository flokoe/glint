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
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/flokoe/clairvoyance/internal/model"
)

type templateRegistry struct {
	templates *template.Template
}

// Implement the echo.Renderer interface.
func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newChat(c echo.Context) error {
	return c.Render(http.StatusOK, "chat.html", "Sun")
}

func admin(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	type AdminData struct {
		Providers []model.Provider
		LLMs      []model.Llm
	}

	var data AdminData

	// Fetch the newly created provider with its models
	if err := db.Find(&data.Providers).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch provider")
	}

	return c.Render(http.StatusOK, "admin.html", data)
}

type providerDTO struct {
	Type string `form:"providerType"`
	URL  string `form:"providerUrl"`
}

func addProvider(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	a := new(providerDTO)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Load into separate struct for security
	provider := model.Provider{
		Type: a.Type,
		URL:  a.URL,
	}

	if err := db.Create(&provider).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to create provider")
	}

	// Fetch the newly created provider with its models
	var providers []model.Provider
	if err := db.Find(&providers).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch provider")
	}

	return c.Render(http.StatusOK, "providerTable.html", providers)
}
