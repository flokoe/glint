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

package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flokoe/clairvoyance/internal/database"
	"github.com/flokoe/clairvoyance/internal/llm"
)

type AdminHandler struct {
	queries *database.Queries
}

func NewAdminHandler(db *sql.DB) *ChatHandler {
	return &ChatHandler{
		queries: database.New(db),
	}
}

type AdminViewData struct {
	Providers []database.Providers
	LLMs      []database.Llms
}

func (h *ChatHandler) AdminView(c echo.Context) error {
	providers, err := h.queries.GetProviders(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch providers")
	}

	LLMs, err := h.queries.GetLLMs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch LLMs")
	}

	data := AdminViewData{
		providers,
		LLMs,
	}

	return c.Render(http.StatusOK, "admin.html", data)
}

type providerDTO struct {
	Type string `form:"providerType"`
	URL  string `form:"providerUrl"`
}

func (h *ChatHandler) ApiAddProvider(c echo.Context) error {
	a := new(providerDTO)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Load into separate struct for security
	provider := database.AddProviderParams{
		Type: a.Type,
		Url:  a.URL,
	}

	providerResult, err := h.queries.AddProvider(c.Request().Context(), provider)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to add providers")
	}

	// Fetch the newly created provider with its models
	providers, err := h.queries.GetProviders(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch provider")
	}

	foo, err := llm.NewProvider(provider.Type, map[string]string{
		"base_url": provider.Url,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to create LLM provider")
	}

	m, err := foo.GetModels()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch models")
	}

	for _, mdl := range *m {
		llm := database.AddLLMParams{
			String:      mdl.String,
			Name:        mdl.String,
			ProviderID:  providerResult.ID,
			ContextSize: mdl.ContextSize,
		}

		// Also save to database
		_, err := h.queries.AddLLM(c.Request().Context(), llm)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to add providers")
		}
	}

	return c.Render(http.StatusOK, "providerTable.html", providers)
}
