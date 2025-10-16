-- ================================================================
-- TAG MANAGEMENT SYSTEM - SEED DATA
-- ================================================================

-- === RATING TYPES ===
INSERT INTO rating_types (name, is_normalized) VALUES
('Star Rating', FALSE),
('Normalized Score', TRUE);

-- === RATINGS ===
INSERT INTO ratings (name, score, description, rating_type_id)
SELECT 'One Star', 1, 'Lowest rating', id FROM rating_types WHERE name = 'Star Rating'
UNION ALL
SELECT 'Two Stars', 2, 'Below average', id FROM rating_types WHERE name = 'Star Rating'
UNION ALL
SELECT 'Three Stars', 3, 'Average', id FROM rating_types WHERE name = 'Star Rating'
UNION ALL
SELECT 'Four Stars', 4, 'Good', id FROM rating_types WHERE name = 'Star Rating'
UNION ALL
SELECT 'Five Stars', 5, 'Excellent', id FROM rating_types WHERE name = 'Star Rating';

-- === PARTS OF SPEECH ===
INSERT INTO parts_of_speech (name, description) VALUES
('noun', 'A person, place, thing, or idea'),
('verb', 'An action or state of being'),
('adjective', 'A word that modifies a noun');

-- === CONTEXTS ===
INSERT INTO contexts (name, classification_type, description) VALUES
('artistic', 'subjective', 'Context of artistic or creative use'),
('technical', 'objective', 'Context of technical or scientific use'),
('personal', 'subjective', 'Personal or emotional context');

-- === TAG RELATIONSHIP TYPES ===
INSERT INTO tag_relationship_types (name) VALUES
('synonym'),
('antonym'),
('related'),
('derivative');

-- === ENTITY RELATIONSHIP TYPES ===
INSERT INTO entity_relationship_types (name) VALUES
('depends_on'),
('derived_from'),
('associated_with');

-- === TAGS ===
INSERT INTO tags (name, display_name, metadata, part_of_speech_id)
SELECT 'creativity', 'Creativity', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'noun'
UNION ALL
SELECT 'analyze', 'Analyze', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'verb'
UNION ALL
SELECT 'innovative', 'Innovative', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'adjective';

-- === TAG RELATIONSHIPS ===
INSERT INTO tag_relationships (tag_a_id, tag_b_id, relationship_type_id, description)
SELECT t1.id, t2.id, tr.id, 'related concept'
FROM tags t1, tags t2, tag_relationship_types tr
WHERE t1.name = 'creativity' AND t2.name = 'innovative' AND tr.name = 'related';

-- === ENTITIES ===
INSERT INTO entities (name, location, metadata) VALUES
('Entity A', '/data/entity_a', '{}'::jsonb),
('Entity B', '/data/entity_b', '{}'::jsonb);

-- === ENTITY PURPOSES ===
INSERT INTO entity_purposes (entity_id, purpose_tag_id, is_primary)
SELECT e.id, t.id, TRUE
FROM entities e, tags t
WHERE e.name = 'Entity A' AND t.name = 'creativity';

-- === ENTITY TAGS ===
INSERT INTO entity_tags (entity_id, tag_id, context_id, metadata)
SELECT e.id, t.id, c.id, '{}'::jsonb
FROM entities e, tags t, contexts c
WHERE e.name = 'Entity A' AND t.name = 'innovative' AND c.name = 'artistic';

-- === UI GROUPS ===
INSERT INTO ui_groups (name) VALUES
('General'),
('Advanced');

-- === UI LAYOUTS ===
INSERT INTO ui_layouts (name) VALUES
('Default Layout');

-- === UI FIELDS ===
INSERT INTO ui_fields (ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order)
SELECT ul.id, ug.id, c.id, t.id, 1
FROM ui_layouts ul, ui_groups ug, contexts c, tags t
WHERE ul.name = 'Default Layout' AND ug.name = 'General' AND c.name = 'artistic' AND t.name = 'creativity';
