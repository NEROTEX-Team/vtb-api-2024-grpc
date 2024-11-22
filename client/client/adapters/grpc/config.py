from dataclasses import dataclass, field
from os import environ


@dataclass(frozen=True, slots=True, kw_only=True)
class GRPCConfig:
    host: str = field(default_factory=lambda: environ.get("APP_GRPC_HOST", "localhost"))
    port: str = field(default_factory=lambda: environ.get("APP_GRPC_PORT", "50051"))

    @property
    def address(self) -> str:
        return f"{self.host}:{self.port}"
