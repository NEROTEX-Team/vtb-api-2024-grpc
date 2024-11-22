from http import HTTPStatus

import grpc
from fastapi import HTTPException, Request
from fastapi.responses import JSONResponse

from client.application.exceptions import ClientException
from client.presentors.rest.routers.api.v1.schemas.common import StatusResponseSchema


async def http_exception_handler(request: Request, exc: HTTPException) -> JSONResponse:
    return exception_json_response(
        status_code=exc.status_code,
        message=exc.detail,
    )


async def grpc_error_handler(
    request: Request, exc: grpc.aio.AioRpcError
) -> JSONResponse:
    return exception_json_response(
        status_code=HTTPStatus.INTERNAL_SERVER_ERROR,
        message=str(exc),
    )


async def client_exception_handler(
    request: Request,
    exc: ClientException,
) -> JSONResponse:
    return exception_json_response(
        status_code=HTTPStatus.INTERNAL_SERVER_ERROR,
        message=exc.message,
    )


def exception_json_response(status_code: int, message: str) -> JSONResponse:
    return JSONResponse(
        status_code=status_code,
        content=StatusResponseSchema(
            ok=False,
            status_code=status_code,
            message=message,
        ).model_dump(mode="json"),
    )
