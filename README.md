## Tagging-service

### Documentation

This system is designed to power a flexible and context-aware Inspector UI for creative assets like images, film stills, documents, and more. Each asset may serve different "purposes" (e.g. a film still, a figure study, a reference photo) and should dynamically display its associated metadata (tags) in a structured way according to those purposes.

At the heart of the system is a flexible tagging model with many-to-many relationships, hierarchical categorization, and purpose-specific rendering logic. Tags can represent people, roles, vehicles, genres, purposes, and any other useful classification. These tags can be reused and linked across assets and entities like films or subjects.

Here’s a high-level summary of the key tables and what they represent:

- **assets**: The core table representing images, documents, etc. Each asset can have one or more purposes.
- **tags**: A universal table for all tag values (e.g., "Pulp Fiction", "Quentin Tarantino", "Actor"). Tags can belong to multiple tag types.
- **tag_types**: Defines the role or category of a tag (e.g., "Person", "Role", "Genre", "Purpose").
- **tag_relationships**: Supports hierarchy and inheritance between tags (e.g., "Datsun 240Z" is a child of "Car").
- Junction Tables
	- **tag_type_links**: A junction table linking the `tags` table to the `tag_types` table. A single tag can have multiple tag types.
	- **asset_tags**: Junction table linking `tags` to `assets`.
	- **asset_purposes**: Junction table allowing assets to have multiple purposes, such as primary and secondary purposes.

This model supports the reuse of tags across different domains (e.g., "Quentin Tarantino" as a tag used on both assets and the movie entity), and supports filtering, hierarchical exploration, and custom UI rendering depending on the asset’s purpose.

The Inspector UI will read structured JSON formatted according to predefined layouts for each purpose (e.g., "Film Still" might expect to show Directors, Writers, Actors, etc.). Tags that do not fit into this predefined structure are shown in a catch-all "Other Tags" section.

### Database

#### Core Tables

##### `assets`

Stores metadata for digital assets.

- `id` (UUID): Primary key
- `filename` (TEXT): File name
- `created_at` / `updated_at` (TIMESTAMP): Timestamps

##### `tags`

Stores all tags (e.g., people, concepts, colors, emotions, etc.)

- `id` (UUID): Primary key
- `name` (TEXT): Unique canonical name
- `created_at` / `updated_at` (TIMESTAMP): Timestamps

##### `tag_types`

Used to classify tags (e.g., Color, Person, Mood, Genre).

- `id` (UUID): Primary key
- `name` (TEXT): Unique type name

##### `tag_tag_types`

Many-to-many table linking tags to one or more types (e.g., a tag can be both a Mood and a Genre).

- `tag_id`, `tag_type_id`: Foreign keys
- Primary key: (`tag_id`, `tag_type_id`)

##### `tag_aliases`

Alternative names or synonyms for a tag. Used in UI for flexible searching and filtering.

- `id` (UUID): Primary key
- `tag_id` (UUID): Foreign key to canonical tag
- `alias` (TEXT): Alias string (must be unique across all aliases)

##### `tag_relationship_types`

Defines the type of relationship between tags.

- `id` (UUID): Primary key
- `name` (TEXT): Enum-style text value (e.g., 'parent', 'related')

##### `tag_relationships`

Defines typed relationships between tags, like hierarchy or equivalence.

- `source_tag_id`, `target_tag_id`: Foreign keys to `tags`
- `relationship_type_id`: Foreign key to `tag_relationship_types`
- Primary key: (`source_tag_id`, `target_tag_id`, `relationship_type_id`)

##### `tag_modifiers`

Defines tag phrases like "very tall" or "bright red".

- `modifier_tag_id`: The adjective or adverb (e.g., 'very')
- `base_tag_id`: The main tag being modified (e.g., 'tall')
- Combined in UI or logic to form virtual phrases

##### `asset_tags`

Links tags to assets.

- `asset_id`, `tag_id`: Foreign keys
- Optional `modifier_tag_id`: To describe e.g., 'very tall'
- Primary key: (`asset_id`, `tag_id`)

##### `asset_purposes`

Defines one or more purposes for an asset.

- `purpose_tag_id`: Must link to a tag that represents a purpose
- `is_primary`: Indicates main purpose

---

#### UI Configuration Tables

