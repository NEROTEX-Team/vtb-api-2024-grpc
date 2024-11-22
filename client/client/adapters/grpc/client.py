import logging
from datetime import UTC, datetime
from typing import Any
from uuid import UUID

import grpc
from grpc.aio import Channel

from client.adapters.grpc.exception import GRPCClientException
from client.adapters.grpc.generated import user_pb2, user_pb2_grpc
from client.application.exceptions import (
    EntityAlreadyExistsException,
    EntityNotFoundException,
)
from client.domain.entities.user import CreateUser, UpdateUser, User, UserList

log = logging.getLogger(__name__)


class GRPCClient:
    __user_stub: user_pb2_grpc.UserV1Stub  # type: ignore

    def __init__(self, channel: Channel) -> None:
        self.__user_stub = user_pb2_grpc.UserV1Stub(channel)  # type: ignore

    async def fetch_user_list(self, *, limit: int = 100, offset: int = 0) -> UserList:
        try:
            user_list_data: user_pb2.FetchUserListResponse = (  # type: ignore
                await self.__user_stub.FetchUserList(
                    user_pb2.FetchUserListRequest(limit=limit, offset=offset),  # type: ignore
                )
            )
        except grpc.aio.AioRpcError as e:
            log.error("Failed to fetch user list: %s", e)
            raise GRPCClientException("Failed to fetch user list") from e
        return UserList(
            total=user_list_data.total,
            items=[convert_grpc_user_to_user(user) for user in user_list_data.users],
        )

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
        return convert_grpc_user_to_user(user_data.user)

    async def create_user(self, *, user_data: CreateUser) -> User:
        try:
            created_user_data = await self.__user_stub.CreateUser(
                user_pb2.CreateUserRequest(  # type: ignore
                    email=user_data.email,
                    password=user_data.password,
                    first_name=user_data.first_name,
                    last_name=user_data.last_name,
                ),
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.ALREADY_EXISTS:
                raise EntityAlreadyExistsException(
                    entity=User, unique_id=user_data.email
                )
            log.error("Failed to create user: %s", e)
            raise GRPCClientException(f"Failed to create user {user_data.email}") from e
        return convert_grpc_user_to_user(created_user_data.user)

    async def update_user(self, *, user_data: UpdateUser) -> User:
        try:
            updated_user_data = await self.__user_stub.UpdateUser(
                user_pb2.UpdateUserRequest(  # type: ignore
                    id=str(user_data.id),
                    email=user_data.email,
                    first_name=user_data.first_name,
                    last_name=user_data.last_name,
                ),
            )
        except grpc.aio.AioRpcError as e:
            if e.code() == grpc.StatusCode.NOT_FOUND:
                raise EntityNotFoundException(entity=User, entity_id=user_data.id)
            log.error("Failed to update user: %s", e)
            raise GRPCClientException(f"Failed to update user {user_data.id}") from e
        return convert_grpc_user_to_user(updated_user_data.user)

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
        return convert_grpc_user_to_user(user_data.user)


def convert_grpc_user_to_user(user: user_pb2.User) -> User:  # type: ignore
    return User(
        id=UUID(user.id),
        email=user.email,
        first_name=user.first_name,
        last_name=user.last_name,
        created_at=grpc_timestamp_to_datetime(user.created_at),
        updated_at=grpc_timestamp_to_datetime(user.updated_at),
    )


def grpc_timestamp_to_datetime(timestamp: Any) -> datetime:
    return timestamp.ToDatetime().replace(tzinfo=UTC)
