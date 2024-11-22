import abc
from typing import Protocol

from client.domain.entities.user import (
    CreateUser,
    UpdateUser,
    User,
    UserId,
    UserList,
    UserListParams,
)


class IUserRepository(Protocol):
    @abc.abstractmethod
    async def fetch_user_list(self, *, params: UserListParams) -> UserList:
        raise NotImplementedError

    @abc.abstractmethod
    async def fetch_user_by_id(self, *, user_id: UserId) -> User | None:
        raise NotImplementedError

    @abc.abstractmethod
    async def fetch_user_by_email(self, *, email: str) -> User | None:
        raise NotImplementedError

    @abc.abstractmethod
    async def create_user(self, *, user_data: CreateUser) -> User:
        raise NotImplementedError

    @abc.abstractmethod
    async def update_user(self, *, user_data: UpdateUser) -> User:
        raise NotImplementedError

    @abc.abstractmethod
    async def delete_user_by_id(self, *, user_id: UserId) -> None:
        raise NotImplementedError
