from sqlalchemy.orm import Session
from typing import List, Optional
from app.db.session import get_db
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
    results = [Tag.from_orm(tag) for tag in tags]
    return PaginatedResponse(results=results, total=total)