##### `ui_groups`

Defines groups of tags to be bundled together visually in the Inspector UI.

- `id` (UUID): Primary key
- `name` (TEXT): Display name for UI section (e.g., "Color", "Mood")

##### `ui_fields`

Describes how tags of a given type or group are displayed/edited in the UI.

- `id` (UUID): Primary key
- `ui_group_id` (UUID): Foreign key
- `tag_type_id` (UUID): Foreign key (optional)
- `control_type` (TEXT): UI control (e.g., dropdown, chips, slider)
- `is_required` (BOOLEAN): Whether user must supply this

---

#### Relationships & Semantics

- A tag can have multiple aliases (e.g., 'strong' for 'powerful').
- Tags may have relationships of type "parent", "related", or others.
- Tags can be composed via `tag_modifiers` (e.g., "slightly sad").
- Tags may belong to multiple `tag_types`, and types can map to `ui_fields`.

---

#### Example Use Case

A user uploads an image of a red car and tags it with:

- Tag: `car` (type: Object)
- Tag: `red` (type: Color)
- Modifier: `bright` → `bright red`

The UI renders this under the "Color" section using a chip control and under "Objects" with a dropdown. Aliases like "scarlet" will match the same tag. Related tags like `vehicle` (parent of `car`) and `transport` (related) are shown in the inspector.

#### 🎯 Concepts & Goals

1. **Unified Tag Table**: All labels—people, roles, genres, objects, purposes—live in a single `tags` table.
2. **Tag Types**: Each tag is classified by one or more `tag_types` (e.g., `person`, `role`, `genre`, `purpose`, `vehicle`).
3. **Tag Relationships**: Tags can be linked hierarchically or associatively via `tag_relationships` (e.g., `movie` → `director`, `movie` → `actor`).
4. **Asset‐Tagging**: Assets reference tags (what’s in the asset) via `asset_tags`.
5. **Asset Purposes**: Assets can have multiple purposes, with exactly one marked primary, via `asset_purposes`.
6. **UI Configuration**: The Inspector UI uses pre‐defined group/field mappings based on an asset’s selected purpose to render structured metadata, falling back to `other_tags` for uncategorized tags.

---

#### 🗃️ Database Schema

