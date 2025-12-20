
# Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services.
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa: F403,F401

router = APIRouter(tags=["utilities:part-of-speech"])


@router.delete(
    "/parts-of-speech/{id}",
    
    status_code=200,
)
async def delete_parts_of_speech__id_(PartOfSpeechId: Any = None,) -> None:
    """
    Delete a specific part of speech
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_part_of_speech.delete_parts_of_speech__id_(...)
    raise NotImplementedError("DELETE /parts-of-speech/{id} not implemented yet")

@router.get(
    "/parts-of-speech",
    
    status_code=200,
)
async def get_parts_of_speech() -> None:
    """
    List all parts of speech
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_part_of_speech.get_parts_of_speech(...)
    raise NotImplementedError("GET /parts-of-speech not implemented yet")

@router.get(
    "/parts-of-speech/{id}",
    response_model=PartOfSpeech,
    status_code=200,
)
async def get_parts_of_speech__id_(PartOfSpeechId: Any = None,) -> PartOfSpeech:
    """
    Retrieve a specific part of speech
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_part_of_speech.get_parts_of_speech__id_(...)
    raise NotImplementedError("GET /parts-of-speech/{id} not implemented yet")

@router.patch(
    "/parts-of-speech/{id}",
    response_model=PartOfSpeech,
    status_code=200,
)
async def patch_parts_of_speech__id_(PartOfSpeechId: Any = None,body: PartOfSpeechInput = Body(...),) -> PartOfSpeech:
    """
    Update a specific part of speech
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_part_of_speech.patch_parts_of_speech__id_(...)
    raise NotImplementedError("PATCH /parts-of-speech/{id} not implemented yet")

@router.post(
    "/parts-of-speech",
    response_model=PartOfSpeech,
    status_code=201,
)
async def post_parts_of_speech(body: PartOfSpeechInput = Body(...),) -> PartOfSpeech:
    """
    Create a new part of speech
    OperationId: 
    Auto-generated stub. Add auth/deps if needed.
    """
    # TODO: Implement business logic in app/services and call it here.
    # Example: return await services.utilities_part_of_speech.post_parts_of_speech(...)
    raise NotImplementedError("POST /parts-of-speech not implemented yet")
