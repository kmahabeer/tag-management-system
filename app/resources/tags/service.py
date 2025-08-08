from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from collections.abc import Sequence
from uuid import UUID
from app.resources.tags.models import Tag
from app.resources.tags.schemas import TagCreate


async def create_tag(db: AsyncSession, tag_in: TagCreate) -> Tag:
    result = await db.execute(select(Tag).where(Tag.name == tag_in.name))
    existing = result.scalar_one_or_none()
    if existing:
        raise ValueError("Tag already exists")

    new_tag = Tag(**tag_in.model_dump())
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
    db: AsyncSession, tag_id: UUID, tag_in: TagCreate
) -> Tag | None:
    result = await db.execute(select(Tag).where(Tag.id == tag_id))
    tag = result.scalar_one_or_none()
    if not tag:
        return None
    for field, value in tag_in.model_dump(exclude_unset=True).items():
        setattr(tag, field, value)
    await db.commit()
    await db.refresh(tag)
    return tag
