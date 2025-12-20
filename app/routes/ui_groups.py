
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["ui:groups"])


@router.delete(
    "/ui/groups/{id}",
    
    status_code=200,
)
async def delete_ui_groups__id_(UiGroupId: Any = None,) -> None:
    """
    Delete a specific UI group
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_groups.delete_ui_groups__id_(...)
    raise NotImplementedError("DELETE /ui/groups/{id} not implemented yet")

@router.get(
    "/ui/groups",
    
    status_code=200,
)
async def get_ui_groups() -> None:
    """
    List all UI groups
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_groups.get_ui_groups(...)
    raise NotImplementedError("GET /ui/groups not implemented yet")

@router.get(
    "/ui/groups/{id}",
    response_model=UiGroup,
    status_code=200,
)
async def get_ui_groups__id_(UiGroupId: Any = None,) -> UiGroup:
    """
    Retrieve a specific UI group
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_groups.get_ui_groups__id_(...)
    raise NotImplementedError("GET /ui/groups/{id} not implemented yet")

@router.patch(
    "/ui/groups/{id}",
    response_model=UiGroup,
    status_code=200,
)
async def patch_ui_groups__id_(UiGroupId: Any = None,body: UiGroupInput = Body(...),) -> UiGroup:
    """
    Update a specific UI group
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_groups.patch_ui_groups__id_(...)
    raise NotImplementedError("PATCH /ui/groups/{id} not implemented yet")

@router.post(
    "/ui/groups",
    response_model=UiGroup,
    status_code=201,
)
async def post_ui_groups(body: UiGroupInput = Body(...),) -> UiGroup:
    """
    Create a new UI group
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_groups.post_ui_groups(...)
    raise NotImplementedError("POST /ui/groups not implemented yet")
