from fastapi import APIRouter
from app.api.v1.resources.tags import endpoints as tags_endpoints
from app.api.v1.resources.tags import schemas as tag_schemas
from app.api.v1.resources.tags import models as tag_models
from app.api.v1.resources.entities import endpoints as entities_endpoints
from app.api.v1.resources.utilities import endpoints as utilities_endpoints


api_router = APIRouter(prefix="/v1")

api_router.include_router(tags_endpoints.router)
api_router.include_router(entities_endpoints.router)
api_router.include_router(utilities_endpoints.router)


class schemas:
    """
    Aggregated imports for version 1 schemas of the API.

    Example:
        from api_v1 import API_V1

        new_tag = API_V1.schemas.TagCreate(name="example")
        tag_model = API_V1.models.Tag(...)
    """

    TagCreate = tag_schemas.TagCreate
    TagOut = tag_schemas.TagOut
    TagUpdate = tag_schemas.TagUpdate


class models:
    Tag = tag_models.Tag
