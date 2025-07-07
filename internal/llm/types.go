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

package llm

// // Common types for LLM interactions
// type Message struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

// type CompletionRequest struct {
// 	Model       string    `json:"model"`
// 	Messages    []Message `json:"messages,omitempty"`
// 	Prompt      string    `json:"prompt,omitempty"`
// 	MaxTokens   int       `json:"max_tokens,omitempty"`
// 	Temperature float64   `json:"temperature,omitempty"`
// 	Stream      bool      `json:"stream,omitempty"`
// }

// type CompletionResponse struct {
// 	ID      string `json:"id"`
// 	Object  string `json:"object"`
// 	Created int64  `json:"created"`
// 	Model   string `json:"model"`
// 	Choices []struct {
// 		Index   int     `json:"index"`
// 		Message Message `json:"message,omitempty"`
// 		Text    string  `json:"text,omitempty"`
// 		Finish  string  `json:"finish_reason"`
// 	} `json:"choices"`
// 	Usage struct {
// 		PromptTokens     int `json:"prompt_tokens"`
// 		CompletionTokens int `json:"completion_tokens"`
// 		TotalTokens      int `json:"total_tokens"`
// 	} `json:"usage"`
// }

type Model struct {
	String      string `json:"id"`
	ContextSize int    `json:"context_size"`
}

type ModelsResponse []Model

// LLMProvider interface defines the contract for all LLM providers
type LLMProvider interface {
	// GetModels returns available models
	GetModels() (*ModelsResponse, error)

	// Completion performs a text completion
	// Completion(request CompletionRequest) (*CompletionResponse, error)

	// ChatCompletion performs a chat completion
	// ChatCompletion(request CompletionRequest) (*CompletionResponse, error)
}
