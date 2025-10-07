---
title: Home
---
# Tag Management System Documentation

This section contains the documentation for the Tag Management System.

The **Tag Management System** is designed to power a flexible and context-aware user interface (UI) for entities like images, film stills, documents, and more. Each entity may serve different "purposes" (e.g. a film still, a figure study, a reference photo) and should dynamically display its associated metadata (tags) in a structured way according to those purposes.

At the heart of the system is a flexible tagging model with many-to-many relationships, hierarchical categorization, and purpose-specific rendering logic. Tags can represent people, roles, vehicles, genres, purposes, and any other useful classification. These tags can be reused and linked across entities like films or subjects.

Frontend clients, such as a web-based graphical user interface (GUI) will read structured JSON according to predefined layouts for each purpose (e.g., "Film Still" might expect to show Directors, Writers, Actors, etc.). Tags that do not fit into this predefined structure are shown in a catch-all "Other Tags" section.

## Features

- Tag entities with hierarchical, typed, and composable labels
- Group tags by UI logic using types and display controls
- Store purposes for each entity to customize rendering
- Link tags to multiple entities with support for aliases and relationships
- Supports modifier phrases like “very tall” or “bright red”

## Usage

The system is designed to expose a RESTful API for creating, updating, and querying tags, as well as managing tag relationships.

## Contents

- [Overview](./overview/_index.md)
	- [Concepts](./overview/concepts.md)
	- [Goals](./overview/goals.md)
- [Database](./database/_index.md)
	- [Entities](./database/entities.md)
	- [Tags](./database/tags.md)
	- [Entity Tagging](./database/entity_tagging.md)
	- [Utilities](ui_configurations.md)
	- [Database Schema](./database/schema/schema.md)
- [API](./api/_index.md)
	- [API Endpoints](./api/endpoints/_index.md)
	- [API Resources](./api/resources/_index.md)
- [Logic](./logic/_index.md)
	- [Composite Tagging](./logic/composite_tagging.md)
	- [Labeled Data Import](./logic/labeled_data_import.md)
- [User Interface (UI)](./ui/_index.md)
	- [UI Rendering Logic](./ui/rendering.md)
	- [Structured JSON Output](./ui/json_examples.md)
- [Examples](./examples/_index.md)
	- [Data](./examples/data/_index.md)
		- [Entities](./examples/data/example_entities.md)
		- [Tags](./examples/data/example_tags.md)
		- [Tag Assignments](./examples/data/example_entity_tags.md)
	- [Queries](./examples/queries/_index.md)
		- [Tag Queries](./examples/queries/tag_queries.md)
		- [Entity Queries](./examples/queries/entity_queries.md)
		- [Other Query Examples](./examples/queries/other_queries.md)
- [Glossary](./glossary.md)
