# Retrieve all tagging contexts

## Parameters

None

## Request

```curl
curl -X GET http://localhost:8000/api/v1/contexts \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `id` (UUID): Context ID
- `name` (string): Name of the context (e.g., "content")
- `classification_type` (string): Either `objective` or `subjective`
- `description` (string)
- `is_active` (boolean)

### Response Body

```json
{
	"contexts": [
		{
			"id": "UUID",
			"name": "content",
			"classification_type": "objective",
			"description": "What is visually present",
			"is_active": true
		}
	]
}
```

### Error Codes

| HTTP status code | Code | Description |
| ---------------- | ---- | ----------- |
| 200              | ok   | Success     |
