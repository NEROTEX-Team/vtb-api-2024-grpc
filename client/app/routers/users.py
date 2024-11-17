import os
from fastapi import APIRouter, Request, Form, HTTPException
from fastapi.responses import HTMLResponse, RedirectResponse
from fastapi.templating import Jinja2Templates
from datetime import datetime

import grpc
from app import user_pb2, user_pb2_grpc

GRPC_SERVER_HOST = os.environ.get('GRPC_SERVER_HOST', 'localhost')
GRPC_SERVER_PORT = os.environ.get('GRPC_SERVER_PORT', '50051')

router = APIRouter()

grpc_server_address = f"{GRPC_SERVER_HOST}:{GRPC_SERVER_PORT}"
templates = Jinja2Templates(directory="app/templates")

channel = grpc.insecure_channel(grpc_server_address)
stub = user_pb2_grpc.UserV1Stub(channel)

def timestamp_to_datetime(ts):
    return datetime.fromtimestamp(ts.seconds)

def handle_grpc_error(e):
    raise HTTPException(status_code=500, detail=str(e))

@router.get("/", response_class=HTMLResponse)
async def read_users(request: Request):
    try:
        grpc_request = user_pb2.FetchAllUsersRequest(limit=100, offset=0)
        grpc_response = stub.FetchListUsers(grpc_request)
        users = grpc_response.users
        return templates.TemplateResponse("index.html", {"request": request, "users": users})
    except grpc.RpcError as e:
        handle_grpc_error(e)

@router.get("/create", response_class=HTMLResponse)
async def create_user_form(request: Request):
    return templates.TemplateResponse("create_user.html", {"request": request})

@router.post("/create")
async def create_user(
    request: Request,
    email: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...)
):
    try:
        grpc_request = user_pb2.CreateUserRequest(
            email=email,
            first_name=first_name,
            last_name=last_name
        )
        stub.CreateUser(grpc_request)
        return RedirectResponse(url=router.url_path_for('read_users'), status_code=303)
    except grpc.RpcError as e:
        handle_grpc_error(e)

@router.get("/{user_id}", response_class=HTMLResponse)
async def user_detail(request: Request, user_id: str):
    try:
        grpc_request = user_pb2.FetchUserByIdRequest(id=user_id)
        grpc_response = stub.FetchUserById(grpc_request)
        user = grpc_response.user
        return templates.TemplateResponse(
            "user_detail.html",
            {"request": request, "user": user, "timestamp_to_datetime": timestamp_to_datetime}
        )
    except grpc.RpcError as e:
        handle_grpc_error(e)

@router.get("/update/{user_id}", response_class=HTMLResponse)
async def update_user_form(request: Request, user_id: str):
    try:
        grpc_request = user_pb2.FetchUserByIdRequest(id=user_id)
        grpc_response = stub.FetchUserById(grpc_request)
        user = grpc_response.user
        return templates.TemplateResponse("update_user.html", {"request": request, "user": user})
    except grpc.RpcError as e:
        handle_grpc_error(e)

@router.post("/update/{user_id}")
async def update_user(
    request: Request,
    user_id: str,
    email: str = Form(...),
    first_name: str = Form(...),
    last_name: str = Form(...)
):
    try:
        grpc_request = user_pb2.UpdateUserRequest(
            id=user_id,
            email=email,
            first_name=first_name,
            last_name=last_name
        )
        stub.UpdateUser(grpc_request)
        return RedirectResponse(url=router.url_path_for('user_detail', user_id=user_id), status_code=303)
    except grpc.RpcError as e:
        handle_grpc_error(e)

@router.post("/delete/{user_id}")
async def delete_user(user_id: str):
    try:
        grpc_request = user_pb2.DeleteUserByIdRequest(id=user_id)
        stub.DeleteUser(grpc_request)
        return RedirectResponse(url=router.url_path_for('read_users'), status_code=303)
    except grpc.RpcError as e:
        handle_grpc_error(e)
