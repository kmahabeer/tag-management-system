# Update tag aliases

## Parameters

- `id` (UUID, required): Tag ID
- `aliases`: List of strings representing alias names to associate with this tag

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id>/aliases \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "aliases": ["automobile", "vehicle"]
  }'
```

## Returns

### Response Attributes

- `aliases`: List of updated alias records

### Response Body

```json
{
	"aliases": [
		{ "id": "UUID", "name": "automobile" },
		{ "id": "UUID", "name": "vehicle" }
	]
}
```

### Error Codes

| HTTP status code | Code         | Description         |
| ---------------- | ------------ | ------------------- |
| 400              | invalid_data | Invalid alias input |
| 404              | not_found    | Tag not found       |
