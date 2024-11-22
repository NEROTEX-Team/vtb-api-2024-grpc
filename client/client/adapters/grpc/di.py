from collections.abc import AsyncIterator

import grpc
from dishka import BaseScope, Component, Provider, Scope, provide

from client.adapters.grpc.client import GRPCClient
from client.adapters.grpc.config import GRPCConfig


class GRPCProvider(Provider):
    def __init__(
        self,
        grpc_config: GRPCConfig,
        scope: BaseScope | None = None,
        component: Component | None = None,
    ) -> None:
        self.grpc_config = grpc_config
        super().__init__(scope=scope, component=component)

    @provide(scope=Scope.APP)
    async def channel(self) -> AsyncIterator[grpc.aio.Channel]:
        async with grpc.aio.insecure_channel(self.grpc_config.address) as channel:
            yield channel

    @provide(scope=Scope.APP)
    def grpc_client(self, channel: grpc.aio.Channel) -> GRPCClient:
        return GRPCClient(channel=channel)
