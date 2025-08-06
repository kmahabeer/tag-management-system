# Update tag composition components

## Parameters

- `id` (UUID, required): Composite tag ID
- `components`: Ordered list of tag component IDs to compose the tag

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id>/compositions \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "components": [
      { "id": "UUID1", "position": 1 },
      { "id": "UUID2", "position": 2 },
      { "id": "UUID3", "position": 3 }
    ]
  }'
```

## Returns

### Response Attributes

- `base_tag_id`: UUID of the composite tag
- `components`: Updated ordered list

### Response Body

```json
{
	"base_tag_id": "UUID",
	"components": [
		{ "id": "UUID1", "name": "very", "position": 1 },
		{ "id": "UUID2", "name": "big", "position": 2 },
		{ "id": "UUID3", "name": "car", "position": 3 }
	]
}
```

### Error Codes

| HTTP status code | Code         | Description                   |
| ---------------- | ------------ | ----------------------------- |
| 400              | invalid_data | Invalid composition structure |
| 404              | not_found    | Tag or component not found    |
