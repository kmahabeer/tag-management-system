
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["tags:compositions"])


@router.delete(
    "/tags/{id}/compositions",
    
    status_code=200,
)
async def delete_tags__id__compositions(TagId: Any = None,) -> None:
    """
    Delete composition for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.delete_tags__id__compositions(...)
    raise NotImplementedError("DELETE /tags/{id}/compositions not implemented yet")

@router.delete(
    "/tags/{id}/compositions/{composition_id}",
    
    status_code=200,
)
async def delete_tags__id__compositions__composition_id_(composition_id: UUID,TagId: Any = None,) -> None:
    """
    Delete composition component
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.delete_tags__id__compositions__composition_id_(...)
    raise NotImplementedError("DELETE /tags/{id}/compositions/{composition_id} not implemented yet")

@router.get(
    "/tags/{id}/compositions",
    
    status_code=200,
)
async def get_tags__id__compositions(TagId: Any = None,) -> None:
    """
    List components of a composite tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.get_tags__id__compositions(...)
    raise NotImplementedError("GET /tags/{id}/compositions not implemented yet")

@router.get(
    "/tags/{id}/compositions/{composition_id}",
    response_model=TagComponent,
    status_code=200,
)
async def get_tags__id__compositions__composition_id_(composition_id: UUID,TagId: Any = None,) -> TagComponent:
    """
    Retrieve component of a composite tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.get_tags__id__compositions__composition_id_(...)
    raise NotImplementedError("GET /tags/{id}/compositions/{composition_id} not implemented yet")

@router.patch(
    "/tags/{id}/compositions",
    
    status_code=200,
)
async def patch_tags__id__compositions(TagId: Any = None,body: CompositionUpdate = Body(...),) -> None:
    """
    Update composition for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.patch_tags__id__compositions(...)
    raise NotImplementedError("PATCH /tags/{id}/compositions not implemented yet")

@router.patch(
    "/tags/{id}/compositions/{composition_id}",
    response_model=TagComponent,
    status_code=200,
)
async def patch_tags__id__compositions__composition_id_(composition_id: UUID,TagId: Any = None,body: Any = Body(...),) -> TagComponent:
    """
    Update composition component
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.patch_tags__id__compositions__composition_id_(...)
    raise NotImplementedError("PATCH /tags/{id}/compositions/{composition_id} not implemented yet")

@router.post(
    "/tags/{id}/compositions",
    response_model=TagComponent,
    status_code=201,
)
async def post_tags__id__compositions(TagId: Any = None,body: TagComponentInput = Body(...),) -> TagComponent:
    """
    Add component to a composite tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_compositions.post_tags__id__compositions(...)
    raise NotImplementedError("POST /tags/{id}/compositions not implemented yet")
