---
title: Entities
parent: Database Schema
has_children: true
nav_order: 1
---
# Entities

The `entities` table represents digital artifacts such as image files, video files, documents, URLs, or any tagged resource.

## `entities` Table

| Column        | Type      | Description   |
| ------------- | --------- | ------------- |
| `id`          | UUID      | Primary key   |
| `name`        | TEXT      | The name of the entity |
| `location`    | TEXT      | Location of the entity (e.g., file path, URI, URL) |
| `is_primary`  | BOOLEAN   | Indicates if this is the primary entity (e.g., latest version of an entity) or if there may exist another version that is the primary. |
| `metadata`    | JSONB     | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`) |

Each **entity** can have one or more associated **purposes**, and is linked to [`tags`](./tags.md) via the `entity_tags` junction table.
