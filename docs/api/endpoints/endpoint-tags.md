# Endpoint: `/tags`

The `/tags` endpoints provide full CRUD access to the tag database.

## GET `/tags`

List all tags.

### Query Parameters

- `type`: (optional) Filter by tag type
- `q`: (optional) Fuzzy search by name or alias

### Example Response

```json
{
  "data": [
    { "id": "uuid1", "name": "Red", "types": ["Color"] },
    { "id": "uuid2", "name": "Car", "types": ["Vehicle"] }
  ]
}
```

## POST `/tags`

Create a new tag.

### Example Request Body

```json
{
  "name": "Director",
  "type_ids": ["uuid-role"]
}
```

### Example Response

```json
{
  "data": { "id": "uuid123", "name": "Director" }
}
```

## PATCH `/tags/{tag_id}`

Update an existing tag.

### Example Request Body

```json
{
  "name": "Director",
  "metadata": { "source": "user_input" }
}
```

## DELETE `/tags/{tag_id}`

Permanently delete a tag.
