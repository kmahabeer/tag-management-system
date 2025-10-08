
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:core"])


@router.delete(
    "/entities/{id}",
    
    status_code=200,
)
async def delete_entities__id_(EntityId: Any = None,) -> None:
    """
    Delete an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_core.delete_entities__id_(...)
    raise NotImplementedError("DELETE /entities/{id} not implemented yet")

@router.get(
    "/entities",
    
    status_code=200,
)
async def get_entities(limit: int = None,offset: int = None,) -> dict:
    """
    List all entities
    OperationId: 
    Auto-generated stub. Add auth/deps if needed. PaginatedResponse detected (results: List[Entity]).
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_core.get_entities(...)
    raise NotImplementedError("GET /entities not implemented yet")

@router.get(
    "/entities/{id}",
    response_model=Entity,
    status_code=200,
)
async def get_entities__id_(EntityId: Any = None,) -> Entity:
    """
    Retrieve entity by ID
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_core.get_entities__id_(...)
    raise NotImplementedError("GET /entities/{id} not implemented yet")

@router.patch(
    "/entities/{id}",
    response_model=Entity,
    status_code=200,
)
async def patch_entities__id_(EntityId: Any = None,body: EntityInput = Body(...),) -> Entity:
    """
    Update an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_core.patch_entities__id_(...)
    raise NotImplementedError("PATCH /entities/{id} not implemented yet")

@router.post(
    "/entities",
    response_model=Entity,
    status_code=201,
)
async def post_entities(body: EntityInput = Body(...),) -> Entity:
    """
    Create a new entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_core.post_entities(...)
    raise NotImplementedError("POST /entities not implemented yet")
