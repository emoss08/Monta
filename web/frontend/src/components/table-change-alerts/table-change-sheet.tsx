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

import { DatepickerField } from "@/components/common/fields/date-picker";
import { TextareaField } from "@/components/common/fields/textarea";
import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { useTopics } from "@/hooks/useQueries";
import {
  EnumDatabaseAction,
  EnumDeliveryMethod,
  databaseActionChoices,
  deliveryMethodChoices,
  statusChoices,
} from "@/lib/choices";
import { cn } from "@/lib/utils";
import { tableChangeAlertSchema } from "@/lib/validations/OrganizationSchema";
import { type TableChangeAlertFormValues as FormValues } from "@/types/organization";
import { type TableSheetProps } from "@/types/tables";
import { yupResolver } from "@hookform/resolvers/yup";
import { useState } from "react";
import { FormProvider, useForm, type Control } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { InputField } from "../common/fields/input";
import { SelectInput } from "../common/fields/select-input";
import { FormControl, FormGroup } from "../ui/form";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/new/new-tabs";

export function TableChangeAlertForm({
  control,
}: {
  control: Control<FormValues>;
}) {
  const { t } = useTranslation("admin.tablechangealert");
  // const { control, watch } = useFormContext<FormValues>();
  const {
    selectTopics,
    isError: isTopicError,
    isLoading: isTopicsLoading,
  } = useTopics();

  // const { selectEmailProfile, isError, isLoading } = useEmailProfiles(open);

  return (
    <FormGroup className="lg:grid-cols-2">
      <FormControl>
        <SelectInput
          name="status"
          rules={{ required: true }}
          control={control}
          options={statusChoices}
          isClearable={false}
          label={t("fields.status.label")}
          placeholder={t("fields.status.placeholder")}
          description={t("fields.status.description")}
        />
      </FormControl>
      <FormControl>
        <InputField
          name="name"
          rules={{ required: true }}
          control={control}
          label={t("fields.name.label")}
          placeholder={t("fields.name.placeholder")}
          description={t("fields.name.description")}
        />
      </FormControl>
      <FormControl>
        <SelectInput
          name="databaseAction"
          rules={{ required: true }}
          options={databaseActionChoices}
          control={control}
          label={t("fields.databaseAction.label")}
          placeholder={t("fields.databaseAction.placeholder")}
          description={t("fields.databaseAction.description")}
        />
      </FormControl>
      <FormControl>
        <SelectInput
          name="topicName"
          rules={{ required: true }}
          options={selectTopics}
          isLoading={isTopicsLoading}
          isFetchError={isTopicError}
          control={control}
          label={t("fields.topic.label")}
          placeholder={t("fields.topic.placeholder")}
          description={t("fields.topic.description")}
        />
      </FormControl>
      <FormControl className="col-span-full">
        <SelectInput
          name="deliveryMethod"
          rules={{ required: true }}
          options={deliveryMethodChoices}
          control={control}
          label={t("fields.deliveryMethod.label")}
          placeholder={t("fields.deliveryMethod.placeholder")}
          description={t("fields.deliveryMethod.description")}
        />
      </FormControl>
      <FormControl className="col-span-full">
        <TextareaField
          name="description"
          control={control}
          label={t("fields.description.label")}
          placeholder={t("fields.description.placeholder")}
          description={t("fields.description.description")}
        />
      </FormControl>
      <FormControl className="col-span-full">
        <InputField
          name="customSubject"
          control={control}
          label={t("fields.customSubject.label")}
          placeholder={t("fields.customSubject.placeholder")}
          description={t("fields.customSubject.description")}
        />
      </FormControl>
      <FormControl className="col-span-full">
        <InputField
          name="emailRecipients"
          rules={{ required: true }}
          control={control}
          label={t("fields.emailRecipients.label")}
          placeholder={t("fields.emailRecipients.placeholder")}
          description={t("fields.emailRecipients.description")}
        />
      </FormControl>
      <FormControl>
        <DatepickerField
          name="effectiveDate"
          control={control}
          label={t("fields.effectiveDate.label")}
          placeholder={t("fields.effectiveDate.placeholder")}
          description={t("fields.effectiveDate.description")}
        />
      </FormControl>
      <FormControl>
        <DatepickerField
          name="expirationDate"
          control={control}
          label={t("fields.expirationDate.label")}
          placeholder={t("fields.expirationDate.placeholder")}
          description={t("fields.expirationDate.description")}
        />
      </FormControl>
    </FormGroup>
  );
}

export function TableChangeAlertBody({
  control,
}: {
  control: Control<FormValues>;
}) {
  const [activeTab, setActiveTab] = useState<string>("info");

  return (
    <Tabs
      defaultValue="info"
      value={activeTab}
      className="mt-10 w-full flex-1"
      onValueChange={setActiveTab}
    >
      <TabsList className="mx-auto space-x-4">
        <TabsTrigger value="info">General Information</TabsTrigger>
        <TabsTrigger value="conditionalLogic">Conditional Logic</TabsTrigger>
      </TabsList>
      <TabsContent value="info">
        <TableChangeAlertForm control={control} />
      </TabsContent>
      <TabsContent value="conditionalLogic">
        <div>
          <p>Coming Soon...</p>
        </div>
      </TabsContent>
    </Tabs>
  );
}

export function TableChangeAlertSheet({ onOpenChange, open }: TableSheetProps) {
  const { t } = useTranslation(["admin.tablechangealert", "common"]);

  const tableChangeAlertForm = useForm<FormValues>({
    resolver: yupResolver(tableChangeAlertSchema),
    defaultValues: {
      status: "Active",
      name: "",
      databaseAction: EnumDatabaseAction.Insert,
      topicName: "",
      description: "",
      emailRecipients: "",
      deliveryMethod: EnumDeliveryMethod.Email,
      // conditionalLogic: {},
      customSubject: "",
      effectiveDate: null,
      expirationDate: null,
    },
  });

  const { control, reset, handleSubmit } = tableChangeAlertForm;

  const mutation = useCustomMutation<FormValues>(control, {
    method: "POST",
    path: "/table-change-alerts/",
    successMessage: t("formMessages.postSuccess"),
    queryKeysToInvalidate: "tableChangeAlerts",
    closeModal: true,
    reset,
    errorMessage: t("formMessages.postError"),
  });

  const onSubmit = (values: FormValues) => mutation.mutate(values);

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className={cn("w-full xl:w-[700px]")}>
        <SheetHeader>
          <SheetTitle>{t("title")}</SheetTitle>
          <SheetDescription>{t("subTitle")}</SheetDescription>
        </SheetHeader>
        <FormProvider {...tableChangeAlertForm}>
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex h-full flex-col overflow-y-auto"
          >
            <TableChangeAlertBody control={control} />
            <SheetFooter className="mb-12">
              <Button
                type="reset"
                variant="secondary"
                onClick={() => onOpenChange(false)}
                className="w-full"
              >
                {t("buttons.cancel", { ns: "common" })}
              </Button>
              <Button
                type="submit"
                isLoading={mutation.isPending}
                className="w-full"
              >
                {t("buttons.save", { ns: "common" })}
              </Button>
            </SheetFooter>
          </form>
        </FormProvider>
      </SheetContent>
    </Sheet>
  );
}
