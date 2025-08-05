# Database Reference

This section details the database schema used by the Tagging Service.

- [Entities](entities.md)
- [Tags](./tags.md)
- [Entity Tagging](./entity_tagging.md)
- [UI Configurations](./ui_configuration.md)
- [Database Schema](./schema.md)

> **Note:** Every table in the system includes audit columns: `created_at`, `updated_at` (both of type `TIMESTAMP`), and `created_by`, `updated_by` to track when records are created and last modified. These fields are maintained automatically by the service layer or database triggers.
