import React from "react";
import { Button } from "@/components/ui/button";
import { RefreshCcw } from "lucide-react";

export default function Toolbar(props: {
  onRefresh?: () => void;
  right?: React.ReactNode;
  title: string;
  subtitle?: string;
}) {
  return (
    <div className="flex items-center justify-between gap-4">
      <div>
        <div className="text-xl font-semibold">{props.title}</div>
        {props.subtitle && (
          <div className="text-sm text-muted-foreground">{props.subtitle}</div>
        )}
      </div>
      <div className="flex items-center gap-2">
        {props.onRefresh && (
          <Button variant="outline" size="sm" onClick={props.onRefresh}>
            <RefreshCcw className="h-4 w-4 mr-2" /> Refresh
          </Button>
        )}
        {props.right}
      </div>
    </div>
  );
}
