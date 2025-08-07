from collections.abc import AsyncGenerator
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine, async_sessionmaker
from app.core.config import settings

sqlalchemy_engine = create_async_engine(settings.DATABASE_URL, echo=True)
async_session_factory = async_sessionmaker(
    bind=sqlalchemy_engine,
    expire_on_commit=False,
)


async def get_db() -> AsyncGenerator[AsyncSession, None]:
    async with async_session_factory() as session:
        yield session
