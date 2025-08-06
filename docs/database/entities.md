# Entities

The `entities` table represents digital artifacts such as image files, video files, documents, URLs, or any tagged resource.

## `entities` Table

| Column       | Type    | Description                                                                                                                            |
| ------------ | ------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `id`         | UUID    | Primary key                                                                                                                            |
| `name`       | TEXT    | The name of the entity                                                                                                                 |
| `location`   | TEXT    | Location of the entity (e.g., file path, URI, URL)                                                                                     |
| `is_primary` | BOOLEAN | Indicates if this is the primary entity (e.g., latest version of an entity) or if there may exist another version that is the primary. |
| `metadata`   | JSONB   | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`)                                               |

Each **entity** can have one or more associated **purposes**, and is linked to [`tags`](./tags.md) via the `entity_tags` junction table.

## Join Tables

### `entity_purposes`

Many-to-many join table which defines one or more **purposes** for an entity. A "purpose" is an entry within the [`tags`](./tags.md) table. Only one "purpose" should be marked as _primary_ for each entity.

| Column           | Type    | Description                                                                        |
| ---------------- | ------- | ---------------------------------------------------------------------------------- |
| `id`             | UUID    | Primary key                                                                        |
| `entity_id`      |         | Foreign key to the [`entities`](./entities.md) table                               |
| `purpose_tag_id` |         | Foreign key to the [`tags`](./tags.md) table (must be a tag of the type "purpose") |
| `is_primary`     | BOOLEAN | Indicates if this is the primary purpose                                           |

Constraint: `purpose_tag_id` must refer to a tag of type `"Purpose"`.

> [!warning] TODO
>
> - Enforce via a SQL CHECK using a function or trigger if needed.
>   Example: An image file might be both a "Figure Study" and a "Film Still", but only one is primary.

### `entity_relationship_types`

| Column | Type | Description                                                          |
| ------ | ---- | -------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                          |
| `name` | TEXT | Name for a tag relationship type (e.g. "Parent-Child", "Associated") |

### `entity_relationships`

Many-to-one join table linking many versions of an entity to the primary version of the entity. Defines versioning, grouping, or derivational relationships between entities.

| Column                 | Type | Description                                                                                 |
| ---------------------- | ---- | ------------------------------------------------------------------------------------------- |
| `id`                   | UUID | Primary key                                                                                 |
| `entity_a_id`          |      | Foreign key to [**Entity A**](./entities.md#entities)                                       |
| `entity_b_id`          |      | Foreign key to [**Entity B**](./entities.md#entities)                                       |
| `relationship_type_id` |      | Foreign key to [`entity_relationship_types`](./entities.md#entity_relationship_types) table |

Examples:

- “Film Still B” → alternate version of → “Film Still A”
- “GIF Loop” → derived from → “Full Video”
- “Sketch A” → belongs to group → “Figure Study Set”
