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
along with this program. If not, see
<https://www.gnu.org/licenses/>.

SPDX-License-Identifier: AGPL-3.0-only
*/

package model

// Llm represents a Large Language Model.
type Llm struct {
	ID           uint
	String       string `gorm:"not null"` // Model identifier within the provider
	Name         string `gorm:"not null"`
	ProviderID   uint   `gorm:"not null"`
	ContextSize  int    `gorm:"not null"`
	Capabilities string `gorm:"not null"` // e.g. "chat", "completion", "embedding"
	IsEnabled    bool   `gorm:"not null;default:true"`
	CreatedAt    int64
	UpdatedAt    int64

	Provider Provider
}
