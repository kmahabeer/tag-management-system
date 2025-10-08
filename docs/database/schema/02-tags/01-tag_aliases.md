---
title: Tag Aliases
parent: Tags
nav_order: 1
---
# Tag Aliases

Defines alternative names or synonyms for canonical tags.  
Aliases improve searchability, tag normalization, and user input flexibility.

## `tag_aliases` Table

Represents alternate names or synonyms that point to a single canonical tag.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Unique alias name used as an alternate label for a tag. |
| `tag_id` | UUID | Foreign key to the [`tags`](../index.md#tags-table) table, referencing the canonical tag this alias points to. |

Example:

| id | name | tag_id |
|----|------|---------|
| a1 | "puppy" | tag_dog |

## Constraints

- Each alias name must be **globally unique**.  
- Each alias must map to exactly **one** canonical tag.  
- A tag may have multiple aliases, but no circular aliasing is allowed (an alias cannot point to another alias).

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Tag Aliases** | `/tags/{id}/aliases` | **GET**, **POST**, **PATCH**, **DELETE** | Manage aliases for a specific tag (e.g., add or edit synonyms). |
| | `/tags/{id}/aliases/{alias_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a single alias linked to the tag. |
| **System-Wide Tag Aliases** | `/tag-aliases`, `/tag-aliases/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage alias records across all tags in the system. |

### CRUD Parity Overview

| Scope | CRUD | Description |
|--------|------|--------------|
| Per Tag | ✅ Full | Manage aliases for individual tags. |
| System-Wide | ✅ Full | Manage all alias records globally. |

✅ = Full REST-compliant CRUD support.
