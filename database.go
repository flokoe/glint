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

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/flokoe/clairvoyance/internal/model"
)

// MigrateAndSeed performs database migration and initial seeding.
func MigrateAndSeed(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Provider{}, &model.Llm{}, &model.User{}, &model.Conversation{}, &model.Message{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if db.Migrator().HasTable(&model.User{}) {
		var u model.User
		if err := db.First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&model.User{Email: "zaphod@pangalactic.gov", Name: "Zaphod Beeblebrox"})
		}
	}
}
