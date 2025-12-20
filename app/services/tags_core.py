from sqlalchemy.orm import Session
from typing import Optional
from app.models.tag_model import TagModel
from app.schemas.models import Tag, TagInput, PaginatedResponse


async def get_tags_service(
    db: Session,
    limit: Optional[int] = None,
    offset: Optional[int] = None,
) -> PaginatedResponse:
    """
    Retrieve a list of all tags with optional pagination.
    """
    query = db.query(TagModel)
    total = query.count()

    if offset:
        query = query.offset(offset)
    if limit:
        query = query.limit(limit)

    tags = query.all()

    # Manual Pydantic mapping to handle reserved 'metadata' column
    results = [
        Tag(
            id=tag.id,
            name=tag.name,
            display_name=tag.display_name,
            metadata=tag.metadata_json,  # map ORM attr -> schema field
            part_of_speech_id=tag.part_of_speech_id,
            created_at=tag.created_at,
            updated_at=tag.updated_at,
        )
        for tag in tags
    ]

    return PaginatedResponse(results=results, total=total)


async def create_tag_service(db: Session, tag_input: TagInput) -> Tag:
    """
    Create a new tag.
    """
    from uuid import uuid4
    from datetime import datetime

    new_tag = TagModel(
        id=uuid4(),
        name=tag_input.name,
        display_name=tag_input.display_name or tag_input.name,
        metadata_json=tag_input.metadata or {},
        part_of_speech_id=tag_input.part_of_speech_id,
        created_at=datetime.utcnow(),
        updated_at=datetime.utcnow(),
    )
    db.add(new_tag)
    db.commit()
    db.refresh(new_tag)

    return Tag(
        id=new_tag.id,
        name=new_tag.name,
        display_name=new_tag.display_name,
        metadata=new_tag.metadata_json,
        part_of_speech_id=new_tag.part_of_speech_id,
        created_at=new_tag.created_at,
        updated_at=new_tag.updated_at,
    )
