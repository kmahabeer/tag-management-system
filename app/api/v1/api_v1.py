from fastapi import APIRouter
from app.api.v1.endpoints import tags, entities, utilities


api_router = APIRouter(prefix="/v1")

api_router.include_router(tags.router)
api_router.include_router(entities.router)
api_router.include_router(utilities.router)
