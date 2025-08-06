# Database Reference

This section details the database schema used by the Tagging Service.

- [Entities](entities.md)
- [Tags](./tags.md)
- [Entity Tagging](./entity_tagging.md)
- Utilities
	- [User Interface (UI) Configurations](./utilities/ui_configurations.md)
	- [Ratings](./utilities/ratings.md)
- [Database Schema](./schema.md)

> **Notes**: All foreign keys in this schema refer to UUID primary keys unless otherwise noted.

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
