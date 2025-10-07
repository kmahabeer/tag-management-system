---
title: Database Scripts
parent: Database
nav_order: 2
---

# Database Schema Generation

This document provides the complete **SQL build scripts** required to generate the Tag Management System database from scratch. It is intended for initializing a clean PostgreSQL instance and establishing all tables, relationships, and triggers that power the Tag Management System.

These scripts:

- Create all required lookup, tagging, and entity tables
- Establish relationships, constraints, and foreign keys
- Enable audit fields and automatic timestamp updates
- Enforce semantic integrity and relationship uniqueness
- Provide a reusable, migration-friendly foundation for development and testing

PostgreSQL (version ≥ 13) is required, with the `uuid-ossp` extension enabled for UUIDv4 primary keys. All tables include standard audit columns (`created_at`, `updated_at`, `created_by`, `updated_by`), and triggers automatically maintain timestamp consistency.

> The steps below should be executed sequentially in a fresh database environment to fully reproduce the Tagging Service schema.

## Step 1: Enable required PostgreSQL extensions

PostgreSQL needs this to generate UUIDs:

```sql
-- Enables uuid_generate_v4() function
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## Step 2: Create Shared Lookup Tables

These define **types and classifications** used by other tables.

### Step 2a: `parts_of_speech`

```sql
CREATE TABLE parts_of_speech (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 2b: `contexts`

```sql
CREATE TABLE contexts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    classification_type TEXT CHECK (classification_type IN ('subjective', 'objective')),
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 2c: `tag_relationship_types`

```sql
CREATE TABLE tag_relationship_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 2d: `entity_relationship_types`

```sql
CREATE TABLE entity_relationship_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 2e: `rating_types`

```sql
CREATE TABLE rating_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    is_normalized BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 2f: `ratings`

```sql
CREATE TABLE ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    score INT NOT NULL,
    description TEXT,
    rating_type_id UUID REFERENCES rating_types(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Step 3: Define Core Tag Tables

### Step 3a: `tags`

```sql
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    display_name TEXT,
    metadata JSONB,
    part_of_speech_id UUID REFERENCES parts_of_speech(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 3b: `tag_aliases`

```sql
CREATE TABLE tag_aliases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 3c: `tag_relationships`

```sql
CREATE TABLE tag_relationships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_a_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    tag_b_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    relationship_type_id UUID NOT NULL REFERENCES tag_relationship_types(id),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 3d: `tag_compositions`

```sql
CREATE TABLE tag_compositions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    base_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    component_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    position INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(base_tag_id, component_tag_id, position)
);
```

### Step 3e: `tag_context_ratings`

```sql
CREATE TABLE tag_context_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    rating_id UUID NOT NULL REFERENCES ratings(id),
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 3f: `tag_relationship_ratings`

```sql
CREATE TABLE tag_relationship_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_a_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    tag_b_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    rating_id UUID NOT NULL REFERENCES ratings(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Step 5: Define Entity Tables

### Step 4a: `entities`

```sql
CREATE TABLE entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    location TEXT,
    is_primary BOOLEAN DEFAULT TRUE,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 4b: `entity_purposes`

```sql
CREATE TABLE entity_purposes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    purpose_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 4c: `entity_relationships`

```sql
CREATE TABLE entity_relationships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_a_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    entity_b_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    relationship_type_id UUID NOT NULL REFERENCES entity_relationship_types(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 4d: `entity_tags`

```sql
CREATE TABLE entity_tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Step 6: Define UI Configuration Tables

### Step 5a: `ui_layouts`

```sql
CREATE TABLE ui_layouts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    purpose_tag_id UUID REFERENCES tags(id),
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 5b: `ui_groups`

```sql
CREATE TABLE ui_groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Step 5c: `ui_fields`

```sql
CREATE TABLE ui_fields (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ui_layout_id UUID NOT NULL REFERENCES ui_layouts(id) ON DELETE CASCADE,
    ui_group_id UUID NOT NULL REFERENCES ui_groups(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    category_tag_id UUID NOT NULL REFERENCES tags(id),
    sort_order INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Step 6: Create Table Relationships and Constraints

These are additional **constraints and rules** applied across tables to enforce semantic correctness, data consistency, and uniqueness.

### Step 6a: Enforce One Primary Purpose per Entity

```sql
-- Only one primary purpose per entity
CREATE UNIQUE INDEX uniq_primary_purpose_per_entity
ON entity_purposes(entity_id)
WHERE is_primary = TRUE;
```

### Step 6b: Enforce Tag Composition Uniqueness

```sql
-- Prevent duplicate composite structures
CREATE UNIQUE INDEX uniq_tag_composition_structure
ON tag_compositions(base_tag_id, component_tag_id, position);
```

### Step 6c: Prevent Duplicate Tag Relationships

```sql
-- Prevent duplicate relationships between the same tags and type
CREATE UNIQUE INDEX uniq_tag_relationship
ON tag_relationships(tag_a_id, tag_b_id, relationship_type_id);
```

### Step 6d: Prevent Circular Entity Relationships

You may want to enforce this with a function/trigger later (e.g., `entity_a_id != entity_b_id`).

```sql
-- Prevent self-linking
ALTER TABLE entity_relationships
ADD CONSTRAINT no_self_referencing_entities
CHECK (entity_a_id <> entity_b_id);
```

### Step 6e: Optional: Rating Normalization (if needed)

Example constraint if you later enforce normalization on rating scores (optional):

```sql
-- Example CHECK constraint for normalized scores
-- Only apply this if you're enforcing it through rating_types
-- ALTER TABLE ratings ADD CHECK (
--     (is_normalized = TRUE AND score BETWEEN 0 AND 1) OR
--     (is_normalized = FALSE AND score BETWEEN 1 AND 10)
-- );
```

## Step 7: Create Timestamp Trigger

To automatically keep `updated_at` fresh:

```sql
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

## Step 8: Attach Triggers to All Tables

> Applies the trigger to update `updated_at` before each row update. Run this once per table that contains an `updated_at` field.

Here are trigger statements for all tables:

```sql
-- === Shared Tables ===
CREATE TRIGGER trg_set_updated_at_parts_of_speech
BEFORE UPDATE ON parts_of_speech
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_contexts
BEFORE UPDATE ON contexts
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_relationship_types
BEFORE UPDATE ON tag_relationship_types
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_entity_relationship_types
BEFORE UPDATE ON entity_relationship_types
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_rating_types
BEFORE UPDATE ON rating_types
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_ratings
BEFORE UPDATE ON ratings
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- === Tags ===
CREATE TRIGGER trg_set_updated_at_tags
BEFORE UPDATE ON tags
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_aliases
BEFORE UPDATE ON tag_aliases
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_relationships
BEFORE UPDATE ON tag_relationships
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_compositions
BEFORE UPDATE ON tag_compositions
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_context_ratings
BEFORE UPDATE ON tag_context_ratings
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_tag_relationship_ratings
BEFORE UPDATE ON tag_relationship_ratings
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- === Entities ===
CREATE TRIGGER trg_set_updated_at_entities
BEFORE UPDATE ON entities
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_entity_purposes
BEFORE UPDATE ON entity_purposes
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_entity_relationships
BEFORE UPDATE ON entity_relationships
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_entity_tags
BEFORE UPDATE ON entity_tags
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- === UI Configuration ===
CREATE TRIGGER trg_set_updated_at_ui_layouts
BEFORE UPDATE ON ui_layouts
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_ui_groups
BEFORE UPDATE ON ui_groups
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_set_updated_at_ui_fields
BEFORE UPDATE ON ui_fields
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```
