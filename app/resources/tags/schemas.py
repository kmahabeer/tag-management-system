from pydantic import BaseModel
from typing import Optional
from uuid import UUID


class TagCreate(BaseModel):
    name: str
    display_name: Optional[str] = None
    # part_of_speech_id: Optional[UUID] = None
    # metadata: Optional[dict] = None


class TagOut(TagCreate):
    id: UUID
