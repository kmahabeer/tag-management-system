
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["ui:layouts"])


@router.delete(
    "/ui/layouts/{id}",
    
    status_code=200,
)
async def delete_ui_layouts__id_(UiLayoutId: Any = None,) -> None:
    """
    Delete a specific UI layout
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_layouts.delete_ui_layouts__id_(...)
    raise NotImplementedError("DELETE /ui/layouts/{id} not implemented yet")

@router.get(
    "/ui/layouts",
    
    status_code=200,
)
async def get_ui_layouts() -> None:
    """
    List all UI layouts
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_layouts.get_ui_layouts(...)
    raise NotImplementedError("GET /ui/layouts not implemented yet")

@router.get(
    "/ui/layouts/{id}",
    response_model=UiLayout,
    status_code=200,
)
async def get_ui_layouts__id_(UiLayoutId: Any = None,) -> UiLayout:
    """
    Retrieve a specific UI layout
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_layouts.get_ui_layouts__id_(...)
    raise NotImplementedError("GET /ui/layouts/{id} not implemented yet")

@router.patch(
    "/ui/layouts/{id}",
    response_model=UiLayout,
    status_code=200,
)
async def patch_ui_layouts__id_(UiLayoutId: Any = None,body: UiLayoutInput = Body(...),) -> UiLayout:
    """
    Update a specific UI layout
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_layouts.patch_ui_layouts__id_(...)
    raise NotImplementedError("PATCH /ui/layouts/{id} not implemented yet")

@router.post(
    "/ui/layouts",
    response_model=UiLayout,
    status_code=201,
)
async def post_ui_layouts(body: UiLayoutInput = Body(...),) -> UiLayout:
    """
    Create a new UI layout
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.ui_layouts.post_ui_layouts(...)
    raise NotImplementedError("POST /ui/layouts not implemented yet")
