import { DataTableColumnHeader } from "@/components/data-table/_components/data-table-column-header";
import { StatusBadge } from "@/components/status-badge";
import { Checkbox } from "@/components/ui/checkbox";
import { type ServiceTypeSchema } from "@/lib/schemas/service-type-schema";
import { type ColumnDef } from "@tanstack/react-table";

export function getColumns(): ColumnDef<ServiceTypeSchema>[] {
  return [
    {
      accessorKey: "select",
      id: "select",
      header: ({ table }) => {
        return (
          <Checkbox
            checked={
              table.getIsAllPageRowsSelected() ||
              (table.getIsSomePageRowsSelected() && "indeterminate")
            }
            onCheckedChange={(checked) =>
              table.toggleAllPageRowsSelected(!!checked)
            }
            aria-label="Select all"
            className="translate-y-0.5"
          />
        );
      },
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(checked) => row.toggleSelected(!!checked)}
          aria-label="Select row"
          className="translate-y-0.5"
        />
      ),
      enableSorting: false,
      enableHiding: false,
    },
    {
      accessorKey: "status",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Status" />
      ),
      cell: ({ row }) => {
        const status = row.original.status;
        return <StatusBadge status={status} />;
      },
    },
    {
      accessorKey: "code",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Code" />
      ),
      cell: ({ row }) => {
        const isColor = !!row.original.color;
        return isColor ? (
          <div className="flex items-center gap-x-1.5 text-sm font-medium text-foreground">
            <div
              className="size-2 rounded-full"
              style={{
                backgroundColor: row.original.color,
              }}
            />
            <p>{row.original.code}</p>
          </div>
        ) : (
          <p>{row.original.code}</p>
        );
      },
    },
    {
      accessorKey: "description",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Description" />
      ),
    },
  ];
}