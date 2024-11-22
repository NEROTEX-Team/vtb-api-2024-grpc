from dishka.integrations.fastapi import DishkaRoute
from fastapi import APIRouter

router = APIRouter(prefix="/users", route_class=DishkaRoute, tags=["Users"])
