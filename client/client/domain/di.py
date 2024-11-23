from dishka import Provider, Scope, provide

from client.domain.interfaces.repositories.user import IUserRepository
from client.domain.services.user import UserService


class DomainProvider(Provider):
    @provide(scope=Scope.REQUEST)
    def user_service(self, user_repository: IUserRepository) -> UserService:
        return UserService(user_repository=user_repository)
