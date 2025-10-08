
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["utilities:ratings:system"])


@router.delete(
    "/rating-types/{id}",
    
    status_code=200,
)
async def delete_rating_types__id_(RatingTypeId: Any = None,) -> None:
    """
    Delete a specific rating type
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.delete_rating_types__id_(...)
    raise NotImplementedError("DELETE /rating-types/{id} not implemented yet")

@router.delete(
    "/ratings/{id}",
    
    status_code=200,
)
async def delete_ratings__id_(RatingId: Any = None,) -> None:
    """
    Delete a specific rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.delete_ratings__id_(...)
    raise NotImplementedError("DELETE /ratings/{id} not implemented yet")

@router.get(
    "/rating-types",
    
    status_code=200,
)
async def get_rating_types() -> None:
    """
    List all rating types
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.get_rating_types(...)
    raise NotImplementedError("GET /rating-types not implemented yet")

@router.get(
    "/rating-types/{id}",
    response_model=RatingType,
    status_code=200,
)
async def get_rating_types__id_(RatingTypeId: Any = None,) -> RatingType:
    """
    Retrieve a specific rating type
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.get_rating_types__id_(...)
    raise NotImplementedError("GET /rating-types/{id} not implemented yet")

@router.get(
    "/ratings",
    
    status_code=200,
)
async def get_ratings() -> None:
    """
    List all ratings
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.get_ratings(...)
    raise NotImplementedError("GET /ratings not implemented yet")

@router.get(
    "/ratings/{id}",
    response_model=Rating,
    status_code=200,
)
async def get_ratings__id_(RatingId: Any = None,) -> Rating:
    """
    Retrieve a specific rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.get_ratings__id_(...)
    raise NotImplementedError("GET /ratings/{id} not implemented yet")

@router.patch(
    "/rating-types/{id}",
    response_model=RatingType,
    status_code=200,
)
async def patch_rating_types__id_(RatingTypeId: Any = None,body: RatingTypeInput = Body(...),) -> RatingType:
    """
    Update a specific rating type
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.patch_rating_types__id_(...)
    raise NotImplementedError("PATCH /rating-types/{id} not implemented yet")

@router.patch(
    "/ratings/{id}",
    response_model=Rating,
    status_code=200,
)
async def patch_ratings__id_(RatingId: Any = None,body: RatingInput = Body(...),) -> Rating:
    """
    Update a specific rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.patch_ratings__id_(...)
    raise NotImplementedError("PATCH /ratings/{id} not implemented yet")

@router.post(
    "/rating-types",
    response_model=RatingType,
    status_code=201,
)
async def post_rating_types(body: RatingTypeInput = Body(...),) -> RatingType:
    """
    Create a new rating type
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.post_rating_types(...)
    raise NotImplementedError("POST /rating-types not implemented yet")

@router.post(
    "/ratings",
    response_model=Rating,
    status_code=201,
)
async def post_ratings(body: RatingInput = Body(...),) -> Rating:
    """
    Create a new rating
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_ratings_system.post_ratings(...)
    raise NotImplementedError("POST /ratings not implemented yet")
