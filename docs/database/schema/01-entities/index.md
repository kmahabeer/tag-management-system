---
title: Entities
parent: Database Schema
has_children: true
nav_order: 1
---
# Entities

The `entities` table represents digital artifacts such as images, video files, documents, URLs, or any tagged resource managed by the system.

## `entities` Table

| Column        | Type      | Description                                                                                     |
| ------------- | --------- | ----------------------------------------------------------------------------------------------- |
| `id`          | UUID      | Primary key.                                                                                   |
| `name`        | TEXT      | Logical or human-friendly name of the entity.                                                  |
| `location`    | TEXT      | File path, URI, or reference to where the resource resides.                                    |
| `is_primary`  | BOOLEAN   | Indicates whether this entity is the canonical (primary) version.                              |
| `metadata`    | JSONB     | Arbitrary key-value metadata about the entity (e.g., `{ "source": "Label Studio" }`).          |

Each entity can have one or more assigned **tags**, associated **purposes**, and **relationships** to other entities.

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Core Entities** | `/entities` | **POST**, **GET** | Create and list entity records. |
| | `/entities/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific entity. |
| **Entity Tags** | `/entities/{id}/tags` | **GET**, **PATCH** | Assign or update contextual tags for a specific entity. |
| | `/entities/{id}/tags/{tag_id}` | **GET**, **DELETE** | Inspect or remove a specific tag assignment. |
| **Entity Relationships** | `/entities/{id}/relationships` | **GET**, **POST**, **PATCH** | Manage entity-to-entity relationships where this entity is the source. |
| | `/entities/{id}/relationships/{relationship_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or delete a single relationship. |
| **Entity Relationship Ratings** | `/entities/{id}/relationship_ratings` | **GET**, **PATCH** | Apply contextual ratings (e.g., clarity, confidence) to outgoing relationships. |
| | `/entities/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or delete a specific relationship rating. |
| **Entity Ratings** | `/entities/{id}/ratings` | **GET**, **PATCH** | Apply contextual ratings directly to this entity (e.g., quality, accuracy). |
| | `/entities/{id}/ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a specific entity rating. |
| **Entity Purposes** | `/entities/{id}/purposes` | **GET**, **PATCH** | Manage purpose tag assignments describing an entity’s role or function. |
| | `/entities/{id}/purposes/{purpose_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a specific purpose tag. |
| **Entity Versions** *(logical view)* | `/entities/{id}/versions` | **GET**, **PATCH** | Manage “version-of” relationships linking this entity to prior or derived versions. |
| | `/entities/{id}/versions/{version_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a specific version link. |
| **System-Wide Relationships** | `/entity-relationships`, `/entity-relationships/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Global view for listing and managing all entity relationships. |
| **System-Wide Relationship Ratings** | `/entity-relationship-ratings`, `/entity-relationship-ratings/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage all relationship rating records globally. |
| **System-Wide Entity Ratings** | `/entity-ratings`, `/entity-ratings/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Global view of contextual ratings applied to entities. |
| **System-Wide Entity Purposes** | `/entity-purposes`, `/entity-purposes/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage global purpose tag assignments across all entities. |
| **System-Wide Entity Versions** | `/entity-versions`, `/entity-versions/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage version-type relationships across entities system-wide. |

### CRUD Parity Overview

| Subsystem | Collection | Individual | System-Wide |
|------------|------------|-------------|-------------|
| Entities | ✅ | ✅ | — |
| Relationships | ✅ | ✅ | ✅ |
| Relationship Ratings | ✅ | ✅ | ✅ |
| Purposes | ✅ | ✅ | ✅ |
| Ratings | ✅ | ✅ | ✅ |
| Versions | ✅ | ✅ | ✅ |

✅ = Full REST-compliant CRUD coverage (GET, POST, PATCH, DELETE).
