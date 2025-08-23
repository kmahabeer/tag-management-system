from pydantic import BaseModel, Field, ConfigDict
from typing import Optional
from uuid import UUID


class TagCreate(BaseModel):
    name: str
    display_name: Optional[str] = None
    part_of_speech_id: Optional[UUID] = None
    metadata: Optional[dict] = None


class TagOut(TagCreate):
    # Enable constructing from ORM objects:
    model_config = ConfigDict(from_attributes=True)

    id: UUID
    name: str
    display_name: Optional[str] = None
    part_of_speech_id: Optional[UUID] = None

    # Pull the "meta" attribute from the ORM model but serialize it as "metadata"
    metadata: Optional[dict] = Field(
        default=None,
        validation_alias="meta",
        serialization_alias="metadata",
    )


class TagUpdate(BaseModel):
    name: Optional[str] = None
    display_name: Optional[str] = None
    part_of_speech_id: Optional[UUID] = None
    metadata: Optional[dict] = None
