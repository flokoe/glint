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

package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/flokoe/clairvoyance/internal/llm"
	"github.com/flokoe/clairvoyance/internal/model"
)

func AdminView(c echo.Context) error {
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

	if err := db.Find(&data.LLMs).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch provider")
	}

	return c.Render(http.StatusOK, "admin.html", data)
}

type providerDTO struct {
	Type string `form:"providerType"`
	URL  string `form:"providerUrl"`
}

func ApiAdminAddProvider(c echo.Context) error {
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

	foo, err := llm.NewProvider(provider.Type, map[string]string{
		"base_url": provider.URL,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to create LLM provider")
	}

	m, err := foo.GetModels()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch models")
	}

	fmt.Println(provider.ID)

	var M []model.Llm
	for _, mdl := range *m {
		M = append(M, model.Llm{
			String:      mdl.String,
			Name:        mdl.String,
			ProviderID:  provider.ID,
			ContextSize: mdl.ContextSize,
			IsEnabled:   true,
		})
	}

	// Also save to database
	if err := db.Create(&M).Error; err != nil {
		return c.String(http.StatusInternalServerError, "failed to create model")
	}

	return c.Render(http.StatusOK, "providerTable.html", providers)
}
