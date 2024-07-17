/**
 * Copyright (c) 2024 Trenova Technologies, LLC
 *
 * Licensed under the Business Source License 1.1 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://trenova.app/pricing/
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *
 * Key Terms:
 * - Non-production use only
 * - Change Date: 2026-11-16
 * - Change License: GNU General Public License v2 or later
 *
 * For full license text, see the LICENSE file in the root directory.
 */

import { Checkbox } from "@/components/common/fields/checkbox";
import { DataTable } from "@/components/common/table/data-table";
import { DataTableColumnHeader } from "@/components/common/table/data-table-column-header";
import { StatusBadge } from "@/components/common/table/data-table-components";
import { EquipTypeEditSheet } from "@/components/equipment-type-edit-table-dialog";
import { EquipTypeDialog } from "@/components/equipment-type-table-dialog";
import { equipmentClassChoices, tableStatusChoices } from "@/lib/choices";
import { truncateText } from "@/lib/utils";
import { type EquipmentType } from "@/types/equipment";
import { type FilterConfig } from "@/types/tables";
import { type ColumnDef } from "@tanstack/react-table";

const columns: ColumnDef<EquipmentType>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="translate-y-[2px]"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
        className="translate-y-[2px]"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    id: "status",
    accessorFn: (row) => row.status,
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Status" />
    ),
    cell: ({ row }) => <StatusBadge status={row.original.status} />,
  },
  {
    id: "code",
    accessorFn: (row) => row.code,
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Code" />
    ),
    cell: ({ row }) => {
      if (row.original.color) {
        return (
          <div className="text-foreground flex items-center space-x-2 text-sm font-medium">
            <div
              className={"mx-2 size-2 rounded-xl"}
              style={{ backgroundColor: row.original.color }}
            />
            {row.original.code}
          </div>
        );
      }

      return row.original.code;
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id));
    },
  },
  {
    accessorKey: "description",
    header: "Description",
    cell: ({ row }) => truncateText(row.original.description as string, 30),
  },
  {
    accessorKey: "equipmentClass",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Equip. Class" />
    ),
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id));
    },
  },
];

const filters: FilterConfig<EquipmentType>[] = [
  {
    columnName: "status",
    title: "Status",
    options: tableStatusChoices,
  },
  {
    columnName: "equipmentClass",
    title: "Equip. Class",
    options: equipmentClassChoices,
  },
];

export default function EquipmentTypes() {
  return (
    <DataTable
      queryKey="equipmentTypes"
      columns={columns}
      link="/equipment-types/"
      name="Equip. Types"
      exportModelName="equipment_types"
      filterColumn="code"
      tableFacetedFilters={filters}
      TableSheet={EquipTypeDialog}
      TableEditSheet={EquipTypeEditSheet}
      addPermissionName="equipmenttype.add"
    />
  );
}
