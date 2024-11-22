from collections.abc import Sequence
from dataclasses import dataclass
from datetime import datetime
from typing import NewType
from uuid import UUID

UserId = NewType("UserId", UUID)


@dataclass(frozen=True, slots=True, kw_only=True)
class User:
    id: UUID
    email: str
    first_name: str
    last_name: str
    created_at: datetime
    updated_at: datetime


@dataclass(frozen=True, slots=True, kw_only=True)
class UserList:
    total: int
    users: Sequence[User]


@dataclass(frozen=True, slots=True, kw_only=True)
class UserListParams:
    limit: int
    offset: int


@dataclass(frozen=True, slots=True, kw_only=True)
class CreateUser:
    email: str
    first_name: str
    last_name: str


@dataclass(frozen=True, slots=True, kw_only=True)
class UpdateUser:
    id: UserId
    email: str
    first_name: str
    last_name: str
