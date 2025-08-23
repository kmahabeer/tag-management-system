import React, { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { TableCell, TableRow } from "@/components/ui/table";
import { Loader2, Pencil, Trash2 } from "lucide-react";
import { toast } from "sonner";
import { api } from "./api";
import { Tag, UUID } from "./types";
import { mockDeleteTag, mockUpdateTag, useMock } from "./mock";

export default function EditableTagRow({
  tag,
  onChanged,
  onDeleted,
  baseUrl,
}: {
  tag: Tag;
  onChanged: (t: Tag) => void;
  onDeleted: (id: UUID) => void;
  baseUrl?: string;
}) {
  const [editing, setEditing] = useState(false);
  const [name, setName] = useState(tag.name);
  const [display, setDisplay] = useState(tag.display_name ?? "");
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    setName(tag.name);
    setDisplay(tag.display_name ?? "");
  }, [tag]);

  async function save() {
    try {
      setSaving(true);
      const payload: Partial<Tag> = {
        name: name.trim(),
        display_name: display.trim(),
      };
      const updated = useMock
        ? await mockUpdateTag(tag.id, payload)
        : await api<Tag>(`/tags/${tag.id}`, {
            method: "PATCH",
            body: JSON.stringify(payload),
            baseUrl,
          });
      onChanged(updated);
      setEditing(false);
      toast.success("Saved");
    } catch (e: any) {
      toast.error(e.message ?? "Failed to save");
    } finally {
      setSaving(false);
    }
  }

  async function destroy() {
    try {
      if (!confirm(`Delete tag "${tag.display_name || tag.name}"?`)) return;
      if (useMock) await mockDeleteTag(tag.id);
      else await api<void>(`/tags/${tag.id}`, { method: "DELETE", baseUrl });
      onDeleted(tag.id);
      toast.success("Deleted");
    } catch (e: any) {
      toast.error(e.message ?? "Failed to delete");
    }
  }

  return (
    <TableRow>
      <TableCell className="w-[40%]">
        {editing ? (
          <Input value={name} onChange={(e) => setName(e.target.value)} />
        ) : (
          <div className="font-medium">{tag.name}</div>
        )}
      </TableCell>
      <TableCell className="w-[40%]">
        {editing ? (
          <Input value={display} onChange={(e) => setDisplay(e.target.value)} />
        ) : (
          <div>
            {tag.display_name || (
              <span className="text-muted-foreground">—</span>
            )}
          </div>
        )}
      </TableCell>
      <TableCell className="text-right">
        {editing ? (
          <div className="flex justify-end gap-2">
            <Button
              size="sm"
              variant="outline"
              onClick={() => setEditing(false)}
            >
              Cancel
            </Button>
            <Button size="sm" onClick={save} disabled={saving || !name.trim()}>
              {saving && <Loader2 className="h-4 w-4 mr-2 animate-spin" />}Save
            </Button>
          </div>
        ) : (
          <div className="flex justify-end gap-2">
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setEditing(true)}
              aria-label="Edit"
            >
              <Pencil className="h-4 w-4" />
            </Button>
            <Button
              variant="ghost"
              size="icon"
              onClick={destroy}
              aria-label="Delete"
            >
              <Trash2 className="h-4 w-4" />
            </Button>
          </div>
        )}
      </TableCell>
    </TableRow>
  );
}
