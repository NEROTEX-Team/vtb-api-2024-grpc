from typing import Any


class ClientException(Exception):
    def __init__(self, message: str) -> None:
        self.message = message
        super().__init__(self.message)


class EntityNotFoundException(ClientException):
    def __init__(self, entity: type, entity_id: Any) -> None:
        super().__init__(f"{entity.__name__} with id {entity_id} not found")


class EntityAlreadyExistsException(ClientException):
    def __init__(self, entity: type, unique_id: Any) -> None:
        super().__init__(f"{entity.__name__} with unique id {unique_id} already exists")
