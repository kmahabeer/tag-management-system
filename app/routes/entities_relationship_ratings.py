
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["entities:relationship_ratings"])


@router.delete(
    "/entities/{id}/relationship_ratings/{rating_id}",
    
    status_code=200,
)
async def delete_entities__id__relationship_ratings__rating_id_(rating_id: UUID,EntityId: Any = None,) -> None:
    """
    Delete a single entity relationship rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_relationship_ratings.delete_entities__id__relationship_ratings__rating_id_(...)
    raise NotImplementedError("DELETE /entities/{id}/relationship_ratings/{rating_id} not implemented yet")

@router.get(
    "/entities/{id}/relationship_ratings",
    
    status_code=200,
)
async def get_entities__id__relationship_ratings(EntityId: Any = None,) -> None:
    """
    List relationship ratings for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_relationship_ratings.get_entities__id__relationship_ratings(...)
    raise NotImplementedError("GET /entities/{id}/relationship_ratings not implemented yet")

@router.get(
    "/entities/{id}/relationship_ratings/{rating_id}",
    response_model=EntityRelationshipRatingInput,
    status_code=200,
)
async def get_entities__id__relationship_ratings__rating_id_(rating_id: UUID,EntityId: Any = None,) -> EntityRelationshipRatingInput:
    """
    Retrieve a single entity relationship rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_relationship_ratings.get_entities__id__relationship_ratings__rating_id_(...)
    raise NotImplementedError("GET /entities/{id}/relationship_ratings/{rating_id} not implemented yet")

@router.patch(
    "/entities/{id}/relationship_ratings",
    
    status_code=200,
)
async def patch_entities__id__relationship_ratings(EntityId: Any = None,body: EntityRelationshipRatingUpdate = Body(...),) -> None:
    """
    Update or create relationship ratings for an entity
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_relationship_ratings.patch_entities__id__relationship_ratings(...)
    raise NotImplementedError("PATCH /entities/{id}/relationship_ratings not implemented yet")

@router.patch(
    "/entities/{id}/relationship_ratings/{rating_id}",
    response_model=EntityRelationshipRatingInput,
    status_code=200,
)
async def patch_entities__id__relationship_ratings__rating_id_(rating_id: UUID,EntityId: Any = None,body: Any = Body(...),) -> EntityRelationshipRatingInput:
    """
    Update a specific entity relationship rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.entities_relationship_ratings.patch_entities__id__relationship_ratings__rating_id_(...)
    raise NotImplementedError("PATCH /entities/{id}/relationship_ratings/{rating_id} not implemented yet")
