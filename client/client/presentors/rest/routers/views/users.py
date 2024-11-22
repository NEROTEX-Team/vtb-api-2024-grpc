from datetime import datetime
from typing import Any
from uuid import UUID

from dishka.integrations.fastapi import DishkaRoute, FromDishka
from fastapi import APIRouter, Form, Request
from fastapi.responses import HTMLResponse, RedirectResponse
from fastapi.templating import Jinja2Templates

from client.adapters.grpc.client import GRPCClient

router = APIRouter(
    prefix="/users",
    route_class=DishkaRoute,
)


def timestamp_to_datetime(ts: Any) -> datetime:
    return datetime.fromtimestamp(ts.seconds)


@router.get("/", response_class=HTMLResponse, name="users:read_users")
async def read_users(
    request: Request,
    templates: FromDishka[Jinja2Templates],
    grpc_client: FromDishka[GRPCClient],
) -> HTMLResponse:
    user_list = await grpc_client.fetch_user_list()
    return templates.TemplateResponse(
        "index.html.j2", {"request": request, "user_list": user_list}
    )


@router.get("/create", response_class=HTMLResponse)
async def create_user_form(
    request: Request, templates: FromDishka[Jinja2Templates]
) -> HTMLResponse:
    return templates.TemplateResponse("create_user.html.j2", {"request": request})


@router.post("/create")
async def create_user(
    grpc_client: FromDishka[GRPCClient],
    email: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...),
) -> RedirectResponse:
    await grpc_client.create_user(
        email=email, first_name=first_name, last_name=last_name
    )
    return RedirectResponse(url=router.url_path_for("read_users"), status_code=303)


@router.get("/{user_id}", response_class=HTMLResponse)
async def user_detail(
    request: Request,
    user_id: UUID,
    templates: FromDishka[Jinja2Templates],
    grpc_client: FromDishka[GRPCClient],
) -> HTMLResponse:
    user = await grpc_client.fetch_user_by_id(user_id=user_id)
    return templates.TemplateResponse(
        "user_detail.html.j2",
        {
            "request": request,
            "user": user,
            "timestamp_to_datetime": timestamp_to_datetime,
        },
    )


@router.get("/update/{user_id}", response_class=HTMLResponse)
async def update_user_form(
    request: Request,
    user_id: UUID,
    templates: FromDishka[Jinja2Templates],
    grpc_client: FromDishka[GRPCClient],
) -> HTMLResponse:
    user = await grpc_client.fetch_user_by_id(user_id=user_id)
    return templates.TemplateResponse(
        "update_user.html.j2", {"request": request, "user": user}
    )


@router.post("/update/{user_id}")
async def update_user(
    grpc_client: FromDishka[GRPCClient],
    user_id: UUID,
    email: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...),
) -> RedirectResponse:
    await grpc_client.update_user(
        user_id=user_id, email=email, first_name=first_name, last_name=last_name
    )
    return RedirectResponse(
        url=router.url_path_for("user_detail", user_id=user_id), status_code=303
    )


@router.post("/delete/{user_id}")
async def delete_user(
    user_id: UUID, grpc_client: FromDishka[GRPCClient]
) -> RedirectResponse:
    await grpc_client.delete_user(user_id=user_id)
    return RedirectResponse(url=router.url_path_for("read_users"), status_code=303)
