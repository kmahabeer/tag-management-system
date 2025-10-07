---
title: Endpoints - Entities
parent: API Endpoints
has_children: true
nav_order: 1
---
# Entities

Entities represent digital artifacts that are being described or classified by tags. This includes images, documents, video stills, URLs, and more.

## General Entity Endpoints

| HTTP Method | Endpoint                                      | Description                 |
| ----------- | --------------------------------------------- | --------------------------- |
| POST        | [`/v1/entities`](./entities/create.md)        | Create a new entity         |
| GET         | [`/v1/entities`](./entities/list.md)          | List all existing entities  |
| GET         | [`/v1/entities/{id}`](./entities/retrieve.md) | Retrieve an existing entity |
| PATCH       | [`/v1/entities/{id}`](./entities/update.md)   | Update an existing entity   |
| DELETE      | [`/v1/entities/{id}`](./entities/delete.md)   | Delete an existing entity   |

## Tag(s) assigned to an entity

| HTTP Method | Endpoint                                                   | Description                            |
| ----------- | ---------------------------------------------------------- | -------------------------------------- |
| GET         | [`/v1/entities/{id}/tags`](./entities/{id}/tags/assign.md) | List all tags assigned to an entity    |
| PATCH       | [`/v1/entities/{id}/tags`](./entities/{id}/tags/update.md) | Assign or unassign tag(s) to an entity |

## Entity Purpose(s)

| HTTP Method | Endpoint                                                         | Description                     |
| ----------- | ---------------------------------------------------------------- | ------------------------------- |
| PATCH       | [`/v1/entities/{id}/purpose`](./entities/{id}/purpose/update.md) | Update purpose(s) for an entity |

## Entity Relationship(s)

| HTTP Method | Endpoint                                                                       | Description                 |
| ----------- | ------------------------------------------------------------------------------ | --------------------------- |
| GET         | [`/v1/entities/{id}/relationships`](./entities/{id}/relationships/retrieve.md) | View entity relationships   |
| PATCH       | [`/v1/entities/{id}/relationships`](./entities/{id}/relationships/update.md)   | Update entity relationships |
