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

type ChatHandler struct {
	queries *database.Queries
	infer   llm.LLMProvider
}

func NewChatHandler(db *sql.DB, provider llm.LLMProvider) *ChatHandler {
	return &ChatHandler{
		queries: database.New(db),
		infer: provider,
	}
}

type chatData struct {
	LLMs []database.Llm
}

func (h *ChatHandler) ChatView(c echo.Context) error {
	LLMs, err := h.queries.GetLLMs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch LLMs")
	}

	data := chatData{
		LLMs,
	}

	return c.Render(http.StatusOK, "chat.html", data)
}

type chatInputDTO struct {
	Content string `form:"content"`
	Llm     string `form:"llm"`
}

type responseData struct {
	UserMessage string
	Completion  string
}

func (h *ChatHandler) ApiCompletion(c echo.Context) error {
	a := new(chatInputDTO)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	result, err := h.infer.ChatCompletion(llm.CompletionRequest{
		Model:       a.Llm,
		Messages:    []llm.Message{{Role: "user", Content: a.Content}},
	})
	if err != nil {
		c.Logger().Error("Completion failed:", err)
		return c.String(http.StatusInternalServerError, "Completion failed")
	}

	data := responseData{
		UserMessage: a.Content,
		Completion:  result.Choices[0].Message.Content,
	}

	return c.Render(http.StatusOK, "firstCompletion.html", data)
}
