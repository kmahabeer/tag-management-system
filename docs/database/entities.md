# Entities

The `entities` table represents digital artifacts such as image files, video files, documents, URLs, or any tagged resource.

## `entities` Table

| Column       | Type    | Description                                                                                                                            |
| ------------ | ------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `id`         | UUID    | Primary key                                                                                                                            |
| `name`       | TEXT    | The name of the entity                                                                                                                 |
| `location`   | TEXT    | Location of the entity (e.g., file path, URI, URL)                                                                                     |
| `is_primary` | BOOLEAN | Indicates if this is the primary entity (e.g., latest version of an entity) or if there may exist another version that is the primary. |
| `metadata`  | JSONB   | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`)                                               |

Each **entity** can have one or more associated **purposes**, and is linked to `tags` via the `entity_tags` junction table.

## Join Tables

### `entity_purposes`

Many-to-many join table which defines one or more **purposes** for an entity. A "purpose" is an entry within the `tags` table. Only one "purpose" should be marked as *primary* for each entity.

| Column           | Type    | Description                                                 |
| ---------------- | ------- | ----------------------------------------------------------- |
| `entity_id`      | UUID    | Foreign key to `entity`                                     |
| `purpose_tag_id` | UUID    | Foreign key to `tags` (must be a tag of the type "purpose") |
| `is_primary`     | BOOLEAN | Indicates if this is the primary purpose                    |

Example: An image file might be both a "Figure Study" and a "Film Still", but only one is primary.

### `entity_versions`

Many-to-one join table linking many versions of an entity to the primary version of the entity.

| Column                | Type | Description                                                             |
| --------------------- | ---- | ----------------------------------------------------------------------- |
| `id`                  |      | Primary key; composite of `primary_entity_id` and `alternate_entity_id` |
| `primary_entity_id`   | UUID | Foreign key to `entity`                                                 |
| `alternate_entity_id` | UUID | Foreign key to `entity`                                                 |
