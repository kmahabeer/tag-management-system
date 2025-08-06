# Update entity relationships

## Parameters

- `id` (UUID, required): Entity A ID (source)
- `relationships`: List of relationships to assign

  - `entity_b_id` (UUID, required): Target entity
  - `relationship_type_id` (UUID, required)

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/entities/<entity_id>/relationships \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "relationships": [
      {
        "entity_b_id": "UUID",
        "relationship_type_id": "UUID"
      }
    ]
  }'
```

## Returns

### Response Attributes

- `entity_a_id`: UUID
- `relationships`: List of added relationships

### Response Body

```json
{
	"entity_a_id": "UUID",
	"relationships": [
		{
			"entity_b_id": "UUID",
			"relationship_type_id": "UUID"
		}
	]
}
```

### Error Codes

| HTTP status code | Code         | Description                |
| ---------------- | ------------ | -------------------------- |
| 400              | invalid_data | Invalid input or duplicate |
| 404              | not_found    | Entity not found           |
