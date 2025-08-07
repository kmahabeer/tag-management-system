from fastapi import APIRouter
from app.resources.utilities import endpoints
from app.resources.entities import endpoints
from app.resources.tags import endpoints


api_router = APIRouter(prefix="/v1")

api_router.include_router(endpoints.router)
api_router.include_router(endpoints.router)
api_router.include_router(endpoints.router)
