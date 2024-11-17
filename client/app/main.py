from fastapi import FastAPI
from app.routers import users

app = FastAPI(title="User Management System")

app.include_router(users.router, prefix="/users", tags=["users"])
