from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from sqlalchemy import and_
from collections.abc import Sequence
from uuid import UUID
from app.resources.tags.models import Tag
from app.resources.tags.schemas import TagCreate, TagUpdate


async def create_tag(db: AsyncSession, tag_in: TagCreate) -> Tag:
    result = await db.execute(select(Tag).where(Tag.name == tag_in.name))
    existing = result.scalar_one_or_none()
    if existing:
        raise ValueError("Tag already exists")

    payload = tag_in.model_dump(exclude_unset=True)
    if "metadata" in payload:
        payload["meta"] = payload.pop("metadata")

    new_tag = Tag(**payload)
    db.add(new_tag)
    await db.commit()
    await db.refresh(new_tag)
    return new_tag


async def get_all_tags(db: AsyncSession) -> Sequence[Tag]:
    result = await db.execute(select(Tag))
    return result.scalars().all()


async def get_tag_by_id(db: AsyncSession, id: UUID) -> Tag | None:
    result = await db.execute(select(Tag).where(Tag.id == id))
    return result.scalar_one_or_none()


async def update_tag_by_id(
    db: AsyncSession, tag_id: UUID, tag_in: TagUpdate
) -> Tag | None:
    result = await db.execute(select(Tag).where(Tag.id == tag_id))
    tag = result.scalar_one_or_none()
    if not tag:
        return None
    tag_in_data = tag_in.model_dump(exclude_unset=True)

    # If name is being updated, then enforce name uniqueness (excluding this row)
    if "name" in tag_in_data and tag_in_data["name"] != tag.name:
        conflicting_tag = await db.execute(
            select(Tag).where(and_(Tag.name == tag_in_data["name"], Tag.id != tag_id))
        )
        if conflicting_tag.scalar_one_or_none():
            raise ValueError("Tag name already exist")

    # Map API field -> ORM attribute
    if "metadata" in tag_in_data:
        tag_in_data["meta"] = tag_in_data.pop("metadata")

    for field, value in tag_in_data.items():
        setattr(tag, field, value)
    await db.commit()
    await db.refresh(tag)
    return tag


async def delete_tag_by_id(db: AsyncSession, tag_id: UUID) -> bool:
    result = await db.execute(select(Tag).where(Tag.id == tag_id))
    tag = result.scalar_one_or_none()
    if not tag:
        return False
    await db.delete(tag)
    await db.commit()
    return True
