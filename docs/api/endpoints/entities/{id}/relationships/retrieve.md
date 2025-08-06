# View entity relationships

## Parameters

- `id` (UUID, required): Entity ID

## Request

```curl
curl -X GET http://localhost:8000/api/v1/entities/<entity_id>/relationships \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `relationships`: List of relationships where this entity is either source or target

  - `entity_a_id`, `entity_b_id`, `relationship_type_id`

### Response Body

```json
{
	"relationships": [
		{
			"entity_a_id": "UUID",
			"entity_b_id": "UUID",
			"relationship_type_id": "UUID"
		}
	]
}
```

### Error Codes

| HTTP status code | Code      | Description      |
| ---------------- | --------- | ---------------- |
| 404              | not_found | Entity not found |
