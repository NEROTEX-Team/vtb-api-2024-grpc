import logging
from collections.abc import AsyncIterator

import grpc
from dishka import BaseScope, Component, Provider, Scope, provide

from client.adapters.grpc.client import GRPCClient
from client.adapters.grpc.config import GRPCConfig

log = logging.getLogger(__name__)


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
    def tls_credentials(self) -> grpc.ChannelCredentials | None:
        if not self.grpc_config.use_tls:
            log.warning("TLS is not enabled")
            return None

        with open(self.grpc_config.trusted_cert_path, "rb") as f:
            trusted = f.read()

        with open(self.grpc_config.client_cert_path, "rb") as f:
            client_cert = f.read()

        with open(self.grpc_config.client_key_path, "rb") as f:
            client_key = f.read()

        return grpc.ssl_channel_credentials(
            root_certificates=trusted,
            private_key=client_key,
            certificate_chain=client_cert,
        )

    @provide(scope=Scope.APP)
    async def channel(
        self, tls_credentials: grpc.ChannelCredentials | None
    ) -> AsyncIterator[grpc.aio.Channel]:
        if tls_credentials is None:
            async with grpc.aio.insecure_channel(self.grpc_config.address) as channel:
                yield channel
        else:
            async with grpc.aio.secure_channel(
                self.grpc_config.address, tls_credentials
            ) as channel:
                yield channel

    @provide(scope=Scope.APP)
    def grpc_client(self, channel: grpc.aio.Channel) -> GRPCClient:
        return GRPCClient(channel=channel)