```sql
-- Assets
CREATE TABLE assets (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  filename TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tags (people, roles, genres, purposes, objects…)
CREATE TABLE tags (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tag Types (classification for UI grouping/filtering)
CREATE TABLE tag_types (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Many-to-many: tags ↔ tag_types
CREATE TABLE tag_tag_types (
  tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
  tag_type_id UUID REFERENCES tag_types(id) ON DELETE CASCADE,
  PRIMARY KEY(tag_id, tag_type_id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Hierarchical/associative tag relationships
CREATE TABLE tag_relationships (
  parent_tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
  child_tag_id  UUID REFERENCES tags(id) ON DELETE CASCADE,
  relationship  TEXT NOT NULL,
  PRIMARY KEY(parent_tag_id, child_tag_id, relationship),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Asset ↔ Tag assignments (what’s in or about the asset)
CREATE TABLE asset_tags (
  asset_id UUID REFERENCES assets(id) ON DELETE CASCADE,
  tag_id   UUID REFERENCES tags(id) ON DELETE CASCADE,
  PRIMARY KEY(asset_id, tag_id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Asset Purposes (supports multiple, with one primary)
CREATE TABLE asset_purposes (
  asset_id       UUID REFERENCES assets(id) ON DELETE CASCADE,
  purpose_tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
  is_primary     BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY(asset_id, purpose_tag_id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

---

#### 🛠️ Sample Data

Below is mock data illustrating **many-to-many**, **hierarchies**, and **primary/secondary purposes**.

##### 1. Assets

```text
id | filename
---|------------------------------
1  | fargo_still_01.jpg
2  | pulpfiction_suitcase.jpg
3  | hateful_eight_table.jpg
4  | datsun_240z.jpg
5  | figure_study_001.jpg
```

##### 2. Tags

```text
id | name
---|--------------------------
1  | Quentin Tarantino
2  | Joel Coen
3  | Ethan Coen
4  | William H. Macy
5  | Frances McDormand
6  | Steve Buscemi
7  | Samuel L. Jackson
8  | Kurt Russell
9  | Film Still
10 | Script
11 | Figure Study
12 | Pulp Fiction
13 | The Hateful Eight
14 | Director
15 | Writer
16 | Actor
17 | Genre
18 | Crime
19 | Dark Comedy
20 | Drama
21 | Western
22 | Datsun 240z
23 | Car
24 | Vehicle
25 | Subject
26 | Character
27 | Prop
28 | Costume
29 | Movie
```

##### 3. tag_types

```text
id | name
---|---------
1  | person
2  | role
3  | genre
4  | purpose
5  | vehicle
6  | subject
7  | character
8  | prop
9  | costume
10 | movie
```

##### 4. tag_tag_types

```text
tag_id | tag_type_id
-------|-------------
1 (QT)      | 1 (person)
1           | 2 (role)
2 (Joel)    | 1
3 (Ethan)   | 1
4 (WHM)     | 1
5 (FM)      | 1
6 (SB)      | 1
7 (SLJ)     | 1
8 (KR)      | 1
14 (Director) | 2
15 (Writer)   | 2
16 (Actor)    | 2
17 (Genre)    | 3
18 (Crime)    | 3
19 (Dark Comedy) | 3
20 (Drama)      | 3
21 (Western)    | 3
9  (Film Still)| 4
10 (Script)     | 4
11 (Figure Study)| 4
22 (Datsun 240z)| 5
23 (Car)        | 5
24 (Vehicle)    | 5
25 (Subject)    | 6
26 (Character)  | 7
27 (Prop)       | 8
28 (Costume)    | 9
12 (Pulp Fiction)       | 10
13 (The Hateful Eight)  | 10
29 (Movie)              | 10
```

##### 5. tag_relationships

```text
parent_tag_id | child_tag_id | relationship
--------------|--------------|-------------
14 (Director) | 1 (QT)       | role-of
14            | 2 (Joel)     | role-of
14            | 3 (Ethan)    | role-of
15 (Writer)   | 1            | role-of
15            | 3            | role-of
16 (Actor)    | 4            | role-of
16            | 5            | role-of
16            | 6            | role-of
16            | 7            | role-of
16            | 8            | role-of
12 (Pulp Fiction) | 1        | movie-of
12                  | 7      | movie-of
12                  | 6      | movie-of
12                  | 4      | movie-of
12                  | 5      | movie-of
12                  | 19     | movie-of
12                  | 18     | movie-of
12                  | 20     | movie-of
13 (The Hateful Eight) | 1    | movie-of
13                     | 8    | movie-of
13                     | 7    | movie-of
13                     | 16   | movie-of
13                     | 21   | movie-of
29 (Movie)           | 12      | category-of
29                   | 13      | category-of
```

##### 6. asset_tags

```text
asset_id | tag_id
---------|-------
1        | 2 (Joel Coen)
1        | 3 (Ethan Coen)
1        | 4 (William H. Macy)
1        | 5 (Frances McDormand)
1        | 6 (Steve Buscemi)
1        | 7 (Samuel L. Jackson)
1        | 9 (Film Still)
1        | 12 (Pulp Fiction)
1        | 26 (Character) -- e.g. multiple instances in child table for each character
1        | 28 (Costume)      -- e.g. "Black Suit"
1        | 27 (Prop)         -- e.g. "Gun"

2        | 1 (Quentin Tarantino)
2        | 7 (Samuel L. Jackson)
2        | 9 (Film Still)
2        | 12 (Pulp Fiction)

3        | 7 (Samuel L. Jackson)
3        | 8 (Kurt Russell)
3        | 9 (Film Still)
3        | 13 (The Hateful Eight)

4        | 22 (Datsun 240z)
4        | 23 (Car)
4        | 24 (Vehicle)

