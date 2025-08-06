# Entity Tagging

## `entity_tags` Join Table

Many-to-many join table which links [`tags`](./tags.md) to [`entities`](./entities.md).

| Column      | Type  | Description                                                                                                                          |
| ----------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `id`        | UUID  | Primary key                                                                                                                          |
| `entity_id` |       | Foreign key to the [`entities`](./entities.md) table                                                                                 |
| `tag_id`    |       | Foreign key to the [`tags`](./tags.md) table                                                                                         |
| `context`   |       | Foreign key to the [`contexts`](./utilities.md#contexts) table                                                                       |
| `metadata`  | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio", "confidence": 0.92, "annotator": "user123" }`) |

`Metadata` in this context informs about how the tag is used on an entity.

## Tagging Contexts

![](./utilities.md#contexts)
