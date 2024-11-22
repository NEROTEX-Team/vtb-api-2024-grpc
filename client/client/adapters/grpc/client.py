import logging
from uuid import UUID

import grpc
from grpc.aio import Channel

from client.adapters.grpc.exception import GRPCClientException
from client.adapters.grpc.generated import user_pb2, user_pb2_grpc
from client.application.exceptions import EntityNotFoundException
from client.domain.entities.user import User, UserList

log = logging.getLogger(__name__)


class GRPCClient:
    __user_stub: user_pb2_grpc.UserV1Stub  # type: ignore

    def __init__(self, channel: Channel) -> None:
        self.__user_stub = user_pb2_grpc.UserV1Stub(channel)  # type: ignore

    async def fetch_user_list(self, *, limit: int = 100, offset: int = 0) -> UserList:
        try:
            user_list_data = await self.__user_stub.FetchListUsers(
                user_pb2.FetchListUsersRequest(limit=limit, offset=offset),  # type: ignore
            )
        except grpc.aio.AioRpcError as e:
            log.error("Failed to fetch user list: %s", e)
            raise GRPCClientException("Failed to fetch user list") from e
        return user_list_data

    async def fetch_user_by_id(self, *, user_id: UUID) -> User | None:
        try:
            user_data = await self.__user_stub.FetchUserById(
                user_pb2.FetchUserByIdRequest(id=str(user_id)),  # type: ignore
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.NOT_FOUND:
                return None
            log.error("Failed to fetch user by id: %s", e)
            raise GRPCClientException(f"Failed to fetch user by id {user_id}") from e
        return user_data

    async def create_user(self, *, email: str, first_name: str, last_name: str) -> User:
        try:
            user_data = await self.__user_stub.CreateUser(
                user_pb2.CreateUserRequest(  # type: ignore
                    email=email, first_name=first_name, last_name=last_name
                ),
            )
        except grpc.aio.AioRpcError as e:
            log.error("Failed to create user: %s", e)
            raise GRPCClientException(f"Failed to create user {email}") from e
        return user_data

    async def update_user(
        self, *, user_id: UUID, email: str, first_name: str, last_name: str
    ) -> User:
        try:
            user_data = await self.__user_stub.UpdateUser(
                user_pb2.UpdateUserRequest(  # type: ignore
                    id=str(user_id),
                    email=email,
                    first_name=first_name,
                    last_name=last_name,
                ),
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.NOT_FOUND:
                raise EntityNotFoundException(entity=User, entity_id=user_id)
            log.error("Failed to update user: %s", e)
            raise GRPCClientException(f"Failed to update user {user_id}") from e
        return user_data

    async def delete_user(self, *, user_id: UUID) -> None:
        try:
            await self.__user_stub.DeleteUser(
                user_pb2.DeleteUserByIdRequest(id=str(user_id)),  # type: ignore
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.NOT_FOUND:
                raise EntityNotFoundException(entity=User, entity_id=user_id)
            log.error("Failed to delete user: %s", e)
            raise GRPCClientException(f"Failed to delete user {user_id}") from e

    async def fetch_user_by_email(self, *, email: str) -> User | None:
        try:
            user_data = await self.__user_stub.FetchUserByEmail(
                user_pb2.FetchUserByEmailRequest(email=email),  # type: ignore
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.NOT_FOUND:
                return None
            log.error("Failed to fetch user by email: %s", e)
            raise GRPCClientException(f"Failed to fetch user by email {email}") from e
        return user_data
