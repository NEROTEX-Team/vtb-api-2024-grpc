from pydantic import PositiveInt

from client.presentors.rest.schemas import BaseSchema


class StatusResponseSchema(BaseSchema):
    ok: bool
    status_code: PositiveInt
    message: str
