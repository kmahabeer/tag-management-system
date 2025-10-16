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
('adjective', 'A word that modifies a noun'),
('adverb', 'A word that modifies a verb or adjective'),
('preposition', 'A word that shows relationship between nouns or pronouns');

-- === CONTEXTS ===
INSERT INTO contexts (name, classification_type, description) VALUES
('artistic', 'subjective', 'Context of artistic or creative use'),
('technical', 'objective', 'Context of technical or scientific use'),
('personal', 'subjective', 'Personal or emotional context'),
('academic', 'objective', 'Used in educational or research environments'),
('marketing', 'subjective', 'Context related to promotion, branding, or audience engagement');

-- === TAG RELATIONSHIP TYPES ===
INSERT INTO tag_relationship_types (name) VALUES
('synonym'),
('antonym'),
('related'),
('derivative'),
('subset'),
('superset');

-- === ENTITY RELATIONSHIP TYPES ===
INSERT INTO entity_relationship_types (name) VALUES
('depends_on'),
('derived_from'),
('associated_with'),
('inspired_by'),
('references');

-- === TAGS ===
INSERT INTO tags (name, display_name, metadata, part_of_speech_id)
SELECT 'creativity', 'Creativity', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'noun'
UNION ALL
SELECT 'innovation', 'Innovation', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'noun'
UNION ALL
SELECT 'design', 'Design', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'noun'
UNION ALL
SELECT 'analyze', 'Analyze', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'verb'
UNION ALL
SELECT 'develop', 'Develop', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'verb'
UNION ALL
SELECT 'imagine', 'Imagine', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'verb'
UNION ALL
SELECT 'innovative', 'Innovative', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'adjective'
UNION ALL
SELECT 'efficient', 'Efficient', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'adjective'
UNION ALL
SELECT 'precisely', 'Precisely', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'adverb'
UNION ALL
SELECT 'boldly', 'Boldly', '{}'::jsonb, id FROM parts_of_speech WHERE name = 'adverb';

-- === TAG RELATIONSHIPS ===
INSERT INTO tag_relationships (tag_a_id, tag_b_id, relationship_type_id, description)
SELECT t1.id, t2.id, tr.id, 'innovation as an expression of creativity'
FROM tags t1, tags t2, tag_relationship_types tr
WHERE t1.name = 'creativity' AND t2.name = 'innovation' AND tr.name = 'related';

INSERT INTO tag_relationships (tag_a_id, tag_b_id, relationship_type_id, description)
SELECT t1.id, t2.id, tr.id, 'design is a subset of creativity'
FROM tags t1, tags t2, tag_relationship_types tr
WHERE t1.name = 'design' AND t2.name = 'creativity' AND tr.name = 'subset';

INSERT INTO tag_relationships (tag_a_id, tag_b_id, relationship_type_id, description)
SELECT t1.id, t2.id, tr.id, 'innovation builds upon analysis'
FROM tags t1, tags t2, tag_relationship_types tr
WHERE t1.name = 'innovation' AND t2.name = 'analyze' AND tr.name = 'derived';

-- === TAG CONTEXT RATINGS ===
INSERT INTO tag_context_ratings (tag_id, context_id, rating_id)
SELECT t.id, c.id, r.id
FROM tags t, contexts c, ratings r
WHERE t.name = 'innovation' AND c.name = 'technical' AND r.name = 'Five Stars';

INSERT INTO tag_context_ratings (tag_id, context_id, rating_id)
SELECT t.id, c.id, r.id
FROM tags t, contexts c, ratings r
WHERE t.name = 'creativity' AND c.name = 'artistic' AND r.name = 'Five Stars';

-- === ENTITIES ===
INSERT INTO entities (name, location, metadata) VALUES
('Entity A', '/data/entity_a', '{"type":"dataset","source":"internal"}'),
('Entity B', '/data/entity_b', '{"type":"image","source":"external"}'),
('Entity C', '/data/entity_c', '{"type":"document","source":"research"}'),
('Entity D', '/data/entity_d', '{"type":"video","source":"archive"}');

