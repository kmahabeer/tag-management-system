import uuid
from sqlalchemy import String, Column, Boolean, Text
from sqlalchemy.dialects.postgresql import UUID
from app.resources.shared.models.base import Base


class PartOfSpeech(Base):
    __tablename__ = "parts_of_speech"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    name = Column(String, unique=True, nullable=False)
    description = Column(Text, nullable=True)
    is_active = Column(Boolean, nullable=True)
