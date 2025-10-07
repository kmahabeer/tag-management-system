---
title: User Interface (UI) Configurations
parent: Database Schema
has_children: false
nav_order: 4
---
# User Interface (UI) Configurations

The system includes a small set of tables used by front-end clients to define how tag data should be displayed. These are **metadata-only** and do not render UI elements directly.

## `ui_layouts` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `purpose_tag_id` | UUID | Foreign key to `tags`; determines which purpose this layout applies to. |
| `name` | TEXT | Optional display name for debugging or UI reference. |

## `ui_groups` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Label shown in the UI (e.g., “Photographer,” “Camera Metadata”). |

## `ui_fields` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `ui_layout_id` | UUID | Foreign key to [`ui_layouts`](#ui_layouts-table). |
| `ui_group_id` | UUID | Foreign key to [`ui_groups`](#ui_groups-table). |
| `context_id` | UUID | Foreign key to [`contexts`](../utilities/index.md#contexts). |
| `category_tag_id` | UUID | Foreign key to [`tags`](../tags/index.md). |
| `sort_order` | INT | Determines display order. |

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **UI Layouts** | `/ui/layouts`, `/ui/layouts/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage top-level UI layout profiles defining grouping structure. |
| **UI Groups** | `/ui/groups`, `/ui/groups/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage reusable labeled sections within layouts. |
| **UI Fields** | `/ui/fields`, `/ui/fields/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage mappings between tags, contexts, and layout sections. |

### Coverage Summary

| UI Component | CRUD | Description |
|---------------|------|--------------|
| Layouts | ✅ Full | Define the top-level structure of tag groupings in the UI. |
| Groups | ✅ Full | Define reusable labeled sections used inside layouts. |
| Fields | ✅ Full | Bind tags, contexts, and categories to specific layout sections. |
