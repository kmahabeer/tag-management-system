export type UUID = string

export type Tag = {
    id: UUID
    name: string
    display_name?: string | null
    metadata?: Record<string, unknown> | null
    part_of_speech_id?: UUID | null
    created_at?: string
    updated_at?: string
}

export type PartOfSpeech = {
    id: UUID
    name: string
}

export type RelationshipType = {
    id: UUID
    name: string
}

export type TagRelationship = {
    id: UUID
    tag_a_id: UUID
    tag_b_id: UUID
    relationship_type_id: UUID
    description?: string | null
}
