# Tags

Tags are the core classification mechanism in the Tagging Service. A tag may represent a person, role, genre, object, emotion, or any other label applicable to an entity or asset.

Tags are **unified** — there is no distinction between a tag for a “Person” and one for a “Genre” at the schema level. Instead, their function is defined by type and relationships.

## Endpoints

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
