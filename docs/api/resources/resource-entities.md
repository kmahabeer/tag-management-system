# Resource: Entity

Entities represent digital artifacts that are being described or classified by tags. This includes images, documents, video stills, URLs, and more.

## Core Properties

| Field      | Type    | Description                                                  |
|------------|---------|--------------------------------------------------------------|
| `id`       | UUID    | Unique identifier                                            |
| `name`     | Text    | Human-readable name for the entity                           |
| `location` | Text    | Path, URL, or pointer to file location                       |
| `is_primary` | Boolean | Whether this is the canonical version of the entity        |
| `metadata` | JSONB   | Arbitrary information (e.g., source, timestamps)             |

## Relationships

- **Entity ↔ Tags**: Each entity can have many tags.
- **Entity ↔ Purpose**: Tags can define the intended purpose of the entity (e.g., `Figure Study`, `Film Still`).
- **Entity ↔ Versions**: Alternate versions of an entity (e.g., cropped, edited, different format) can be linked together.

## Use Cases

Entities allow the Tagging Service to provide context-aware metadata:
- Render tags differently based on the purpose
- Track the lineage of modified files
- Normalize annotation payloads from external tools like Label Studio

Entities are used as the primary surface for inspection, search, and UI rendering.
