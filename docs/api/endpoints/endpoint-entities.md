# Endpoint: `/entities`

The `/entities` endpoints manage tagged resources.

## GET `/entities`

Returns all registered entities.

### Query Parameters

- `purpose`: (optional) Filter by purpose tag
- `q`: (optional) Filter by name

## POST `/entities`

Create a new entity.

### Example Request Body

```json
{
  "name": "fargo_still_01.jpg",
  "location": "/media/fargo_still_01.jpg"
}
```

### Example Response

```json
{
  "data": { "id": "uuid-entity", "name": "fargo_still_01.jpg" }
}
```

## GET `/entities/{id}`

Fetch a specific entity, including its tags and purposes.

## POST `/entities/{id}/tags`

Assign one or more tags to the entity.

### Example Request Body

```json
{
  "tag_ids": ["uuid-red", "uuid-car"]
}
```

## DELETE `/entities/{id}/tags`

Unassign one or more tags from the entity.

### Example Request Body

```json
{
  "tag_ids": ["uuid-car"]
}
```
