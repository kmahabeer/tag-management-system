import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { api } from "./api";
import { RelationshipType, Tag, TagRelationship } from "./types";
import { mockCreateTagRel, useMock } from "./mock";

export default function RelationshipCreator({
  tags,
  relTypes,
  onCreated,
  baseUrl,
}: {
  tags: Tag[];
  relTypes: RelationshipType[];
  onCreated: (r: TagRelationship) => void;
  baseUrl?: string;
}) {
  const [a, setA] = useState<string>("");
  const [b, setB] = useState<string>("");
  const [typeId, setTypeId] = useState<string>("");
  const [desc, setDesc] = useState<string>("");
  const [saving, setSaving] = useState(false);

  const valid = a && b && typeId && a !== b;

  async function createRel() {
    try {
      setSaving(true);
      const payload = {
        tag_a_id: a,
        tag_b_id: b,
        relationship_type_id: typeId,
        description: desc || undefined,
      } as Omit<TagRelationship, "id">;
      const created = useMock
        ? await mockCreateTagRel(payload)
        : await api<TagRelationship>("/tag_relationships", {
            method: "POST",
            body: JSON.stringify(payload),
            baseUrl,
          });
      onCreated(created);
      setA("");
      setB("");
      setTypeId("");
      setDesc("");
      toast.success("Relationship created");
    } catch (e: any) {
      toast.error(e.message ?? "Failed to create relationship");
    } finally {
      setSaving(false);
    }
  }

  return (
    <div className="grid gap-3 md:grid-cols-2 lg:grid-cols-4">
      <div className="grid gap-2">
        <Label>Tag A</Label>
        <Select value={a} onValueChange={setA}>
          <SelectTrigger>
            <SelectValue placeholder="Select tag A" />
          </SelectTrigger>
          <SelectContent>
            {tags.map((t) => (
              <SelectItem key={t.id} value={t.id}>
                {t.display_name || t.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
      <div className="grid gap-2">
        <Label>Relationship Type</Label>
        <Select value={typeId} onValueChange={setTypeId}>
          <SelectTrigger>
            <SelectValue placeholder="Select type" />
          </SelectTrigger>
          <SelectContent>
            {relTypes.map((rt) => (
              <SelectItem key={rt.id} value={rt.id}>
                {rt.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
      <div className="grid gap-2">
        <Label>Tag B</Label>
        <Select value={b} onValueChange={setB}>
          <SelectTrigger>
            <SelectValue placeholder="Select tag B" />
          </SelectTrigger>
          <SelectContent>
            {tags.map((t) => (
              <SelectItem key={t.id} value={t.id}>
                {t.display_name || t.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
      <div className="grid gap-2 lg:col-span-1 md:col-span-2">
        <Label>Description</Label>
        <Textarea
          rows={1}
          value={desc}
          onChange={(e) => setDesc(e.target.value)}
          placeholder={'e.g., "Vehicle is a parent of Car"'}
        />
      </div>
      <div className="md:col-span-2 lg:col-span-4 flex justify-end">
        <Button onClick={createRel} disabled={!valid || saving}>
          {saving && <Loader2 className="h-4 w-4 mr-2 animate-spin" />}Create
          relationship
        </Button>
      </div>
    </div>
  );
}
