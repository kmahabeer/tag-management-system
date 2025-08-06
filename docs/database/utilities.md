# Utilities

## User Interface (UI) Configuration Tables

Although the Tagging Service is **not** responsible for rendering the User Interface (UI), it provides optional configuration tables that help downstream UI systems interpret and organize tags and entities. These tables define how tags should be grouped, typed, and displayed based on an entity’s **purpose**, enabling UIs to automatically render structured metadata without hardcoding display logic. This configuration enables purpose-specific rendering without requiring hardcoded frontend logic per use case.

### `ui_groups` Table

Defines UI-level sections, like “Color”, “Mood”, or “Props”.

| Column | Type | Description                  |
| ------ | ---- | ---------------------------- |
| `id`   | UUID | Primary key                  |
| `name` | TEXT | Display label for UI section |

### `ui_fields` Table

Junction table which maps a **UI section** to a specific **tag category**, enabling downstream UIs to render grouped tag selectors.

| Column                 | Type    | Description                                                                                                                           |
| ---------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                   | UUID    | Primary key                                                                                                                           |
| `ui_group_id`          |         | Foreign key to the [`ui_groups`](./utilities.md#ui_groups-table) table                                                                |
| `tag_relationships_id` |         | Foreign key to a [`tag_relationships`](./tags.md#tag_relationships-table) that represents the category (e.g., `"Color"`, `"Vehicle"`) |
| `control_type`         | TEXT    | UI widget type (e.g., "chips", "dropdown")                                                                                            |
| `is_required`          | BOOLEAN | Whether the UI expects this tag to be supplied                                                                                        |

Tags shown in a control are derived by querying all tags that are related to `tag_category_id` through a `"is a"` relationship.

> [!TODO] TODO
> `ui_fields` references `tag_relationships`, but that indirection might be fragile. Consider allowing UI mapping directly to a `tag_category_id` column in tags or add a `tag_type` = 'category' and map directly.

## Ratings

### `rating_types` Table

| Column          | Type    | Description                                         |
| --------------- | ------- | --------------------------------------------------- |
| `id`            | UUID    | Primary key                                         |
| `name`          | TEXT    | e.g., `"likeness"`, `"confidence"`, `"clarity"`     |
| `is_normalized` | BOOLEAN | Whether the rating should be on a 1–10 or 0–1 scale |

### `ratings` Lookup Table

| Column           | Type | Description                                                                  |
| ---------------- | ---- | ---------------------------------------------------------------------------- |
| `id`             | UUID | Primary key                                                                  |
| `name`           | TEXT |                                                                              |
| `score`          | INT  | Numeric scale                                                                |
| `description`    | TEXT |                                                                              |
| `rating_type_id` |      | Foreign key to the [`rating_types`](./utilities.md#rating_types-table) table |

Rating types add semantic meaning to user selected ratings. This way:

- "10/10" can mean "I love it" if it's a *likeness* rating.
- "9/10" can mean "very clear" if it's a *clarity* rating.
- You can use the same rating table for different semantics based on context.

## Contexts

Tags can be assigned to entities in **different semantic roles**, depending on how they relate to the entity. For example, a tag like `"Dog"` could indicate that a dog is **depicted in** the image (content), that the image is **about** dogs (subject), or that the tag is part of **annotation metadata**.

To support this, the `entity_tags` table includes a `context` field that captures the *purpose or meaning* of each tag assignment.

### `contexts` Lookup Table

| Column                | Type    | Description                 |
| --------------------- | ------- | --------------------------- |
| `id`                  | UUID    | Primary key                 |
| `name`                | TEXT    |                             |
| `classification_type` |         | `subjective` or `objective` |
| `description`         | TEXT    |                             |
| `is_active`           | BOOLEAN |                             |

This enables:

- Assigning the **same tag** to the **same entity** in multiple ways (e.g. an image of Obama having the tag `"Obama"` as both "content" and "subject").
- **Filtering, displaying, or editing tags** differently depending on the workflow or UI configuration.
- **Improved semantic clarity** in how tags are interpreted by downstream systems.

The list of allowed contexts is defined and enforced at the schema level, and can be extended over time. See the table below for currently supported contexts. The following table defines the **Taxonomy of Tagging Contexts**, grouped by *objective* and *subjective* classifications.

### Objective Context Classifications

| Context    | Description                                                                 | Use Case                                   | Example Tags               |
| ---------- | --------------------------------------------------------------------------- | ------------------------------------------ | -------------------------- |
| `subject`  | What the entity is **about**<br>(e.g., themes, topics, high-level meaning)  | A documentary image *about* climate issues | "Climate Change"           |
| `content`  | What is **in** the entity<br>(e.g., visual elements, objects, people shown) | A person or animal shown *in* a photo      | "Person", "Dog"            |
| `metadata` | Descriptive info not from the content, but about the file or process        | Indicates origin or status info            | "Generated", "Low Quality" |

### Subjective Context Classifications

| Context      | Description                                                   | Use Case                                  | Example Tags                  |
| ------------ | ------------------------------------------------------------- | ----------------------------------------- | ----------------------------- |
| `style`      | Artistic or aesthetic style the entity belongs to             | Art, photography, or design               | "Impressionism", "Monochrome" |
| `emotion`    | Emotional tone conveyed by the entity                         | Classifying mood of a film still or image | "Sad", "Joyful"               |
| `annotation` | Notes made by an annotator; subjective or observational input | For human-led review or labeling projects | "Crowd", "Blurry"             |
