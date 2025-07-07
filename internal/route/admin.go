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

	llamaProvider := llm.NewOpenAICompatibleClient("http://127.0.0.1:8080", "")

	// Use the same interface regardless of provider
	models, _ := llamaProvider.GetModels()
	if models != nil {
		for _, model := range *models {
			fmt.Printf("Model: %s, Context Size: %d\n", model.String, model.ContextSize)
		}
	}

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

	return c.Render(http.StatusOK, "providerTable.html", providers)
}