-- === ENTITY PURPOSES ===
INSERT INTO entity_purposes (entity_id, purpose_tag_id, is_primary)
SELECT e.id, t.id, TRUE
FROM entities e, tags t
WHERE e.name = 'Entity A' AND t.name = 'creativity';

INSERT INTO entity_purposes (entity_id, purpose_tag_id, is_primary)
SELECT e.id, t.id, TRUE
FROM entities e, tags t
WHERE e.name = 'Entity B' AND t.name = 'design';

INSERT INTO entity_purposes (entity_id, purpose_tag_id, is_primary)
SELECT e.id, t.id, FALSE
FROM entities e, tags t
WHERE e.name = 'Entity C' AND t.name = 'analyze';

-- === ENTITY RELATIONSHIPS ===
INSERT INTO entity_relationships (entity_a_id, entity_b_id, relationship_type_id)
SELECT ea.id, eb.id, er.id
FROM entities ea, entities eb, entity_relationship_types er
WHERE ea.name = 'Entity A' AND eb.name = 'Entity B' AND er.name = 'associated_with';

INSERT INTO entity_relationships (entity_a_id, entity_b_id, relationship_type_id)
SELECT ea.id, eb.id, er.id
FROM entities ea, entities eb, entity_relationship_types er
WHERE ea.name = 'Entity B' AND eb.name = 'Entity C' AND er.name = 'derived_from';

-- === ENTITY TAGS ===
INSERT INTO entity_tags (entity_id, tag_id, context_id, metadata)
SELECT e.id, t.id, c.id, '{"confidence":0.95}'::jsonb
FROM entities e, tags t, contexts c
WHERE e.name = 'Entity A' AND t.name = 'innovative' AND c.name = 'artistic';

INSERT INTO entity_tags (entity_id, tag_id, context_id, metadata)
SELECT e.id, t.id, c.id, '{"confidence":0.89}'::jsonb
FROM entities e, tags t, contexts c
WHERE e.name = 'Entity B' AND t.name = 'efficient' AND c.name = 'technical';

INSERT INTO entity_tags (entity_id, tag_id, context_id, metadata)
SELECT e.id, t.id, c.id, '{"confidence":0.82}'::jsonb
FROM entities e, tags t, contexts c
WHERE e.name = 'Entity C' AND t.name = 'creativity' AND c.name = 'academic';

-- === ENTITY RELATIONSHIP RATINGS ===
INSERT INTO entity_relationship_ratings (entity_a_id, entity_b_id, context_id, rating_id)
SELECT ea.id, eb.id, c.id, r.id
FROM entities ea, entities eb, contexts c, ratings r
WHERE ea.name = 'Entity A' AND eb.name = 'Entity B' AND c.name = 'artistic' AND r.name = 'Four Stars';

-- === UI GROUPS ===
INSERT INTO ui_groups (name) VALUES
('General'),
('Advanced'),
('Metadata');

-- === UI LAYOUTS ===
INSERT INTO ui_layouts (name, purpose_tag_id)
SELECT 'Default Layout', t.id FROM tags t WHERE t.name = 'creativity'
UNION ALL
SELECT 'Technical Layout', t.id FROM tags t WHERE t.name = 'innovation';

-- === UI FIELDS ===
INSERT INTO ui_fields (ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order)
SELECT ul.id, ug.id, c.id, t.id, 1
FROM ui_layouts ul, ui_groups ug, contexts c, tags t
WHERE ul.name = 'Default Layout' AND ug.name = 'General' AND c.name = 'artistic' AND t.name = 'creativity';

INSERT INTO ui_fields (ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order)
SELECT ul.id, ug.id, c.id, t.id, 2
FROM ui_layouts ul, ui_groups ug, contexts c, tags t
WHERE ul.name = 'Technical Layout' AND ug.name = 'Advanced' AND c.name = 'technical' AND t.name = 'innovation';
