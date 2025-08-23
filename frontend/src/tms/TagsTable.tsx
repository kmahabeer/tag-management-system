import React, { useEffect, useState } from "react";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import Toolbar from "./Toolbar";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Table,
  TableBody,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { api } from "./api";
import { Tag } from "./types";
import { mockListTags, useMock } from "./mock";
import CreateTagDialog from "./CreateTagDialog";
import EditableTagRow from "./EditableTagRow";

export default function TagsTable({ baseUrl }: { baseUrl?: string }) {
  const [rows, setRows] = useState<Tag[]>([]);
  const [loading, setLoading] = useState(true);

  async function load() {
    try {
      setLoading(true);
      const data = useMock
        ? await mockListTags()
        : await api<Tag[]>("/tags", { baseUrl });
      setRows(data);
    } catch (e: any) {
      toast.error(e.message ?? "Failed to load tags");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    load();
  }, [baseUrl]);

  return (
    <Card className="shadow-sm">
      <CardHeader className="pb-3">
        <Toolbar
          title="Tags"
          subtitle="Create, edit, and delete tags"
          onRefresh={load}
          right={
            <CreateTagDialog
              baseUrl={baseUrl}
              onCreated={(t) => setRows((prev) => [t, ...prev])}
            />
          }
        />
      </CardHeader>
      <CardContent>
        <ScrollArea className="max-h-[420px] rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Display name</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {loading ? (
                <TableRow>
                  <td
                    colSpan={3}
                    className="text-center py-8 text-muted-foreground"
                  >
                    <Loader2 className="h-4 w-4 inline mr-2 animate-spin" />{" "}
                    Loading…
                  </td>
                </TableRow>
              ) : rows.length === 0 ? (
                <TableRow>
                  <td
                    colSpan={3}
                    className="text-center py-10 text-muted-foreground"
                  >
                    No tags yet. Create your first tag.
                  </td>
                </TableRow>
              ) : (
                rows.map((t) => (
                  <EditableTagRow
                    key={t.id}
                    tag={t}
                    onChanged={(u) =>
                      setRows((prev) =>
                        prev.map((x) => (x.id === u.id ? u : x))
                      )
                    }
                    onDeleted={(id) =>
                      setRows((prev) => prev.filter((x) => x.id !== id))
                    }
                    baseUrl={baseUrl}
                  />
                ))
              )}
            </TableBody>
          </Table>
        </ScrollArea>
      </CardContent>
    </Card>
  );
}
