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

import { Button } from "@/components/ui/button";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { formatToUserTimezone } from "@/lib/date";
import { LocationCategorySchema as formSchema } from "@/lib/validations/LocationSchema";
import { useTableStore } from "@/stores/TableStore";
import type {
  LocationCategoryFormValues as FormValues,
  LocationCategory,
} from "@/types/location";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { LCForm } from "./location-category-table-dialog";
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

export function LCEditForm({
  locationCategory,
}: {
  locationCategory: LocationCategory;
}) {
  const { control, reset, handleSubmit } = useForm<FormValues>({
    resolver: yupResolver(formSchema),
    defaultValues: locationCategory,
  });

  const mutation = useCustomMutation<FormValues>(control, {
    method: "PUT",
    path: `/location-categories/${locationCategory.id}/`,
    successMessage: "Location Category updated successfully.",
    queryKeysToInvalidate: "locationCategories",
    closeModal: true,
    reset,
    errorMessage: "Failed to update location category.",
  });

  const onSubmit = (values: FormValues) => mutation.mutate(values);

  return (
    <CredenzaBody>
      <form onSubmit={handleSubmit(onSubmit)}>
        <LCForm control={control} />
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

export function LocationCategoryEditDialog({
  open,
  onOpenChange,
}: {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  const [locationCategory] = useTableStore.use(
    "currentRecord",
  ) as LocationCategory[];

  if (!locationCategory) return null;

  return (
    <Credenza open={open} onOpenChange={onOpenChange}>
      <CredenzaContent>
        <CredenzaHeader>
          <CredenzaTitle className="flex">
            <span>{locationCategory.name}</span>
            <Badge className="ml-5" variant="purple">
              {locationCategory.id}
            </Badge>
          </CredenzaTitle>
        </CredenzaHeader>
        <CredenzaDescription>
          Last updated on {formatToUserTimezone(locationCategory.updatedAt)}
        </CredenzaDescription>
        <LCEditForm locationCategory={locationCategory} />
      </CredenzaContent>
    </Credenza>
  );
}
