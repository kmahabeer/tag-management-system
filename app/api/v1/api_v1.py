from fastapi import APIRouter
from app.api.v1.resources.tags import endpoints as tags_endpoints
from app.api.v1.resources.entities import endpoints as entities_endpoints
from app.api.v1.resources.utilities import endpoints as utilities_endpoints


api_router = APIRouter(prefix="/v1")

api_router.include_router(tags_endpoints.router)
api_router.include_router(entities_endpoints.router)
api_router.include_router(utilities_endpoints.router)
