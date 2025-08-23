from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.core.database import sqlalchemy_engine
from app.api.v1.api_v1 import api_router
from app.core.config import settings
from app.api.v1.resources.shared.models.base import Base
from app.api.v1.resources.tags import models as _
from app.api.v1.resources.utilities import models as _


@asynccontextmanager
async def lifespan(app: FastAPI):
    if settings.DATABASE_URL.startswith("sqlite"):
        print("Initializing SQLite DB...")
        async with sqlalchemy_engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)
    yield


app = FastAPI(title=settings.PROJECT_NAME, root_path="/api", lifespan=lifespan)

app.include_router(api_router)


@app.get("/")
def read_root():
    return {"message": "FastAPI is running."}
