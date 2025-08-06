# Update purpose(s) for an entity

## Parameters

- `id` (UUID, required): Entity ID
- `purposes`: List of purpose tag IDs to assign

  - `purpose_tag_id` (UUID, required)
  - `is_primary` (boolean, optional)

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/entities/<entity_id>/purpose \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "purposes": [
      { "purpose_tag_id": "UUID", "is_primary": true },
      { "purpose_tag_id": "UUID", "is_primary": false }
    ]
  }'
```

## Returns

### Response Attributes

- `entity_id`: UUID
- `purposes`: Array of assigned purposes (with `is_primary` flags)

### Response Body

```json
{
	"entity_id": "UUID",
	"purposes": [
		{ "purpose_tag_id": "UUID", "is_primary": true },
		{ "purpose_tag_id": "UUID", "is_primary": false }
	]
}
```

### Error Codes

| HTTP status code | Code         | Description                       |
| ---------------- | ------------ | --------------------------------- |
| 400              | invalid_data | Invalid tag or multiple primaries |
| 404              | not_found    | Entity or tag not found           |
