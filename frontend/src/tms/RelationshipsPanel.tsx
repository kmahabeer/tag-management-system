import React, { useEffect, useMemo, useState } from "react";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import Toolbar from "./Toolbar";
import RelationshipCreator from "./RelationshipCreator";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Table,
  TableBody,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Loader2, Trash2, Tags as TagsIcon } from "lucide-react";
import { toast } from "sonner";
import { api } from "./api";
import { RelationshipType, Tag, TagRelationship, UUID } from "./types";
import {
  mockDeleteTagRel,
  mockListRelTypes,
  mockListTagRels,
  mockListTags,
  useMock,
} from "./mock";

export default function RelationshipsPanel({ baseUrl }: { baseUrl?: string }) {
  const [tags, setTags] = useState<Tag[]>([]);
  const [relTypes, setRelTypes] = useState<RelationshipType[]>([]);
  const [rels, setRels] = useState<TagRelationship[]>([]);
  const [loading, setLoading] = useState(true);

  async function load() {
    try {
      setLoading(true);
      const [t, rt, r] = await Promise.all([
        useMock ? mockListTags() : api<Tag[]>("/tags", { baseUrl }),
        useMock
          ? mockListRelTypes()
          : api<RelationshipType[]>("/tag_relationship_types", { baseUrl }),
        useMock
          ? mockListTagRels()
          : api<TagRelationship[]>("/tag_relationships", { baseUrl }),
      ]);
      setTags(t);
      setRelTypes(rt);
      setRels(r);
    } catch (e: any) {
      toast.error(e.message ?? "Failed to load relationships");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    load();
  }, [baseUrl]);

  const tagById = useMemo(
    () => Object.fromEntries(tags.map((t) => [t.id, t])),
    [tags]
  );
  const relTypeById = useMemo(
    () => Object.fromEntries(relTypes.map((rt) => [rt.id, rt])),
    [relTypes]
  );

  async function destroy(id: UUID) {
    try {
      if (!confirm("Delete this relationship?")) return;
      if (useMock) await mockDeleteTagRel(id);
      else
        await api<void>(`/tag_relationships/${id}`, {
          method: "DELETE",
          baseUrl,
        });
      setRels((prev) => prev.filter((r) => r.id !== id));
      toast.success("Relationship deleted");
    } catch (e: any) {
      toast.error(e.message ?? "Failed to delete");
    }
  }

  return (
    <Card className="shadow-sm">
      <CardHeader className="pb-3">
        <Toolbar
          title="Tag Relationships"
          subtitle="Create and manage relationships between tags"
          onRefresh={load}
        />
      </CardHeader>
      <CardContent className="space-y-6">
        <RelationshipCreator
          baseUrl={baseUrl}
          tags={tags}
          relTypes={relTypes}
          onCreated={(r) => setRels((prev) => [r, ...prev])}
        />
        <ScrollArea className="max-h-[360px] rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Tag A</TableHead>
                <TableHead>Type</TableHead>
                <TableHead>Tag B</TableHead>
                <TableHead>Description</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {loading ? (
                <TableRow>
                  <td
                    colSpan={5}
                    className="text-center py-8 text-muted-foreground"
                  >
                    <Loader2 className="h-4 w-4 inline mr-2 animate-spin" />{" "}
                    Loading…
                  </td>
                </TableRow>
              ) : rels.length === 0 ? (
                <TableRow>
                  <td
                    colSpan={5}
                    className="text-center py-10 text-muted-foreground"
                  >
                    No relationships yet.
                  </td>
                </TableRow>
              ) : (
                rels.map((r) => (
                  <TableRow key={r.id}>
                    <td>
                      <Badge variant="secondary">
                        <TagsIcon className="h-3 w-3 mr-1" />
                        {tagById[r.tag_a_id]?.display_name ||
                          tagById[r.tag_a_id]?.name ||
                          r.tag_a_id}
                      </Badge>
                    </td>
                    <td>
                      {relTypeById[r.relationship_type_id]?.name ||
                        r.relationship_type_id}
                    </td>
                    <td>
                      <Badge variant="outline">
                        {tagById[r.tag_b_id]?.display_name ||
                          tagById[r.tag_b_id]?.name ||
                          r.tag_b_id}
                      </Badge>
                    </td>
                    <td
                      className="max-w-[360px] truncate"
                      title={r.description || undefined}
                    >
                      {r.description || (
                        <span className="text-muted-foreground">—</span>
                      )}
                    </td>
                    <td className="text-right">
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => destroy(r.id)}
                        aria-label="Delete"
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </td>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </ScrollArea>
      </CardContent>
    </Card>
  );
}
