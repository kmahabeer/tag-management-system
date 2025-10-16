from fastapi import FastAPI
from app.routes._register import register_routers


def create_app() -> FastAPI:
    app = FastAPI(title="Your API", version="1.0.0")
    register_routers(app)
    return app


app = create_app()
