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
		Adapters []Adapter
		Models   []Model
	}

	var data AdminData

	// Fetch the newly created adapter with its models
	if err := db.Find(&data.Adapters).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch adapter")
	}

	return c.Render(http.StatusOK, "admin.html", data)
}

type adapterDTO struct {
	Type string `form:"adapterType"`
	URL  string `form:"adapterUrl"`
}

func addAdapter(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	a := new(adapterDTO)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Load into separate struct for security
	adapter := Adapter{
		Type: a.Type,
		URL:  a.URL,
	}

	if err := db.Create(&adapter).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to create adapter")
	}

	// Fetch the newly created adapter with its models
	var adapters []Adapter
	if err := db.Find(&adapters).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch adapter")
	}

	return c.Render(http.StatusOK, "adapterTable.html", adapters)
}
