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

// Message represents a message in a conversation. Many messages can belong to one conversation.
type Message struct {
	ID              uint
	ConversationID  uint
	ParentID        *uint
	Role            string `gorm:"not null"`
	Content         string `gorm:"not null"`
	ContentRendered string
	LlmID           uint
	BranchPath      *string
	CreatedAt       int64 `gorm:"autoUpdateTime:milli"`
	// tokenCount      int do i really need this maz be containen metadata?
	// metdata        JSON evtl besserer name

	Conversation Conversation
	Llm          Llm
	// Parent       *Message
	// Parents      []Message
}
