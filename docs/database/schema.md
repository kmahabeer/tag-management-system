# Database Schema

This document defines the database schema powering the Tagging Service. It supports:

- Unified tagging across diverse entity types
- Hierarchical and associative relationships between tags
- Composite tag construction (e.g., `"very big red car"`)
- Purpose-driven and context-aware metadata
- UI-configurable layouts and grouping
- Ratings and semantic filtering

PostgreSQL with `uuid-ossp` is required for UUID support as UUIDv4 is used as the primary key type throughout. All tables include audit fields: `created_at`, `updated_at`, `created_by`, `updated_by` and PostgreSQL triggers are employed to automatically update timestamps.

> Each section below introduces a logical grouping of tables or configuration steps that collectively define the tagging system’s database foundation.

## Step 1: Enable required PostgreSQL extensions

PostgreSQL needs this to generate UUIDs:

```sql
-- Enables uuid_generate_v4() function
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## Step 2: Create Shared Lookup Tables

These define **types and classifications** used by other tables.

```sql
-- Types used to classify tags (e.g., Person, Genre)
CREATE TABLE tag_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- Types used to describe relationships between tags (e.g., parent-child)
CREATE TABLE tag_relationship_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
```

## Step 3: Define Core Tag Tables

These store the actual tags and their basic properties.

```sql
-- Canonical tag definitions
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- Synonyms / alternate names for tags
CREATE TABLE tag_aliases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
```

## Step 4: Define Tag Relationship Tables

These tables express many-to-many links between tags, types, and other tags.

```sql
-- tags ↔ tag_types
CREATE TABLE tag_tag_types (
    tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    tag_type_id UUID REFERENCES tag_types(id) ON DELETE CASCADE,
    PRIMARY KEY(tag_id, tag_type_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- tag ↔ tag with relationship type
CREATE TABLE tag_relationships (
    tag_a_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    tag_b_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    relationship_type_id UUID REFERENCES tag_relationship_types(id),
    PRIMARY KEY(tag_a_id, tag_b_id, relationship_type_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- 🆕 tag_compositions: composite tags made of multiple atomic tags
CREATE TABLE tag_compositions (
    composite_tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    component_tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    position INTEGER NOT NULL, -- defines order: 1=very, 2=big, 3=red, 4=car
    PRIMARY KEY(composite_tag_id, component_tag_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
```

## Step 5: Define Entity Tables

These represent the things being tagged, such as images, documents, or files.

```sql
-- Core table for things being tagged
CREATE TABLE entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    location TEXT,
    is_primary BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- entity ↔ purpose_tag (e.g., Film Still, Figure Study)
CREATE TABLE entity_purposes (
    entity_id UUID REFERENCES entities(id) ON DELETE CASCADE,
    purpose_tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    is_primary BOOLEAN DEFAULT FALSE,
    PRIMARY KEY(entity_id, purpose_tag_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- Many-to-one entity versioning
CREATE TABLE entity_versions (
    primary_entity_id UUID REFERENCES entities(id) ON DELETE CASCADE,
    alternate_entity_id UUID REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY(primary_entity_id, alternate_entity_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- entity ↔ tag (e.g., “red”, “Quentin Tarantino”)
CREATE TABLE entity_tags (
    entity_id UUID REFERENCES entities(id) ON DELETE CASCADE,
    tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY(entity_id, tag_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
```

## Step 6: Define UI Configuration Tables

These are optional support tables to guide downstream UI rendering logic.

```sql
-- Group tags visually in the UI (e.g., “Mood”, “Color”)
CREATE TABLE ui_groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);

-- Define tag types used in a UI section, and control types
CREATE TABLE ui_fields (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ui_group_id UUID REFERENCES ui_groups(id),
    tag_type_id UUID REFERENCES tag_types(id),
    control_type TEXT NOT NULL,
    is_required BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
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

Here’s one example — repeat for all tables:

```sql
CREATE TRIGGER trg_set_updated_at_tags
BEFORE UPDATE ON tags
FOR EACH ROW EXECUTE PROCEDURE set_updated_at();
```
