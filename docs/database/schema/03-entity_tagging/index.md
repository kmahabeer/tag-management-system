---
title: Entity-Tag Relationships
parent: Database Schema
has_children: false
nav_order: 3
---
# Entity Tagging

The `entity_tags` table links entities and tags, allowing each tag assignment to be contextual and richly annotated.

## `entity_tags` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `entity_id` | UUID | Foreign key to [`entities`](../entities/index.md). |
| `tag_id` | UUID | Foreign key to [`tags`](../tags/index.md). |
| `context_id` | UUID | Foreign key to [`contexts`](../utilities/index.md#contexts). |
| `metadata` | JSONB | Arbitrary metadata about the tag assignment (e.g., `{ "confidence": 0.92, "annotator": "user123" }`). |

This join table enables contextual, multi-layer tagging while keeping the base entities and tag definitions normalized.
