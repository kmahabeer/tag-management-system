import React from "react";
import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Link2 } from "lucide-react";
import TagsTable from "./TagsTable";
import RelationshipsPanel from "./RelationshipsPanel";
import { DEFAULT_BASE_URL } from "./api";
import { useMock } from "./mock";

export default function TagManagementMock({
  baseUrl = DEFAULT_BASE_URL,
}: {
  baseUrl?: string;
}) {
  return (
    <div className="p-6 mx-auto max-w-6xl">
      <div className="mb-6">
        <Card className="bg-background/60 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-muted/50">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Link2 className="h-5 w-5" /> Tag Management System
            </CardTitle>
            <CardDescription>
              Quick mock UI to test read/write tags and relationships. Works in
              local mock mode.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-sm text-muted-foreground">
              API Base:{" "}
              <code className="rounded bg-muted px-2 py-0.5">{baseUrl}</code>
              {useMock && (
                <span className="ml-2 inline-flex items-center gap-2">
                  <Badge className="uppercase tracking-wide">Mock Mode</Badge>
                </span>
              )}
            </div>
          </CardContent>
        </Card>
      </div>
      <Tabs defaultValue="tags" className="w-full">
        <TabsList className="grid w-full grid-cols-2 md:w-auto md:inline-flex">
          <TabsTrigger value="tags">Tags</TabsTrigger>
          <TabsTrigger value="relationships">Relationships</TabsTrigger>
        </TabsList>
        <TabsContent value="tags" className="mt-6">
          <TagsTable baseUrl={baseUrl} />
        </TabsContent>
        <TabsContent value="relationships" className="mt-6">
          <RelationshipsPanel baseUrl={baseUrl} />
        </TabsContent>
      </Tabs>
      <footer className="mt-8 text-xs text-muted-foreground">
        Mock mode is active.
      </footer>
    </div>
  );
}
