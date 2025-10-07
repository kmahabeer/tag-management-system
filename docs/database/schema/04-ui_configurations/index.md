---
title: User Interface (UI) Configurations
parent: Database Schema
nav_order: 4
---
# User Interface (UI) Configurations

Although the Tag Management System does not render the UI itself, it provides configuration tables to support **structured grouping of tag assignments** (`entity_tags`) in user interfaces, based on:

1. The **semantic context** of each tag assignment (e.g., `"content"`, `"subject"`, `"style"`)
2. The **entity’s primary purpose** (as defined in the `entity_purposes` table)
3. A **predefined layout** for that purpose, defining how tags in each context should be grouped

This system allows downstream UIs to display tags in a meaningful and repeatable way across different workflows, without hardcoded logic.

## `ui_layouts` Table

Defines a **layout configuration profile** associated with a specific purpose tag (or used as a fallback default when no purpose is defined).

| Column           | Type | Description                                                                                     |
| ---------------- | ---- | ----------------------------------------------------------------------------------------------- |
| `id`             | UUID | Primary key                                                                                     |
| `purpose_tag_id` | UUID | Foreign key to `tags` table (e.g., `"Product Catalog"`, `"Concept Art"`). Null = default layout |
| `name`           | TEXT | Optional name for debugging or frontend reference                                               |

## `ui_groups` Table

Defines user-facing section labels that group tags visually.

| Column | Type | Description                                                                 |
| ------ | ---- | --------------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                                 |
| `name` | TEXT | Label shown in the user interface (e.g., `"Photographer"`, `"Ingredients"`) |

## `ui_fields` Table

Defines how `entity_tags` of a specific **context** and **category** should be grouped, *within a layout*. This allows the UI to dynamically organize and render assigned tags in semantically meaningful groups, driven by context.

| Column            | Type | Description                                                                                                            |
| ----------------- | ---- | ---------------------------------------------------------------------------------------------------------------------- |
| `id`              | UUID | Primary key                                                                                                            |
| `ui_layout_id`    | UUID | Foreign key to the [`ui_layouts`](ui_configurations.md#ui_layouts-table) table                                               |
| `ui_group_id`     | UUID | Foreign key to the [`ui_groups`](ui_configurations.md#ui_groups-table), used as the section label                            |
| `context_id`      | UUID | Foreign key to the [`contexts`](ui_configurations.md#contexts) table                                                         |
| `category_tag_id` | UUID | Foreign key to the [`tags`](tags.md#tags-table) table — defines the **category parent tag** used to group child tags |
| `sort_order`      | INT  | Display order of this section within the context                                                                       |

## How UI configuration works

To render an entity's assigned tags in the UI:

1. Retrieve all `entity_tags` for the entity
2. Group them by their `context` (e.g., `content`, `subject`, `style`)
3. For each context:
	1. Load the entity’s `primary purpose` via `entity_purposes`
	2. Find the corresponding `ui_layout` (or default layout if none exists)
	3. Filter all `ui_fields` for that `context_id` and layout
	4. For each `ui_field`:
		- Filter `entity_tags` that are descendants of `category_tag_id`
		- Group them under the label defined by `ui_group_id`, respecting `sort_order`

## Fallback: Default Layout

If no purpose is assigned to an entity, the system falls back to a layout where `purpose_tag_id` IS NULL. This layout can define a minimal or generic grouping structure.

## Example: Concept Art

Suppose the entity is tagged with the purpose `"Concept Art"`. The UI layout for this purpose is designed to organize tags in a way that reflects how concept artists and art directors typically explore and retrieve reference material.

The `ui_layout` configuration for `"Concept Art"` might define the following:

```json
{
  "purpose_tag": "Concept Art",
  "ui_layout": [
    {
      "context": "content",
      "groups": [
        {
          "label": "Subjects",
          "category_tag": "Character Type",
          "sort_order": 1
        },
        {
          "label": "Environment",
          "category_tag": "Environment Type",
          "sort_order": 2
        }
      ]
    },
    {
      "context": "style",
      "groups": [
        {
          "label": "Rendering",
          "category_tag": "Rendering Style",
          "sort_order": 1
        }
      ]
    },
    {
      "context": "emotion",
      "groups": [
        {
          "label": "Mood",
          "category_tag": "Mood",
          "sort_order": 1
        }
      ]
    },
    {
      "context": "metadata",
      "groups": [
        {
          "label": "Artist",
          "category_tag": "Artist",
          "sort_order": 1
        }
      ]
    }
  ]
}
```

with the `entity_tags` assignments as follows:

```json
[
  { "tag": "Alien", "context": "content", "category": "Character Type" },
  { "tag": "Human", "context": "content", "category": "Character Type" },
  { "tag": "Jungle", "context": "content", "category": "Environment Type" },
  { "tag": "3D Render", context": "style", "category": "Rendering Style" },
  { "tag": "Whimsical", "context": "emotion",  "category": "Mood" },
  { "tag": "Jane Doe", "context": "metadata", "category": "Artist" }
]
```

resulting in the following UI rendering result for a specific entity:

```txt
[Content]
Subjects:
- Alien
- Human

Environment:
- Jungle

[Style]
Rendering:
- 3D Render

[Emotion]
Mood:
- Whimsical

[Metadata]
Artist:
- Jane Doe
```

### Key Takeaways

- The **entity’s purpose** (`"Concept Art"`) determines which `ui_layout` is used.
- Within that layout, the UI **groups by `context` first**, then follows the groupings defined in `ui_fields`.
- Each grouping rule is based on a **category tag** like `"Character Type"` or `"Environment Type"`; any tag related to it (via `"is a"` or similar relationships) and assigned to the entity under that context is included in the group.
- If no layout exists for the purpose, the system falls back to the default layout (`ui_layout.purpose_tag_id IS NULL`).

## API Endpoint Summary

| Category       | Endpoint           | CRUD Coverage                  | Description                                                                                                    |
| -------------- | ------------------ | ------------------------------ | -------------------------------------------------------------------------------------------------------------- |
| **UI Layouts** | `/ui/layouts`      | **GET**, **POST**              | Retrieve or create layout configuration profiles that determine how tags and contexts are organized in the UI. |
|                | `/ui/layouts/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific UI layout.                                                              |
| **UI Groups**  | `/ui/groups`       | **GET**, **POST**              | Retrieve or create reusable labeled sections for organizing UI layouts.                                        |
|                | `/ui/groups/{id}`  | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific UI group.                                                               |
| **UI Fields**  | `/ui/fields`       | **GET**, **POST**              | Retrieve or create field definitions that link tags, contexts, or categories to sections within a layout.      |
|                | `/ui/fields/{id}`  | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific UI field record.                                                        |

## Coverage Summary

| UI Component | CRUD   | Description                                                      |
| ------------ | ------ | ---------------------------------------------------------------- |
| Layouts      | ✅ Full | Define the top-level structure of tag groupings in the UI.       |
| Groups       | ✅ Full | Define reusable labeled sections used inside layouts.            |
| Fields       | ✅ Full | Bind tags, contexts, and categories to specific layout sections. |
