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

import { CommodityForm } from "@/components/commodity-dialog";
import { Button } from "@/components/ui/button";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { formatToUserTimezone } from "@/lib/date";
import { commoditySchema } from "@/lib/validations/CommoditiesSchema";
import { useTableStore } from "@/stores/TableStore";
import type { Commodity, CommodityFormValues } from "@/types/commodities";
import { yupResolver } from "@hookform/resolvers/yup";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { Badge } from "./ui/badge";
import {
  Credenza,
  CredenzaBody,
  CredenzaClose,
  CredenzaContent,
  CredenzaDescription,
  CredenzaFooter,
  CredenzaHeader,
  CredenzaTitle,
} from "./ui/credenza";

function CommodityEditForm({ commodity }: { commodity: Commodity }) {
  const { control, handleSubmit, watch, reset, setValue } =
    useForm<CommodityFormValues>({
      resolver: yupResolver(commoditySchema),
      defaultValues: commodity,
    });

  const mutation = useCustomMutation<CommodityFormValues>(control, {
    method: "PUT",
    path: `/commodities/${commodity.id}/`,
    successMessage: "Commodity updated successfully.",
    queryKeysToInvalidate: "commodities",
    closeModal: true,
    reset,
    errorMessage: "Failed to update commodity.",
  });

  const onSubmit = (values: CommodityFormValues) => mutation.mutate(values);

  useEffect(() => {
    const subscription = watch((value, { name }) => {
      if (name === "hazardousMaterialId" && value.hazardousMaterialId) {
        setValue("isHazmat", true);
      } else if (name === "hazardousMaterialId" && !value.hazardousMaterialId) {
        setValue("isHazmat", false);
      }
    });

    return () => subscription.unsubscribe();
  }, [watch, setValue]);

  return (
    <CredenzaBody>
      <form onSubmit={handleSubmit(onSubmit)}>
        <CommodityForm control={control} />
        <CredenzaFooter>
          <CredenzaClose asChild>
            <Button variant="outline" type="button">
              Cancel
            </Button>
          </CredenzaClose>
          <Button type="submit" isLoading={mutation.isPending}>
            Save Changes
          </Button>
        </CredenzaFooter>
      </form>
    </CredenzaBody>
  );
}

export function CommodityEditDialog({
  open,
  onOpenChange,
}: {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  const [commodity] = useTableStore.use("currentRecord") as Commodity[];

  if (!commodity) return null;

  return (
    <Credenza open={open} onOpenChange={onOpenChange}>
      <CredenzaContent>
        <CredenzaHeader>
          <CredenzaTitle className="flex">
            <span>{commodity.name}</span>
            <Badge className="ml-5" variant="purple">
              {commodity.id}
            </Badge>
          </CredenzaTitle>
        </CredenzaHeader>
        <CredenzaDescription>
          Last updated on&nbsp;
          {formatToUserTimezone(commodity.updatedAt)}
        </CredenzaDescription>
        <CommodityEditForm commodity={commodity} />
      </CredenzaContent>
    </Credenza>
  );
}
