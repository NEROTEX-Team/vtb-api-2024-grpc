from client.adapters.grpc.client import GRPCClient
from client.domain.entities.user import (
    CreateUser,
    UpdateUser,
    User,
    UserId,
    UserList,
    UserListParams,
)
from client.domain.interfaces.repositories.user import IUserRepository


class UserRepository(IUserRepository):
    def __init__(self, *, grpc_client: GRPCClient) -> None:
        self.__grpc_client = grpc_client

    async def fetch_user_list(self, params: UserListParams) -> UserList:
        return await self.__grpc_client.fetch_user_list(
            limit=params.limit, offset=params.offset
        )

    async def fetch_user_by_id(self, user_id: UserId) -> User | None:
        return await self.__grpc_client.fetch_user_by_id(user_id=user_id)

    async def fetch_user_by_email(self, email: str) -> User | None:
        return await self.__grpc_client.fetch_user_by_email(email=email)

    async def create_user(self, user_data: CreateUser) -> User:
        return await self.__grpc_client.create_user(user_data=user_data)

    async def update_user(self, user_data: UpdateUser) -> User:
        return await self.__grpc_client.update_user(user_data=user_data)

    async def delete_user(self, user_id: UserId) -> None:
        await self.__grpc_client.delete_user(user_id=user_id)
