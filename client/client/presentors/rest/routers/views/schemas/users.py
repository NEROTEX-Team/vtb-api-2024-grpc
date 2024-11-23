from pydantic import BaseModel, Field, NonNegativeInt, PositiveInt


class UserListParamsSchema(BaseModel):
    limit: PositiveInt = Field(default=100, ge=1, le=100)
    offset: NonNegativeInt = Field(default=0, ge=0)
