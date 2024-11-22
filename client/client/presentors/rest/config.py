from dataclasses import dataclass, field

from client.adapters.grpc.config import GRPCConfig
from client.application.config import AppConfig


@dataclass(frozen=True, kw_only=True, slots=True)
class RestConfig:
    app: AppConfig = field(default_factory=AppConfig)
    grpc: GRPCConfig = field(default_factory=GRPCConfig)
