import React, { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { Loader2, Plus } from "lucide-react";
import { toast } from "sonner";
import { api } from "./api";
import { PartOfSpeech, Tag } from "./types";
import { mockCreateTag, mockListPos, useMock } from "./mock";

export default function CreateTagDialog({
  baseUrl,
  onCreated,
}: {
  baseUrl?: string;
  onCreated: (t: Tag) => void;
}) {
  const [open, setOpen] = useState(false);
  const [name, setName] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [posList, setPosList] = useState<PartOfSpeech[]>([]);
  const [posId, setPosId] = useState<string | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      try {
        const pos = useMock
          ? await mockListPos()
          : await api<PartOfSpeech[]>("/parts_of_speech", { baseUrl });
        setPosList(pos);
      } catch {}
    })();
  }, [baseUrl]);

  async function handleCreate() {
    try {
      setLoading(true);
      const payload: Partial<Tag> = {
        name: name.trim(),
        display_name: displayName.trim() || name.trim(),
        part_of_speech_id: posId ?? null,
      };
      const created = useMock
        ? await mockCreateTag(payload)
        : await api<Tag>("/tags", {
            method: "POST",
            body: JSON.stringify(payload),
            baseUrl,
          });
      onCreated(created);
      setOpen(false);
      setName("");
      setDisplayName("");
      setPosId(undefined);
      toast.success("Tag created");
    } catch (e: any) {
      toast.error(e.message ?? "Failed to create tag");
    } finally {
      setLoading(false);
    }
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button size="sm">
          <Plus className="h-4 w-4 mr-2" />
          New Tag
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[500px]">
        <DialogHeader>
          <DialogTitle>Create Tag</DialogTitle>
          <DialogDescription>Define a new tag.</DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-2">
          <div className="grid gap-2">
            <Label>Name</Label>
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="e.g., Vehicle"
            />
          </div>
          <div className="grid gap-2">
            <Label>Display name</Label>
            <Input
              value={displayName}
              onChange={(e) => setDisplayName(e.target.value)}
              placeholder="Shown in UI (optional)"
            />
          </div>
          <div className="grid gap-2">
            <Label>Part of speech</Label>
            <Select value={posId} onValueChange={setPosId}>
              <SelectTrigger>
                <SelectValue placeholder="(optional)" />
              </SelectTrigger>
              <SelectContent>
                {posList.map((p) => (
                  <SelectItem key={p.id} value={p.id}>
                    {p.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>
        <DialogFooter>
          <Button onClick={handleCreate} disabled={!name.trim() || loading}>
            {loading && <Loader2 className="h-4 w-4 mr-2 animate-spin" />}
            Create
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
