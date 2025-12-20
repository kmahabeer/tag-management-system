
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, HTTPException, Path, Query, Body
from sqlalchemy.orm import Session
from app.schemas.models import *  # noqa: F403,F401
from app.db.session import get_db
from app.schemas.models import PaginatedResponse
from app.services import tags_core

router = APIRouter(tags=["tags:core"])


@router.delete(
    "/tags/{id}",

    status_code=200,
)
async def delete_tags__id_(TagId: Any = None,) -> None:
    """
    Delete a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_core.delete_tags__id_(...)
    raise NotImplementedError("DELETE /tags/{id} not implemented yet")


@router.get(
    "/tags",
    response_model=PaginatedResponse,
    status_code=200,
)
async def get_tags(
    limit: Optional[int] = None,
    offset: Optional[int] = None,
    db: Session = Depends(get_db),
) -> PaginatedResponse:
    """
    List all tags.
    """
    try:
        return await tags_core.get_tags_service(db=db, limit=limit, offset=offset)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get(
    "/tags/{id}",
    response_model=Tag,
    status_code=200,
)
async def get_tags__id_(TagId: Any = None,) -> Tag:
    """
    Retrieve tag by ID
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_core.get_tags__id_(...)
    raise NotImplementedError("GET /tags/{id} not implemented yet")


@router.patch(
    "/tags/{id}",
    response_model=Tag,
    status_code=200,
)
async def patch_tags__id_(TagId: Any = None, body: TagInput = Body(...),) -> Tag:
    """
    Update tag metadata
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_core.patch_tags__id_(...)
    raise NotImplementedError("PATCH /tags/{id} not implemented yet")


@router.post(
    "/tags",
    response_model=Tag,
    status_code=201,
)
async def post_tags(body: TagInput = Body(...), db: Session = Depends(get_db)) -> Tag:
    """
    Create a new tag
    OperationId:
    Auto-generated stub. Add auth/deps if needed.
    """
    try:
        return await tags_core.create_tag_service(db=db, tag_input=body)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
