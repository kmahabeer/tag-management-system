# User Interface (UI) Configuration

Although the Tagging Service is **not** responsible for rendering the User Interface (UI), it provides optional configuration tables that help downstream UI systems interpret and organize tags and entities. These tables define how tags should be grouped, typed, and displayed based on an entity’s **purpose**, enabling UIs to automatically render structured metadata without hardcoding display logic. This configuration enables purpose-specific rendering without requiring hardcoded frontend logic per use case.

## UI Tables

### `ui_groups` Table

Defines UI-level sections, like “Color”, “Mood”, or “Props”.

| Column | Type | Description                  |
|--------|------|------------------------------|
| `id`   | UUID | Primary key                  |
| `name` | TEXT | Display label for UI section |

### `ui_fields` Table

Junction table which maps `tag_types` to UI controls and grouping logic.

| Column         | Type    | Description                                    |
| -------------- | ------- | ---------------------------------------------- |
| `id`           | UUID    | Primary key                                    |
| `ui_group_id`  | UUID    | Foreign key to `ui_groups`                     |
| `tag_type_id`  | UUID    | Foreign key to `tag_types` (optional)          |
| `control_type` | TEXT    | UI widget type (e.g., "chips", "dropdown")     |
| `is_required`  | BOOLEAN | Whether the UI expects this tag to be supplied |
