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

from contextlib import asynccontextmanager
from typing import Annotated

from fastapi import Depends, FastAPI
from sqlmodel import Session

import database

SessionDep = Annotated[Session, Depends(database.get_session)]

@asynccontextmanager
async def lifespan(app: FastAPI):
    database.create_schema_and_tables()
    database.seed()
    yield

app = FastAPI(lifespan=lifespan)
