from dataclasses import dataclass, field
from os import environ


@dataclass(frozen=True, slots=True, kw_only=True)
class GRPCConfig:
    host: str = field(default_factory=lambda: environ.get("APP_GRPC_HOST", "localhost"))
    port: str = field(default_factory=lambda: environ.get("APP_GRPC_PORT", "50051"))

    use_tls: bool = field(
        default_factory=lambda: environ.get("APP_TLS_USE", "False").lower() == "true"
    )
    trusted_cert_path: str = field(
        default_factory=lambda: environ.get("APP_TLS_CLIENT_CA_FILE_PATH", "")
    )
    client_cert_path: str = field(
        default_factory=lambda: environ.get("APP_TLS_CLIENT_CERT_FILE_PATH", "")
    )
    client_key_path: str = field(
        default_factory=lambda: environ.get("APP_TLS_CLIENT_KEY_FILE_PATH", "")
    )

    @property
    def address(self) -> str:
        return f"{self.host}:{self.port}"
