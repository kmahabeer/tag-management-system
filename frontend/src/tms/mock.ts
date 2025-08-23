import { PartOfSpeech, RelationshipType, Tag, TagRelationship, UUID } from './types'

export const useMock = false

const mockDB = {
    pos: [
        { id: crypto.randomUUID(), name: "noun" },
        { id: crypto.randomUUID(), name: "adjective" },
        { id: crypto.randomUUID(), name: "adverb" }
    ] as PartOfSpeech[],
    tags: [
        { id: crypto.randomUUID(), name: "Vehicle", display_name: "Vehicle" },
        { id: crypto.randomUUID(), name: "Car", display_name: "Car" },
        { id: crypto.randomUUID(), name: "Truck", display_name: "Truck" }
    ] as Tag[],
    relTypes: [
        { id: crypto.randomUUID(), name: "is a" },
        { id: crypto.randomUUID(), name: "is associated with" },
        { id: crypto.randomUUID(), name: "parent" }
    ] as RelationshipType[],
    tagRels: [] as TagRelationship[]
}

export async function mockListTags(): Promise<Tag[]> {
    return JSON.parse(JSON.stringify(mockDB.tags))
}
export async function mockCreateTag(payload: Partial<Tag>): Promise<Tag> {
    const t: Tag = {
        id: crypto.randomUUID(),
        name: payload.name ?? "",
        display_name: payload.display_name ?? payload.name ?? "",
        part_of_speech_id: payload.part_of_speech_id ?? null,
        metadata: payload.metadata ?? null
    }
    mockDB.tags.unshift(t)
    return t
}
export async function mockUpdateTag(id: UUID, payload: Partial<Tag>): Promise<Tag> {
    const idx = mockDB.tags.findIndex((t) => t.id === id)
    if (idx === -1) throw new Error("Not found")
    mockDB.tags[idx] = { ...mockDB.tags[idx], ...payload }
    return mockDB.tags[idx]
}
export async function mockDeleteTag(id: UUID): Promise<void> {
    mockDB.tags = mockDB.tags.filter((t) => t.id !== id)
}
export async function mockListPos(): Promise<PartOfSpeech[]> {
    return mockDB.pos
}
export async function mockListRelTypes(): Promise<RelationshipType[]> {
    return mockDB.relTypes
}
export async function mockListTagRels(): Promise<TagRelationship[]> {
    return mockDB.tagRels.slice()
}
export async function mockCreateTagRel(payload: Omit<TagRelationship, "id">): Promise<TagRelationship> {
    const r = { id: crypto.randomUUID(), ...payload }
    mockDB.tagRels.unshift(r)
    return r
}
export async function mockDeleteTagRel(id: UUID): Promise<void> {
    mockDB.tagRels = mockDB.tagRels.filter((r) => r.id !== id)
}
