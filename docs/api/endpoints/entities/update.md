# Update an existing entity

## Parameters

- `id` (UUID, required): Entity ID to update
- `name` (string, optional): New name
- `location` (string, optional): New path or URL
- `is_primary` (boolean, optional): Update primary status
- `metadata` (object, optional): Metadata updates

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/entities/<entity_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "is_primary": false
  }'
```

## Returns

### Response Attributes

- Entity fields with updated values

### Response Body

```json
{
	"id": "UUID",
	"name": "Sketch A",
	"location": "images/sketch-a.jpg",
	"is_primary": false,
	"metadata": {},
	"updated_at": "2025-08-06T12:30:00Z"
}
```

### Error Codes

| HTTP status code | Code         | Description           |
| ---------------- | ------------ | --------------------- |
| 404              | not_found    | Entity not found      |
| 400              | invalid_data | Malformed update data |
