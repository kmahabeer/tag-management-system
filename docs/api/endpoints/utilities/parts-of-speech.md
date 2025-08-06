# Retrieve all parts of speech

## Parameters

None

## Request

```curl
curl -X GET http://localhost:8000/api/v1/parts-of-speech \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `id` (UUID): POS ID
- `name` (string): Part of speech name (e.g., "noun")
- `description` (string): Explanation
- `is_active` (boolean)

### Response Body

```json
{
	"parts_of_speech": [
		{
			"id": "UUID",
			"name": "noun",
			"description": "A person, place, thing, or concept",
			"is_active": true
		}
	]
}
```

### Error Codes

| HTTP status code | Code | Description |
| ---------------- | ---- | ----------- |
| 200              | ok   | Success     |
