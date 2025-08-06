# Database Reference

This section details the database schema used by the Tagging Service.

- [Entities](entities.md)
- [Tags](./tags.md)
- [Entity Tagging](./entity_tagging.md)
- [Utilities](./database/utilities.md)
- [Database Schema](./schema.md)

> **Note:** Every table in the system includes audit columns: `created_at`, `updated_at` (both of type `TIMESTAMP`), and `created_by`, `updated_by` to track when records are created and last modified. These fields are maintained automatically by the service layer or database triggers.

Use a trigger-based approach and apply to all tables with a common function:

```sql
CREATE OR REPLACE FUNCTION set_audit_fields()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```
