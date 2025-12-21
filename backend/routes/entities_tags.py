
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:tags"])


@router.delete(
    "/entities/{id}/tags/{tag_id}",
    
    status_code=200,
)
async def delete_entities__id__tags__tag_id_(tag_id: UUID,EntityId: Any = None,) -> None:
    """
    Remove tag assignment from entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_tags.delete_entities__id__tags__tag_id_(...)
    raise NotImplementedError("DELETE /entities/{id}/tags/{tag_id} not implemented yet")

@router.get(
    "/entities/{id}/tags",
    
    status_code=200,
)
async def get_entities__id__tags(EntityId: Any = None,) -> None:
    """
    List tags assigned to an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_tags.get_entities__id__tags(...)
    raise NotImplementedError("GET /entities/{id}/tags not implemented yet")

@router.get(
    "/entities/{id}/tags/{tag_id}",
    response_model=EntityTagAssignment,
    status_code=200,
)
async def get_entities__id__tags__tag_id_(tag_id: UUID,EntityId: Any = None,) -> EntityTagAssignment:
    """
    Retrieve specific tag assignment
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_tags.get_entities__id__tags__tag_id_(...)
    raise NotImplementedError("GET /entities/{id}/tags/{tag_id} not implemented yet")

@router.patch(
    "/entities/{id}/tags",
    
    status_code=200,
)
async def patch_entities__id__tags(EntityId: Any = None,body: EntityTagUpdate = Body(...),) -> None:
    """
    Update tags assigned to an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_tags.patch_entities__id__tags(...)
    raise NotImplementedError("PATCH /entities/{id}/tags not implemented yet")
