---
title: Entity-Tag Relationships
parent: Database Schema
has_children: true
nav_order: 3
---
# Entity Tagging

## `entity_tags` Table

Many-to-many join table which links [`tags`](./tags.md) to [`entities`](./entities.md).

| Column       | Type  | Description                                                                                                                          |
| ------------ | ----- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `id`         | UUID  | Primary key                                                                                                                          |
| `entity_id`  |       | Foreign key to the [`entities`](./entities.md) table                                                                                 |
| `tag_id`     |       | Foreign key to the [`tags`](./tags.md) table                                                                                         |
| `context_id` |       | Foreign key to the [`contexts`](ui_configurations.md#contexts) table                                                                 |
| `metadata`   | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio", "confidence": 0.92, "annotator": "user123" }`) |

`Metadata` in this context informs about how the tag is used on an entity.
