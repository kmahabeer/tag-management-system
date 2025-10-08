
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["tags:aliases"])


@router.delete(
    "/tags/{id}/aliases",
    
    status_code=200,
)
async def delete_tags__id__aliases(TagId: Any = None,) -> None:
    """
    Delete all aliases for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.delete_tags__id__aliases(...)
    raise NotImplementedError("DELETE /tags/{id}/aliases not implemented yet")

@router.delete(
    "/tags/{id}/aliases/{alias_id}",
    
    status_code=200,
)
async def delete_tags__id__aliases__alias_id_(alias_id: UUID,TagId: Any = None,) -> None:
    """
    Delete alias for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.delete_tags__id__aliases__alias_id_(...)
    raise NotImplementedError("DELETE /tags/{id}/aliases/{alias_id} not implemented yet")

@router.get(
    "/tags/{id}/aliases",
    
    status_code=200,
)
async def get_tags__id__aliases(TagId: Any = None,) -> None:
    """
    List aliases for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.get_tags__id__aliases(...)
    raise NotImplementedError("GET /tags/{id}/aliases not implemented yet")

@router.get(
    "/tags/{id}/aliases/{alias_id}",
    response_model=TagAlias,
    status_code=200,
)
async def get_tags__id__aliases__alias_id_(alias_id: UUID,TagId: Any = None,) -> TagAlias:
    """
    Retrieve alias for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.get_tags__id__aliases__alias_id_(...)
    raise NotImplementedError("GET /tags/{id}/aliases/{alias_id} not implemented yet")

@router.patch(
    "/tags/{id}/aliases",
    
    status_code=200,
)
async def patch_tags__id__aliases(TagId: Any = None,body: AliasUpdate = Body(...),) -> None:
    """
    Bulk update aliases for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.patch_tags__id__aliases(...)
    raise NotImplementedError("PATCH /tags/{id}/aliases not implemented yet")

@router.patch(
    "/tags/{id}/aliases/{alias_id}",
    response_model=TagAlias,
    status_code=200,
)
async def patch_tags__id__aliases__alias_id_(alias_id: UUID,TagId: Any = None,body: Any = Body(...),) -> TagAlias:
    """
    Update alias for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.patch_tags__id__aliases__alias_id_(...)
    raise NotImplementedError("PATCH /tags/{id}/aliases/{alias_id} not implemented yet")

@router.post(
    "/tags/{id}/aliases",
    response_model=TagAlias,
    status_code=201,
)
async def post_tags__id__aliases(TagId: Any = None,body: Any = Body(...),) -> TagAlias:
    """
    Create alias for a tag
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.tags_aliases.post_tags__id__aliases(...)
    raise NotImplementedError("POST /tags/{id}/aliases not implemented yet")
