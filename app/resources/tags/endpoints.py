from fastapi import APIRouter, Depends, HTTPException
from uuid import UUID
from sqlalchemy.ext.asyncio import AsyncSession
from app.core.database import get_db
from app.resources.tags.schemas import TagCreate, TagOut
from app.resources.tags.service import create_tag, get_all_tags, get_tag_by_id

router = APIRouter(prefix="/tags", tags=["tags"])


@router.post("", response_model=TagOut, status_code=201)
async def create_tag_endpoint(tag_in: TagCreate, db: AsyncSession = Depends(get_db)):
    return await create_tag(db=db, tag_in=tag_in)


@router.get("", response_model=list[TagOut])
async def get_tags_endpoint(db: AsyncSession = Depends(get_db)):
    return await get_all_tags(db)


@router.get("/{id}", response_model=TagOut)
async def get_tag_by_id_endpoint(id: UUID, db: AsyncSession = Depends(get_db)):
    tag = await get_tag_by_id(db, id)
    if not tag:
        raise HTTPException(status_code=404, detail="Tag not found")
    return tag
