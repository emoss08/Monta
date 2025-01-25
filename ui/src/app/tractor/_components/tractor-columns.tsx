import { DataTableColumnHeader } from "@/components/data-table/_components/data-table-column-header";
import { DataTableColorColumn } from "@/components/data-table/_components/data-table-components";
import { EquipmentStatusBadge } from "@/components/status-badge";
import { Checkbox } from "@/components/ui/checkbox";
import { type Tractor } from "@/types/tractor";
import { type ColumnDef } from "@tanstack/react-table";

export function getColumns(): ColumnDef<Tractor>[] {
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
          />
        );
      },
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(checked) => row.toggleSelected(!!checked)}
          aria-label="Select row"
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
        return <EquipmentStatusBadge status={status} />;
      },
    },
    {
      accessorKey: "code",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Code" />
      ),
    },
    {
      id: "equipmentType",
      accessorKey: "equipmentType",
      header: "Equipment Type",
      cell: ({ row }) => {
        const equipType = row.original.equipmentType;
        const isEquipType = !!equipType;

        return isEquipType ? (
          <DataTableColorColumn
            color={equipType?.color}
            text={equipType?.code ?? ""}
          />
        ) : (
          <p>No equipment type</p>
        );
      },
    },
    {
      id: "assignedWorkers",
      header: "Assigned Workers",
      cell: ({ row }) => {
        const { primaryWorker, secondaryWorker } = row.original;

        const isPrimaryWorker = !!primaryWorker;
        const isSecondaryWorker = !!secondaryWorker;

        return isPrimaryWorker ? (
          <div className="flex flex-col gap-0.5">
            <p>
              {primaryWorker?.firstName} {primaryWorker?.lastName}
            </p>
            {isSecondaryWorker && (
              <p className="text-2xs text-muted-foreground">
                Co-Driver: {secondaryWorker?.firstName}{" "}
                {secondaryWorker?.lastName}
              </p>
            )}
          </div>
        ) : (
          <p>No assigned workers</p>
        );
      },
    },
  ];
}
