import os
from pydantic import BaseModel


def load_env_file(path=".env"):
    if not os.path.exists(path):
        return
    with open(path) as f:
        for line in f:
            line = line.strip()
            if line and not line.startswith("#") and "=" in line:
                key, value = line.split("=", 1)
                os.environ.setdefault(key.strip(), value.strip())


load_env_file()


class Settings(BaseModel):
    PROJECT_NAME: str = "Tag Management System API"
    API_V1_STR: str
    DATABASE_URL: str


settings = Settings(
    PROJECT_NAME=os.getenv("PROJECT_NAME", "Tag Management System API"),
    API_V1_STR=os.getenv("API_V1_STR", "/api/v1"),
    DATABASE_URL=os.getenv("DATABASE_URL", "sqlite+aiosqlite:///./tagging.db"),
)
