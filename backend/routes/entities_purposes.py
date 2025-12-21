
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:purposes"])


@router.delete(
    "/entities/{id}/purposes/{purpose_id}",
    
    status_code=200,
)
async def delete_entities__id__purposes__purpose_id_(purpose_id: UUID,EntityId: Any = None,) -> None:
    """
    Delete a specific purpose tag from an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_purposes.delete_entities__id__purposes__purpose_id_(...)
    raise NotImplementedError("DELETE /entities/{id}/purposes/{purpose_id} not implemented yet")

@router.get(
    "/entities/{id}/purposes",
    
    status_code=200,
)
async def get_entities__id__purposes(EntityId: Any = None,) -> None:
    """
    List purpose tags assigned to an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_purposes.get_entities__id__purposes(...)
    raise NotImplementedError("GET /entities/{id}/purposes not implemented yet")

@router.get(
    "/entities/{id}/purposes/{purpose_id}",
    response_model=EntityPurposeInput,
    status_code=200,
)
async def get_entities__id__purposes__purpose_id_(purpose_id: UUID,EntityId: Any = None,) -> EntityPurposeInput:
    """
    Retrieve a specific purpose tag for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_purposes.get_entities__id__purposes__purpose_id_(...)
    raise NotImplementedError("GET /entities/{id}/purposes/{purpose_id} not implemented yet")

@router.patch(
    "/entities/{id}/purposes",
    
    status_code=200,
)
async def patch_entities__id__purposes(EntityId: Any = None,body: EntityPurposeUpdate = Body(...),) -> None:
    """
    Update purpose tags for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_purposes.patch_entities__id__purposes(...)
    raise NotImplementedError("PATCH /entities/{id}/purposes not implemented yet")

@router.patch(
    "/entities/{id}/purposes/{purpose_id}",
    response_model=EntityPurposeInput,
    status_code=200,
)
async def patch_entities__id__purposes__purpose_id_(purpose_id: UUID,EntityId: Any = None,body: Any = Body(...),) -> EntityPurposeInput:
    """
    Update a specific purpose tag for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_purposes.patch_entities__id__purposes__purpose_id_(...)
    raise NotImplementedError("PATCH /entities/{id}/purposes/{purpose_id} not implemented yet")
