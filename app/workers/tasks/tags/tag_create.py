from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from sqlalchemy import and_
from collections.abc import Sequence
from uuid import UUID
from api.v1.api_v1 import schemas, models

TagCreate = schemas.TagCreate
Tag = models.Tag


async def create_tag(db: AsyncSession, tag_in: TagCreate) -> Tag:
    result = await db.execute(select(Tag).where(Tag.name == tag_in.name))
    existing_tag = result.scalar_one_or_none()
    if existing_tag:
        raise ValueError("Tag already exists")

    payload = tag_in.model_dump(exclude_unset=True)
    if "metadata" in payload:
        payload["meta"] = payload.pop("metadata")

    new_tag = Tag(**payload)
    db.add(new_tag)
    await db.commit()
    await db.refresh(new_tag)
    return new_tag
