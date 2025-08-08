from sqlalchemy import Column, String, Text, JSON, ForeignKey
from sqlalchemy.dialects.postgresql import UUID
from app.resources.shared.models.base import Base
import uuid


class Tag(Base):
    __tablename__ = "tags"
    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    name = Column(String, nullable=False, unique=True)
    display_name = Column(String, nullable=True)
    part_of_speech_id = Column(
        UUID(as_uuid=True),
        ForeignKey("parts_of_speech.id"),
        nullable=True,
    )
    meta = Column("metadata", JSON, nullable=True)
