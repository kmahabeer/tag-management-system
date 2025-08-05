# Entities

Entities represent digital artifacts that are being described or classified by tags. This includes images, documents, video stills, URLs, and more.

## Endpoints

### Create a new entity

- [POST /v1/entities](./entities/create.md)

### Retrieve an existing entity

- [GET /v1/entities/{id}](./entities/retrieve.md)

### List all existing entities

- [GET /v1/entities](./entities/list.md)

### Update an existing entity

- [PATCH /v1/entities/{id}](./entities/update.md)

### Delete an existing entity

- [DELETE /v1/entities/{id}](./entities/delete.md)

### Tag(s) assigned to an entity

#### List all tags assigned to an entity

- [GET /v1/entities/{id}/tags](./entities/{id}/tags/assign.md)

#### Assign or unassign tag(s) to an entity

- [PATCH /v1/entities/{id}/tags](./entities/{id}/tags/update.md)
