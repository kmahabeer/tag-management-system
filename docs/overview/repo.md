---
title: Repository Structure
---
# Repository

```txt
tagging_service/
├── app/
│   ├── api/
│   │   ├── v1/
│   │   │   ├── endpoints/
│   │   │   │   ├── tags.py
│   │   │   │   ├── entities.py
│   │   │   │   ├── utilities.py
│   │   │   ├── api_v1.py
│   ├── core/
│   │   ├── config.py
│   │   ├── database.py
│   ├── models/
│   │   ├── __init__.py
│   │   ├── tag.py
│   │   ├── entity.py
│   │   ├── context.py
│   │   ├── rating.py
│   │   ├── pos.py
│   ├── schemas/
│   │   ├── __init__.py
│   │   ├── tag.py
│   │   ├── entity.py
│   │   ├── context.py
│   │   ├── rating.py
│   │   ├── pos.py
│   ├── services/
│   │   ├── tag_service.py
│   │   ├── entity_service.py
│   ├── main.py
├── alembic/ (optional if using migrations)
├── requirements.txt
├── README.md

```

## Breakdown of Key Modules

### `app/api/v1/endpoints/`

Contains route handlers per resource:

- `tags.py`: `/tags`, `/tags/{id}`, aliases, relationships, compositions, etc.
- `entities.py`: `/entities`, `/entities/{id}`, tagging, purpose, relationships
- `utilities.py`: `/contexts`, `/parts-of-speech`

### `app/models/`

SQLAlchemy models matching your schema:

- `tag.py`: Includes `Tag`, `TagAlias`, `TagRelationship`, `TagComposition`, etc.
- `entity.py`: `Entity`, `EntityPurpose`, `EntityTag`, `EntityRelationship`
- `context.py`: `Context`
- `rating.py`: `Rating`, `RatingType`, `TagContextRating`, `TagRelationshipRating`
- `pos.py`: `PartOfSpeech`

### `app/schemas/`

Pydantic models for request/response validation. Mirrors the structure of `models/`.

### `app/services/`

Business logic:

- `tag_service.py`: Logic for creating/updating tags, enforcing grammar rules, handling aliases/compositions.
- `entity_service.py`: Logic for tagging entities, assigning purposes, managing relationships.

## Difference Between Models and Schemas

| Aspect         | **Models** (`models/`)                       | **Schemas** (`schemas/`)                                 |
| -------------- | -------------------------------------------- | -------------------------------------------------------- |
| Purpose        | Represents the **database structure**        | Represents the **data shape** for API requests/responses |
| Implementation | SQLAlchemy (`Tag`, `Entity`, etc.)           | Pydantic (`TagCreate`, `TagOut`, etc.)                   |
| Used in        | Internal DB logic – e.g., inserts, queries   | API validation and serialization (FastAPI)               |
| Example        | `Tag.name = Column(String)`                  | `TagCreate.name: str`                                    |
| Tied to DB     | Yes – maps to actual DB tables               | No – purely for data validation                          |
| Extra Logic    | May have relationships, indexes, constraints | May include validation rules, default values             |

- Use schemas to validate data coming in or out of your API
- Use models to persist or retrieve that data in the database
