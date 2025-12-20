from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routes._register import register_routers
from app.db.session import engine, Base
from app.models.tag_model import TagModel


def create_app() -> FastAPI:
    Base.metadata.create_all(bind=engine)

    app = FastAPI(title="Your API", version="1.0.0")
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    register_routers(app)
    return app


app = create_app()
