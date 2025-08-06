# View tag composition components

## Parameters

- `id` (UUID, required): Composite tag ID

## Request

```curl
curl -X GET http://localhost:8000/api/v1/tags/<tag_id>/compositions \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `components`: Ordered list of component tags in the composition

  - `id`: UUID
  - `name`: Component tag name
  - `position`: Integer index in the composition

### Response Body

```json
{
	"components": [
		{ "id": "UUID", "name": "very", "position": 1 },
		{ "id": "UUID", "name": "big", "position": 2 },
		{ "id": "UUID", "name": "car", "position": 3 }
	]
}
```

### Error Codes

| HTTP status code | Code      | Description           |
| ---------------- | --------- | --------------------- |
| 404              | not_found | Composite tag missing |
