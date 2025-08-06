# Tags

Tags are the core classification mechanism in the Tagging Service. A tag may represent a person, role, genre, object, emotion, or any other label applicable to an entity or asset.

Tags are **unified** — there is no distinction between a tag for a “Person” and one for a “Genre” at the schema level. Instead, their function is defined by type and relationships.

## General Tag Endpoints

### Create a new tag

- [POST /v1/tags](./tags/create.md)

### Retrieve an existing tag

- [GET /v1/tags/{id}](./tags/retrieve.md)

### List all existing tags

- [GET /v1/tags](./tags/list.md)

### Update an existing tag

- [PATCH /v1/tags/{id}](./tags/update.md)

### Delete an existing tag

- [DELETE /v1/tags/{id}](./tags/delete.md)

## Tag Aliases Endpoints

### View tag aliases

- [GET /v1/tags/{id}/aliases](./tags/{id}/aliases/retrieve.md)

### Update tag aliases

- [PATCH /v1/tags/{id}/aliases](./tags/{id}/aliases/update.md)

## Tag Relationship Endpoints

### View tag relationships

- [GET /v1/tags/{id}/relationships](./tags/{id}/relationships/retrieve.md)

### Update tag relationships

- [PATCH /v1/tags/{id}/relationships](./tags/{id}/relationships/update.md)

## Tag Composition Endpoints

### View tag composition components

- [GET /v1/tags/{id}/compositions](./tags/{id}/compositions/retrieve.md)

### Update tag composition components

- [PATCH /v1/tags/{id}/compositions](./tags/{id}/compositions/update.md)

## Tag Rating Endpoints

### Rate tag in context

- [PATCH /v1/tags/{id}/ratings](./tags/{id}/ratings/contextual.md)

### Rate tag relationship

- [PATCH /v1/tags/{id}/relationship_ratings](./tags/{id}/ratings/relationship.md)
