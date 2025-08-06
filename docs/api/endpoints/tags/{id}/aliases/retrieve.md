# View tag aliases

## Parameters

- `id` (UUID, required): Tag ID

## Request

```curl
curl -X GET http://localhost:8000/api/v1/tags/<tag_id>/aliases \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `aliases`: List of alias objects

  - `id`: UUID
  - `name`: Alias name

### Response Body

```json
{
	"aliases": [
		{
			"id": "UUID",
			"name": "automobile"
		},
		{
			"id": "UUID",
			"name": "vehicle"
		}
	]
}
```

### Error Codes

| HTTP status code | Code      | Description   |
| ---------------- | --------- | ------------- |
| 404              | not_found | Tag not found |
