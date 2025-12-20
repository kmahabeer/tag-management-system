from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routes._register import register_routers


def create_app() -> FastAPI:
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
