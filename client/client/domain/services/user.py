from client.application.exceptions import EntityNotFoundException
from client.domain.entities.user import (
    CreateUser,
    UpdateUser,
    User,
    UserId,
    UserList,
    UserListParams,
)
from client.domain.interfaces.repositories.user import IUserRepository


class UserService:
    __user_repository: IUserRepository

    def __init__(self, *, user_repository: IUserRepository) -> None:
        self.__user_repository = user_repository

    async def fetch_user_list(self, params: UserListParams) -> UserList:
        return await self.__user_repository.fetch_user_list(params=params)

    async def fetch_user_by_id(self, user_id: UserId) -> User:
        user = await self.__user_repository.fetch_user_by_id(user_id=user_id)
        if user is None:
            raise EntityNotFoundException(entity=User, entity_id=user_id)
        return user

    async def create_user(self, user_data: CreateUser) -> User:
        return await self.__user_repository.create_user(user_data=user_data)

    async def update_user(self, user_data: UpdateUser) -> User:
        return await self.__user_repository.update_user(user_data=user_data)

    async def delete_user(self, user_id: UserId) -> None:
        await self.__user_repository.delete_user_by_id(user_id=user_id)
