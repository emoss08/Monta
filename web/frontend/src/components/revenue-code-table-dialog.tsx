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

import { InputField } from "@/components/common/fields/input";
import { SelectInput } from "@/components/common/fields/select-input";
import { Button } from "@/components/ui/button";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { statusChoices } from "@/lib/choices";
import { revenueCodeSchema } from "@/lib/validations/AccountingSchema";
import { type RevenueCodeFormValues as FormValues } from "@/types/accounting";
import { type TableSheetProps } from "@/types/tables";
import { yupResolver } from "@hookform/resolvers/yup";
import { Control, useForm } from "react-hook-form";
import { AsyncSelectInput } from "./common/fields/async-select-input";
import { GradientPicker } from "./common/fields/color-field";
import { TextareaField } from "./common/fields/textarea";
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
import { Form, FormControl, FormGroup } from "./ui/form";

export function RCForm({ control }: { control: Control<FormValues> }) {
  return (
    <Form>
      <FormGroup className="grid gap-6 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-2">
        <FormControl>
          <SelectInput
            name="status"
            rules={{ required: true }}
            control={control}
            label="Status"
            options={statusChoices}
            placeholder="Select Status"
            description="Operational status of the Revenue Code"
            isClearable={false}
          />
        </FormControl>
        <FormControl>
          <InputField
            control={control}
            rules={{ required: true }}
            name="code"
            label="Code"
            autoCapitalize="none"
            autoCorrect="off"
            type="text"
            placeholder="Code"
            description="Code for the Revenue Code"
          />
        </FormControl>
        <FormControl className="col-span-full">
          <TextareaField
            name="description"
            rules={{ required: true }}
            control={control}
            label="Description"
            placeholder="Description"
            description="Description of the Revenue Code"
          />
        </FormControl>
        <FormControl>
          <AsyncSelectInput
            name="expenseAccountId"
            control={control}
            link="/general-ledger-accounts/"
            valueKey="accountNumber"
            label="Expense Account"
            placeholder="Select Expense Account"
            description="The Expense Account associated with the Revenue Code"
            isClearable
            hasPopoutWindow
            popoutLink="/accounting/gl-accounts"
            popoutLinkLabel="GL Account"
          />
        </FormControl>
        <FormControl>
          <AsyncSelectInput
            name="revenueAccountId"
            control={control}
            link="/general-ledger-accounts/"
            valueKey="accountNumber"
            label="Revenue Account"
            placeholder="Select Revenue Account"
            description="The Revneue Account associated with the Revenue Code"
            isClearable
            hasPopoutWindow
            popoutLink="/accounting/gl-accounts"
            popoutLinkLabel="GL Account"
          />
        </FormControl>
        <FormControl className="col-span-full">
          <GradientPicker
            name="color"
            label="Color"
            description="Color Code of the Revenue Code"
            control={control}
          />
        </FormControl>
      </FormGroup>
    </Form>
  );
}

export function RevenueCodeDialog({ onOpenChange, open }: TableSheetProps) {
  const { control, reset, handleSubmit } = useForm<FormValues>({
    resolver: yupResolver(revenueCodeSchema),
    defaultValues: {
      status: "Active",
      code: "",
      description: "",
      expenseAccountId: null,
      revenueAccountId: null,
    },
  });

  const mutation = useCustomMutation<FormValues>(control, {
    method: "POST",
    path: "/revenue-codes/",
    successMessage: "Revenue Code created successfully.",
    queryKeysToInvalidate: "revenueCodes",
    closeModal: true,
    reset,
    errorMessage: "Failed to create new revenue code.",
  });

  const onSubmit = (values: FormValues) => mutation.mutate(values);

  return (
    <Credenza open={open} onOpenChange={onOpenChange}>
      <CredenzaContent>
        <CredenzaHeader>
          <CredenzaTitle>Create New Revenue Code</CredenzaTitle>
        </CredenzaHeader>
        <CredenzaDescription>
          Please fill out the form below to create a new Revenue Code.
        </CredenzaDescription>
        <CredenzaBody>
          <form onSubmit={handleSubmit(onSubmit)}>
            <RCForm control={control} />
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
      </CredenzaContent>
    </Credenza>
  );
}
