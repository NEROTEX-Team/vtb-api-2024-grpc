from fastapi import APIRouter

from client.presentors.rest.routers.api.v1.endpoints.users import router as user_router

router = APIRouter(prefix="/v1")
router.include_router(user_router)
