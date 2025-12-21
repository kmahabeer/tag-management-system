
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:ratings"])


@router.delete(
    "/entities/{id}/ratings/{rating_id}",
    
    status_code=200,
)
async def delete_entities__id__ratings__rating_id_(rating_id: UUID,EntityId: Any = None,) -> None:
    """
    Delete a specific entity rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings.delete_entities__id__ratings__rating_id_(...)
    raise NotImplementedError("DELETE /entities/{id}/ratings/{rating_id} not implemented yet")

@router.get(
    "/entities/{id}/ratings",
    
    status_code=200,
)
async def get_entities__id__ratings(EntityId: Any = None,) -> None:
    """
    List ratings applied to an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings.get_entities__id__ratings(...)
    raise NotImplementedError("GET /entities/{id}/ratings not implemented yet")

@router.get(
    "/entities/{id}/ratings/{rating_id}",
    response_model=EntityContextualRatingInput,
    status_code=200,
)
async def get_entities__id__ratings__rating_id_(rating_id: UUID,EntityId: Any = None,) -> EntityContextualRatingInput:
    """
    Retrieve a specific rating for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings.get_entities__id__ratings__rating_id_(...)
    raise NotImplementedError("GET /entities/{id}/ratings/{rating_id} not implemented yet")

@router.patch(
    "/entities/{id}/ratings",
    
    status_code=200,
)
async def patch_entities__id__ratings(EntityId: Any = None,body: EntityContextualRatingUpdate = Body(...),) -> None:
    """
    Update or apply ratings for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings.patch_entities__id__ratings(...)
    raise NotImplementedError("PATCH /entities/{id}/ratings not implemented yet")

@router.patch(
    "/entities/{id}/ratings/{rating_id}",
    response_model=EntityContextualRatingInput,
    status_code=200,
)
async def patch_entities__id__ratings__rating_id_(rating_id: UUID,EntityId: Any = None,body: Any = Body(...),) -> EntityContextualRatingInput:
    """
    Update a specific entity rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_ratings.patch_entities__id__ratings__rating_id_(...)
    raise NotImplementedError("PATCH /entities/{id}/ratings/{rating_id} not implemented yet")
