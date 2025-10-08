
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:ratings:system"])


@router.delete(
    "/entity-ratings/{id}",
    
    status_code=200,
)
async def delete_entity_ratings__id_(id: UUID,) -> None:
    """
    Delete a specific entity rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings_system.delete_entity_ratings__id_(...)
    raise NotImplementedError("DELETE /entity-ratings/{id} not implemented yet")

@router.get(
    "/entity-ratings",
    
    status_code=200,
)
async def get_entity_ratings() -> None:
    """
    List all entity ratings
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings_system.get_entity_ratings(...)
    raise NotImplementedError("GET /entity-ratings not implemented yet")

@router.patch(
    "/entity-ratings/{id}",
    response_model=EntityContextualRatingInput,
    status_code=200,
)
async def patch_entity_ratings__id_(id: UUID,body: EntityContextualRatingInput = Body(...),) -> EntityContextualRatingInput:
    """
    Update a specific entity rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings_system.patch_entity_ratings__id_(...)
    raise NotImplementedError("PATCH /entity-ratings/{id} not implemented yet")

@router.post(
    "/entity-ratings",
    response_model=EntityContextualRatingInput,
    status_code=201,
)
async def post_entity_ratings(body: EntityContextualRatingInput = Body(...),) -> EntityContextualRatingInput:
    """
    Create a new entity rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings_system.post_entity_ratings(...)
    raise NotImplementedError("POST /entity-ratings not implemented yet")
