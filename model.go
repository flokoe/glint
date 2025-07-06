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

// Adapter represents an AI API Provider/Endpoint.
type Adapter struct {
	ID        uint
	Type      string `gorm:"not null"` // e.g. "openai-compatible", "llama-stack"
	URL       string `gorm:"not null"` // Base URL for the adapter, e.g. "https://api.openai.com/v1"
	CreatedAt int64

	Models []Model
}

// Model represents a Large Language Model.
type Model struct {
	ID           uint
	String       string `gorm:"not null"` // Model identifier within the adapter
	Name         string `gorm:"not null"`
	AdapterID    uint   `gorm:"not null"`
	ContextSize  int    `gorm:"not null"`
	Capabilities string `gorm:"not null"` // e.g. "chat", "completion", "embedding"
	IsEnabled    bool   `gorm:"not null;default:true"`
	CreatedAt    int64
	UpdatedAt    int64

	Adapter Adapter
}

// User represents a user in the system.
type User struct {
	ID        uint   // Standard field for the primary key
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	CreatedAt int64  // Automatically managed by GORM for creation time

	Conversations []Conversation
}

// Conversation represents a chat conversation between a user and the AI.
type Conversation struct {
	ID        uint
	UserID    uint
	UUID      string `gorm:"unique;not null"`
	Title     *string
	CreatedAt int64
	UpdatedAt int64
	// last_message_id or branch path of past message
	// metadata JSON

	User     User
	Messages []Message
}

// Message represents a message in a conversation. Many messages can belong to one conversation.
type Message struct {
	ID              uint
	ConversationID  uint
	ParentID        *uint
	Role            string `gorm:"not null"`
	Content         string `gorm:"not null"`
	ContentRendered string
	ModelID         uint
	BranchPath      *string
	CreatedAt       int64 `gorm:"autoUpdateTime:milli"`
	// tokenCount      int do i really need this maz be containen metadata?
	// metdata        JSON evtl besserer name

	Conversation Conversation
	Model        Model
	// Parent       *Message
	// Parents      []Message
}
