from fastapi import APIRouter

from client.presentors.rest.routers.views.users import router as user_templates

router = APIRouter(include_in_schema=False)
router.include_router(user_templates)
