# Tags

Tags are the core classification mechanism in the Tagging Service. A tag may represent a person, role, genre, object, emotion, or any other label applicable to an entity or asset.

Tags are **unified** — there is no distinction between a tag for a “Person” and one for a “Genre” at the schema level. Instead, their function is defined by type and relationships.

## General Tag Endpoints

| HTTP Method | Endpoint                              | Description              |
| ----------- | ------------------------------------- | ------------------------ |
| POST        | [`/v1/tags`](./tags/create.md)        | Create a new tag         |
| GET         | [`/v1/tags`](./tags/list.md)          | List all existing tags   |
| GET         | [`/v1/tags/{id}`](./tags/retrieve.md) | Retrieve an existing tag |
| PATCH       | [`/v1/tags/{id}`](./tags/update.md)   | Update an existing tag   |
| DELETE      | [`/v1/tags/{id}`](./tags/delete.md)   | Delete an existing tag   |

## Tag Aliases Endpoints

| HTTP Method | Endpoint                                                   | Description        |
| ----------- | ---------------------------------------------------------- | ------------------ |
| GET         | [`/v1/tags/{id}/aliases`](./tags/{id}/aliases/retrieve.md) | View tag aliases   |
| PATCH       | [`/v1/tags/{id}/aliases`](./tags/{id}/aliases/update.md)    | Update tag aliases |

## Tag Relationship Endpoints

| HTTP Method | Endpoint                                                              | Description              |
| ----------- | --------------------------------------------------------------------- | ------------------------ |
| GET         | [`/v1/tags/{id}/relationships`](./tags/{id}/relationships/retrieve.md) | View tag relationships   |
| PATCH       | [`/v1/tags/{id}/relationships`](./tags/{id}/relationships/update.md)  | Update tag relationships |

## Tag Composition Endpoints

| HTTP Method | Endpoint                                                             | Description                       |
| ----------- | -------------------------------------------------------------------- | --------------------------------- |
| GET         | [`/v1/tags/{id}/compositions`](./tags/{id}/compositions/retrieve.md) | View tag composition components   |
| PATCH       | [`/v1/tags/{id}/compositions`](./tags/{id}/compositions/update.md)   | Update tag composition components |

## Tag Rating Endpoints

| HTTP Method | Endpoint                                                                    | Description           |
| ----------- | --------------------------------------------------------------------------- | --------------------- |
| PATCH       | [`/v1/tags/{id}/ratings`](./tags/{id}/ratings/contextual.md)                | Rate tag in context   |
| PATCH       | [`/v1/tags/{id}/relationship_ratings`](./tags/{id}/ratings/relationship.md) | Rate tag relationship |
