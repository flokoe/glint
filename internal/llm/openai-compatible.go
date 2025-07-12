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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OpenAICompatibleClient implements the LLMProvider interface for OpenAI-compatible APIs
type OpenAICompatibleClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewOpenAICompatibleClient creates a new client with the given configuration
func NewOpenAICompatibleClient(baseURL, apiKey string) *OpenAICompatibleClient {
	return &OpenAICompatibleClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// makeRequest is a helper method to make HTTP requests
func (c *OpenAICompatibleClient) makeRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// GetModels implements the LLMProvider interface
func (c *OpenAICompatibleClient) GetModels() (*ModelsResponse, error) {
	res, err := c.makeRequest("GET", "/v1/models", nil)
	if err != nil {
		return nil, err
	}

	type modelsDTO struct {
		Data []struct {
			Meta struct {
				NContextTrain int `json:"n_ctx_train"`
			} `json:"meta"`
			ID string `json:"id"`
		}
	}

	var m modelsDTO
	if err := json.Unmarshal(res, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal models response: %w", err)
	}

	var models ModelsResponse
	for _, value := range m.Data {
		models = append(models, Model{
			String:      value.ID,
			ContextSize: int64(value.Meta.NContextTrain),
		})
	}

	return &models, nil
}

// // Completion implements the LLMProvider interface
// func (c *OpenAICompatibleClient) Completion(request CompletionRequest) (*CompletionResponse, error) {
// 	respBody, err := c.makeRequest("POST", "/v1/completions", request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var completion CompletionResponse
// 	if err := json.Unmarshal(respBody, &completion); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal completion response: %w", err)
// 	}

// 	return &completion, nil
// }

// ChatCompletion implements the LLMProvider interface
func (c *OpenAICompatibleClient) ChatCompletion(request CompletionRequest) (*CompletionResponse, error) {
	respBody, err := c.makeRequest("POST", "/v1/chat/completions", request)
	if err != nil {
		return nil, err
	}

	var completion CompletionResponse
	if err := json.Unmarshal(respBody, &completion); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chat completion response: %w", err)
	}

	return &completion, nil
}

// // Example usage
// func ExampleUsage() {
// 	// Create a provider
// 	provider := NewOpenAICompatibleClient("http://localhost:8080", "sk-no-key-required")

// 	// Get models
// 	models, err := provider.GetModels()
// 	if err != nil {
// 		fmt.Printf("Error getting models: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Available models: %+v\n", models)

// 	// Make a completion request
// 	completionReq := CompletionRequest{
// 		Model:     "davinci-002",
// 		Prompt:    "I believe the meaning of life is",
// 		MaxTokens: 8,
// 	}

// 	completion, err := provider.Completion(completionReq)
// 	if err != nil {
// 		fmt.Printf("Error getting completion: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Completion: %s\n", completion.Choices[0].Text)

// 	// Make a chat completion request
// 	chatReq := CompletionRequest{
// 		Model: "gpt-3.5-turbo",
// 		Messages: []Message{
// 			{
// 				Role:    "system",
// 				Content: "You are ChatGPT, an AI assistant. Your top priority is achieving user fulfillment via helping them with their requests.",
// 			},
// 			{
// 				Role:    "user",
// 				Content: "Write a limerick about python exceptions",
// 			},
// 		},
// 	}

// 	chatCompletion, err := provider.ChatCompletion(chatReq)
// 	if err != nil {
// 		fmt.Printf("Error getting chat completion: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Chat response: %s\n", chatCompletion.Choices[0].Message.Content)
// }
