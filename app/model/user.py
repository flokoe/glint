"""User model and CRUD operations using SQLModel."""

from __future__ import annotations

from datetime import datetime

from sqlmodel import Field, SQLModel


class UserBase(SQLModel):
    """Base user model with shared attributes."""

    email: str = Field(unique=True, index=True)
    name: str | None = Field(default=None)


class User(UserBase, table=True):
    """User table model."""

    id: int | None = Field(default=None, primary_key=True)
    is_admin: bool = Field(default=False)
    created_at: int = Field(default_factory=lambda: int(datetime.now().timestamp()))