5        | 25 (Subject)     -- e.g., "Model"
5        | 11 (Figure Study)
5        | 26 (Character)? -- or different subject tags
```

##### 7. asset_purposes

```text
asset_id | purpose_tag_id | is_primary
---------|----------------|-----------
1        | 9 (Film Still)   | true
1        | 29 (Movie reference) | false
2        | 9                | true
2        | 29               | false
3        | 9                | true
4        | 11 (Figure Study)| true
5        | 11               | true
```

---

#### 🔍 Structured Inspector JSON

For **asset 1** (Fargo still):

```json
{
  "asset_id": 1,
  "filename": "fargo_still_01.jpg",
  "available_purposes": ["Film Still", "Movie reference"],
  "selected_purpose": "Film Still",
  "movie": {
    "title": "Pulp Fiction",
    "director": ["Quentin Tarantino", "Joel Coen", "Ethan Coen"],
    "writer": ["Quentin Tarantino", "Joel Coen", "Ethan Coen"],
    "actors": [
      "William H. Macy",
      "Frances McDormand",
      "Steve Buscemi",
      "Samuel L. Jackson"
    ],
    "genre": ["Dark Comedy", "Crime", "Drama"]
  },
  "content_tags": {
    "subject": [
      "William H. Macy",
      "Frances McDormand",
      "Steve Buscemi",
      "Samuel L. Jackson"
    ],
    "characters": ["Vincent Vega", "Jules Winnfield"],
    "costumes": ["Black Suit"],
    "props": ["Gun"]
  },
  "other_tags": []
}
```

For **asset 3** (Hateful Eight still):

```json
{
  "asset_id": 3,
  "filename": "hateful_eight_table.jpg",
  "available_purposes": ["Film Still"],
  "selected_purpose": "Film Still",
  "movie": {
    "title": "The Hateful Eight",
    "director": ["Quentin Tarantino"],
    "writer": ["Quentin Tarantino"],
    "actors": ["Kurt Russell", "Samuel L. Jackson"],
    "genre": ["Western"]
  },
  "content_tags": {
    "subject": ["Kurt Russell", "Samuel L. Jackson"],
    "props": ["Table"],
    "setting": ["Winter"]
  },
  "other_tags": []
}
```

For **asset 4** (Car photo):

```json
{
  "asset_id": 4,
  "filename": "datsun_240z.jpg",
  "available_purposes": ["Figure Study", "Movie reference"],
  "selected_purpose": "Figure Study",
  "content_tags": {
    "subject": ["Datsun 240z"],
    "vehicle": ["Car", "Vehicle"]
  },
  "other_tags": []
}
```

For **asset 5** (Figure study):

```json
{
  "asset_id": 5,
  "filename": "figure_study_001.jpg",
  "available_purposes": ["Figure Study"],
  "selected_purpose": "Figure Study",
  "content_tags": {
    "subject": ["Model"],
    "pose": [],
    "lighting": [],
    "setting": []
  },
  "other_tags": []
}
```

---

#### 🛢️ SQL Query Examples

1. **Fetch asset’s tags for a given purpose**

```sql
SELECT t.name, tt.name AS type
FROM asset_purposes ap
JOIN assets a        ON a.id = ap.asset_id
JOIN asset_tags at   ON at.asset_id = a.id
JOIN tags t          ON t.id = at.tag_id
JOIN tag_tag_types ttt ON t.id = ttt.tag_id
JOIN tag_types tt    ON tt.id = ttt.tag_type_id
WHERE a.id = 1 AND ap.is_primary = TRUE
ORDER BY tt.name, t.name;
```

1. **Assemble movie section (child tags of the movie tag)**

```sql
WITH movie_tag AS (
  SELECT tag_id FROM asset_tags WHERE asset_id = 1 AND tag_id IN (SELECT id FROM tags WHERE name = 'Pulp Fiction')
)
SELECT rel.relationship, child.name
FROM movie_tag m
JOIN tag_relationships rel ON rel.parent_tag_id = m.tag_id
JOIN tags child           ON child.id = rel.child_tag_id;
```

1. **List available purposes for an asset**

```sql
SELECT t.name, ap.is_primary
FROM asset_purposes ap
JOIN tags t ON t.id = ap.purpose_tag_id
WHERE ap.asset_id = 1;
```

1. **Fetch uncategorized tags (other_tags)**

```sql
-- Assuming you have logic to collect categorized tag_ids
WITH categorized AS (
  SELECT tag_id FROM asset_tags WHERE asset_id = 1 AND tag_id IN (
    -- list of tag_ids used in 'movie' or 'content_tags'
  )
)
SELECT t.name
FROM asset_tags at
JOIN tags t ON t.id = at.tag_id
WHERE at.asset_id = 1
  AND at.tag_id NOT IN (SELECT tag_id FROM categorized);
```
