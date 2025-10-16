from sqlalchemy.orm import Session
from typing import Optional
from app.models.tag_model import TagModel
from app.schemas.models import Tag, PaginatedResponse


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
