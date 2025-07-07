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
	"fmt"
)

// Helper function to create providers based on configuration
func NewProvider(providerType string, config map[string]string) (LLMProvider, error) {
	switch providerType {
	case "openai-compatible":
		return NewOpenAICompatibleClient(config["base_url"], config["api_key"]), nil
	// case "custom":
	// 	return &CustomLLMProvider{
	// 		Endpoint: config["endpoint"],
	// 		Token:    config["token"],
	// 		Client:   &http.Client{Timeout: 30 * time.Second},
	// 	}, nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", providerType)
	}
}
