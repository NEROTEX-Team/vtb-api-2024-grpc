from dataclasses import dataclass, field
from os import environ


@dataclass(frozen=True, kw_only=True, slots=True)
class AppConfig:
    title: str = field(default_factory=lambda: environ.get("APP_TITLE", "Client"))
    description: str = field(
        default_factory=lambda: environ.get(
            "APP_DESCRIPTION", "Web Service - GRPC client for VTB API 2024"
        )
    )
    version: str = field(default_factory=lambda: environ.get("APP_VERSION", "1.0.0"))
    debug: bool = field(
        default_factory=lambda: environ.get("APP_DEBUG", "False").lower() == "true"
    )
