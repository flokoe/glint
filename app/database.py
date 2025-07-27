# Clairvoyance - A feature-rich web based AI chat UI
# Copyright (C) 2025 Florian Köhler

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

# SPDX-License-Identifier: AGPL-3.0-only

from sqlmodel import Session, SQLModel, create_engine

from config import settings
from model.user import User

sqlite_url = "sqlite:///clairvoyance.db"
engine = create_engine(sqlite_url, connect_args={"check_same_thread": False})


def create_schema_and_tables():
    SQLModel.metadata.create_all(engine)


def seed() -> None:
    # AIDEV-NOTE: Only create default user in single-user mode
    if not settings.single_user_mode:
        return

    with Session(engine) as session:
        # AIDEV-NOTE: Check if default user already exists to avoid duplicates
        existing_user = session.get(User, 1)
        if existing_user is None:
            user = User(email="zaphod@pangalactic.gov", name="Zaphod Beeblebrox", is_admin=True)
            session.add(user)
            session.commit()


def get_session():
    with Session(engine) as session:
        yield session
