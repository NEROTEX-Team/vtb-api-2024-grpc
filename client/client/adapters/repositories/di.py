from dishka import AnyOf, Provider, Scope, provide

from client.adapters.grpc.client import GRPCClient
from client.adapters.repositories.user import UserRepository
from client.domain.interfaces.repositories.user import IUserRepository


class RepositoryProvider(Provider):
    @provide(scope=Scope.REQUEST)
    def user_repository(
        self, grpc_client: GRPCClient
    ) -> AnyOf[UserRepository, IUserRepository]:
        return UserRepository(grpc_client=grpc_client)
