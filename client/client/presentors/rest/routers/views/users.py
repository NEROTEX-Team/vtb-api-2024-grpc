import logging
from http import HTTPStatus
from uuid import UUID

from dishka.integrations.fastapi import DishkaRoute, FromDishka
from fastapi import APIRouter, Form, Query, Request, Response
from fastapi.responses import HTMLResponse, RedirectResponse
from fastapi.templating import Jinja2Templates

from client.adapters.grpc.client import GRPCClient
from client.application.exceptions import EntityAlreadyExistsException
from client.domain.entities.user import CreateUser, UpdateUser, UserId, UserListParams
from client.domain.services.user import UserService
from client.presentors.rest.routers.views.schemas.users import UserListParamsSchema

log = logging.getLogger(__name__)

router = APIRouter(
    prefix="/users",
    route_class=DishkaRoute,
)


@router.get("/", response_class=HTMLResponse, name="users:list")
async def user_list(
    request: Request,
    templates: FromDishka[Jinja2Templates],
    user_service: FromDishka[UserService],
    params: UserListParamsSchema = Query(...),
) -> HTMLResponse:
    user_list = await user_service.fetch_user_list(
        params=UserListParams(limit=params.limit, offset=params.offset)
    )
    return templates.TemplateResponse(
        "index.html.j2", {"request": request, "user_list": user_list}
    )


@router.get("/create", response_class=HTMLResponse, name="users:create_form")
async def create_user_form(
    request: Request, templates: FromDishka[Jinja2Templates]
) -> HTMLResponse:
    return templates.TemplateResponse("create_user.html.j2", {"request": request})


@router.post("/create", name="users:create")
async def create_user(
    grpc_client: FromDishka[GRPCClient],
    user_service: FromDishka[UserService],
    request: Request,
    templates: FromDishka[Jinja2Templates],
    email: str = Form(...),
    password: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...),
) -> Response:
    try:
        await user_service.create_user(
            user_data=CreateUser(
                email=email,
                password=password,
                first_name=first_name,
                last_name=last_name,
            )
        )
    except EntityAlreadyExistsException:
        log.warning("User with email %s already exists", email)
        return templates.TemplateResponse(
            "create_user.html.j2",
            {"request": request, "errors": [f"User with email {email} already exists"]},
        )
    return RedirectResponse(
        url=router.url_path_for("users:list"),
        status_code=HTTPStatus.FOUND,
    )


@router.get("/{user_id}", response_class=HTMLResponse, name="users:detail")
async def user_detail(
    request: Request,
    user_id: UUID,
    templates: FromDishka[Jinja2Templates],
    user_service: FromDishka[UserService],
) -> HTMLResponse:
    user = await user_service.fetch_user_by_id(user_id=UserId(user_id))
    return templates.TemplateResponse(
        "user_detail.html.j2",
        {
            "request": request,
            "user": user,
        },
    )


@router.get("/update/{user_id}", response_class=HTMLResponse, name="users:update_form")
async def update_user_form(
    request: Request,
    user_id: UUID,
    templates: FromDishka[Jinja2Templates],
    user_service: FromDishka[UserService],
) -> HTMLResponse:
    user = await user_service.fetch_user_by_id(user_id=UserId(user_id))
    return templates.TemplateResponse(
        "update_user.html.j2", {"request": request, "user": user}
    )


@router.post("/update/{user_id}", name="users:update")
async def update_user(
    user_service: FromDishka[UserService],
    user_id: UUID,
    email: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...),
) -> RedirectResponse:
    await user_service.update_user(
        user_data=UpdateUser(
            id=UserId(user_id),
            email=email,
            first_name=first_name,
            last_name=last_name,
        )
    )
    return RedirectResponse(
        url=router.url_path_for("users:detail", user_id=user_id),
        status_code=HTTPStatus.FOUND,
    )


@router.post("/delete/{user_id}", name="users:delete")
async def delete_user(
    user_id: UUID,
    user_service: FromDishka[UserService],
) -> RedirectResponse:
    await user_service.delete_user(user_id=UserId(user_id))
    return RedirectResponse(
        url=router.url_path_for("users:list"),
        status_code=HTTPStatus.FOUND,
    )
