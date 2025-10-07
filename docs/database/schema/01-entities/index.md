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

## Entity API Endpoint Summary

| Category                                                | Endpoint                                          | CRUD Coverage                  | Description                                                                                                    |
| ------------------------------------------------------- | ------------------------------------------------- | ------------------------------ | -------------------------------------------------------------------------------------------------------------- |
| **Core Entities**                                       | `/entities`                                       | **POST**, **GET**              | Create and list entity records representing digital or physical artifacts.                                     |
|                                                         | `/entities/{id}`                                  | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific entity.                                                                 |
| **Entity Tags**                                         | `/entities/{id}/tags`                             | **GET**, **PATCH**             | Manage tag assignments and contextual tagging for a single entity.                                             |
|                                                         | `/entities/{id}/tags/{tag_id}`                    | **GET**, **DELETE**            | Inspect or unassign a specific tag attached to the entity.                                                     |
| **Entity Relationships**                                | `/entities/{id}/relationships`                    | **GET**, **PATCH**             | Manage entity-to-entity relationships where the current entity acts as the source (Entity A).                  |
|                                                         | `/entities/{id}/relationships/{relationship_id}`  | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a single relationship link between two entities.                                    |
| **Entity Relationship Ratings**                         | `/entities/{id}/relationship_ratings`             | **GET**, **PATCH**             | Apply or update contextual ratings (e.g., confidence, clarity) for relationships originating from this entity. |
|                                                         | `/entities/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Retrieve, modify, or delete a single relationship rating record.                                               |
| **Entity Ratings**                                      | `/entities/{id}/ratings`                          | **GET**, **PATCH**             | Apply or update contextual ratings directly on the entity itself (e.g., quality, accuracy).                    |
|                                                         | `/entities/{id}/ratings/{rating_id}`              | **GET**, **PATCH**, **DELETE** | Retrieve, modify, or delete a single rating record applied to the entity.                                      |
| **Entity Purposes**                                     | `/entities/{id}/purposes`                         | **GET**, **PATCH**             | Assign or update the set of purpose tags describing the entity’s intended function or role.                    |
|                                                         | `/entities/{id}/purposes/{purpose_id}`            | **GET**, **PATCH**, **DELETE** | Retrieve, update, or remove a single purpose tag assignment for the entity.                                    |
| **Entity Versions** *(logical view over relationships)* | `/entities/{id}/versions`                         | **GET**, **PATCH**             | Manage “version-of” relationships linking this entity to its prior, derived, or alternate versions.            |
|                                                         | `/entities/{id}/versions/{version_id}`            | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a specific version link between entities.                                           |
| **System-Wide Relationships**                           | `/entity-relationships`                           | **GET**, **POST**              | Global view for listing and creating entity-to-entity relationships across the system.                         |
|                                                         | `/entity-relationships/{id}`                      | **PATCH**, **DELETE**          | Update or delete a specific relationship record globally.                                                      |
| **System-Wide Relationship Ratings**                    | `/entity-relationship-ratings`                    | **GET**, **POST**              | Global view for listing and creating relationship ratings across all entities.                                 |
|                                                         | `/entity-relationship-ratings/{id}`               | **PATCH**, **DELETE**          | Update or delete a specific relationship rating globally.                                                      |
| **System-Wide Entity Ratings**                          | `/entity-ratings`                                 | **GET**, **POST**              | List or create contextual ratings applied directly to entities across the system.                              |
|                                                         | `/entity-ratings/{id}`                            | **PATCH**, **DELETE**          | Update or remove a specific entity rating record.                                                              |
| **System-Wide Entity Purposes**                         | `/entity-purposes`                                | **GET**, **POST**              | Retrieve or create purpose tag assignments across all entities.                                                |
|                                                         | `/entity-purposes/{id}`                           | **PATCH**, **DELETE**          | Update or delete a specific purpose tag record globally.                                                       |
| **System-Wide Entity Versions**                         | `/entity-versions`                                | **GET**, **POST**              | List or create version-type relationships (“version-of”) across the system.                                    |

### CRUD Parity Overview

| Subsystem            | Collection | Individual | System-Wide | CRUD                              		|
| -------------------- | ---------- | ---------- | ----------- | ------------------------------------- |
| Entities             | ✅          | ✅          | —        | Full                              	|
| Tags                 | ✅          | ✅          | —        | Partial (no POST on sub-resource)		|
| Relationships        | ✅          | ✅          | ✅       	| Full                              	|
| Relationship Ratings | ✅          | ✅          | ✅      	| Full                              	|
| Purposes             | ✅          | ✅          | ✅      	| Full                              	|
| Ratings              | ✅          | ✅          | ✅      	| Full                              	|
| Versions             | ✅          | ✅          | ✅      	| Full                              	|

✅ = full REST-compliant CRUD support (GET, POST, PATCH, DELETE).
