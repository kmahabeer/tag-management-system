# Resource: Tag

Tags are the core classification mechanism in the Tagging Service. A tag may represent a person, role, genre, object, emotion, or any other label applicable to an entity or asset.

Tags are **unified** — there is no distinction between a tag for a “Person” and one for a “Genre” at the schema level. Instead, their function is defined by type and relationships.

## Core Properties

| Field      | Type    | Description                                |
|----------- |---------|--------------------------------------------|
| `id`       | UUID    | Unique identifier for the tag              |
| `name`     | Text    | Canonical name (e.g., `Quentin Tarantino`) |
| `metadata` | JSONB   | Optional additional attributes             |
| `created_at`, `updated_at` | Timestamps | Lifecycle audit fields   |

## Relationships

- **Tag ↔ TagType**: A tag can belong to multiple types (e.g., `Person`, `Role`).
- **Tag ↔ Tag**: Tags can be linked through relationships like `parent`, `related`, or `role-of`.
- **Tag ↔ Composite**: A tag may be a component of a composite tag (e.g., “very big red car”).
- **Tag ↔ Entity**: Tags are attached to digital artifacts such as images or videos.

## Special Cases

- Tags may have **aliases** to support flexible matching.
- Tags may participate in **compositions** and form multi-word expressions.
- Tags may be referenced as **purposes** to influence UI rendering.

## See also

See the [Entities resource](./resource-entities.md) to understand how tags are attached to digital artifacts.
