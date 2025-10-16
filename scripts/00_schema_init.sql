-- ================================================================
-- TAG MANAGEMENT SYSTEM - DATABASE SCHEMA INITIALIZATION SCRIPT
-- ================================================================
-- Requires: PostgreSQL ≥ 13
-- Run in a fresh database before loading seed data.
-- ================================================================

-- === STEP 1: EXTENSIONS ===
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- === STEP 2: SHARED LOOKUP TABLES ===
CREATE TABLE parts_of_speech (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contexts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    classification_type TEXT CHECK (classification_type IN ('subjective', 'objective')),
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_relationship_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entity_relationship_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rating_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    is_normalized BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    score INT NOT NULL,
    description TEXT,
    rating_type_id UUID REFERENCES rating_types(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- === STEP 3: CORE TAG TABLES ===
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    display_name TEXT,
    metadata JSONB,
    part_of_speech_id UUID REFERENCES parts_of_speech(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_aliases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_relationships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_a_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    tag_b_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    relationship_type_id UUID NOT NULL REFERENCES tag_relationship_types(id),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_compositions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    base_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    component_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    position INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(base_tag_id, component_tag_id, position)
);

CREATE TABLE tag_context_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    rating_id UUID NOT NULL REFERENCES ratings(id),
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_relationship_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag_a_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    tag_b_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    rating_id UUID NOT NULL REFERENCES ratings(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- === STEP 4: ENTITY TABLES ===
CREATE TABLE entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    location TEXT,
    is_primary BOOLEAN DEFAULT TRUE,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entity_purposes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    purpose_tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entity_relationships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_a_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    entity_b_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    relationship_type_id UUID NOT NULL REFERENCES entity_relationship_types(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT no_self_referencing_entities CHECK (entity_a_id <> entity_b_id)
);

CREATE TABLE entity_relationship_ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_a_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    entity_b_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    rating_id UUID NOT NULL REFERENCES ratings(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entity_tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_id UUID NOT NULL REFERENCES entities(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    context_id UUID NOT NULL REFERENCES contexts(id),
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- === STEP 5: UI CONFIGURATION ===
CREATE TABLE ui_layouts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    purpose_tag_id UUID REFERENCES tags(id),
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ui_groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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

-- === STEP 6: INDEXES & CONSTRAINTS ===
CREATE UNIQUE INDEX uniq_primary_purpose_per_entity
ON entity_purposes(entity_id)
WHERE is_primary = TRUE;

CREATE UNIQUE INDEX uniq_tag_composition_structure
ON tag_compositions(base_tag_id, component_tag_id, position);

CREATE UNIQUE INDEX uniq_tag_relationship
ON tag_relationships(tag_a_id, tag_b_id, relationship_type_id);

-- === STEP 7: TRIGGERS ===
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Attach to all tables that include updated_at
DO $$
DECLARE
  tbl RECORD;
BEGIN
  FOR tbl IN
    SELECT table_name
    FROM information_schema.columns
    WHERE column_name = 'updated_at'
      AND table_schema = 'public'
  LOOP
    EXECUTE format('
      CREATE TRIGGER trg_set_updated_at_%I
      BEFORE UPDATE ON %I
      FOR EACH ROW
      EXECUTE FUNCTION set_updated_at();
    ', tbl.table_name, tbl.table_name);
  END LOOP;
END$$;
