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

import type {
  EmailProtocolChoiceProps,
  EnumDatabaseAction,
  EnumDeliveryMethod,
  RouteDistanceUnitProps,
  RouteModelChoiceProps,
  TimezoneChoices,
} from "@/lib/choices";
import { type StatusChoiceProps } from ".";

export type Organization = {
  id: string;
  name: string;
  scacCode: string;
  dotNumber: string;
  orgType: string;
  timezone: TimezoneChoices;
  logoUrl?: string | null;
  addressLine1?: string;
  addressLine2?: string;
  city?: string;
  state?: {
    abbreviation: string;
    name: string;
  };
  postalCode?: string;
};

export type OrganizationFormValues = Omit<Organization, "id">;

// type DeliveryMethod = "Email" | "Api" | "Local" | "Sms";

export interface TableChangeAlert extends BaseModel {
  id: string;
  status: StatusChoiceProps;
  name: string;
  databaseAction: EnumDatabaseAction;
  topicName: string;
  deliveryMethod: EnumDeliveryMethod;
  description?: string;
  emailRecipients?: string;
  // conditionalLogic?: object | null;
  customSubject?: string;
  effectiveDate?: string | null;
  expirationDate?: string | null;
}

export type TableChangeAlertFormValues = Omit<
  TableChangeAlert,
  "id" | "organizationId" | "createdAt" | "updatedAt" | "version"
>;

export interface EmailProfile extends BaseModel {
  id: string;
  name: string;
  email: string;
  protocol?: EmailProtocolChoiceProps | null;
  host?: string | null;
  port?: number | null;
  username?: string | null;
  password?: string | null;
  isDefault: boolean;
}

export type EmailProfileFormValues = Omit<
  EmailProfile,
  "id" | "organizationId" | "createdAt" | "updatedAt" | "version"
>;

export type Department = {
  id: string;
  name: string;
  organization: string;
  description: string;
  depot: string;
};

/** Types for EmailControl */
export interface EmailControl extends BaseModel {
  id: string;
  billingEmailProfileId?: string | null;
  rateExpirtationEmailProfileId?: string | null;
}

export type EmailControlFormValues = Omit<
  EmailControl,
  "id" | "organizationId" | "createdAt" | "updatedAt" | "version"
>;

export type Depot = BaseModel & {
  id: string;
  name: string;
  description?: string;
};

export type FeatureFlag = {
  name: string;
  code: string;
  description: string;
  beta: boolean;
  preview: string;
};

export type OrganizationFeatureFlag = {
  isEnabled: boolean;
  edges: {
    featureFlag: FeatureFlag;
  };
};

export type GoogleAPI = BaseModel & {
  id: string;
  apiKey?: string | null;
  mileageUnit: RouteDistanceUnitProps;
  trafficModel: RouteModelChoiceProps;
  addCustomerLocation: boolean;
  addLocation: boolean;
  autoGeocode: boolean;
};

export type TableName = {
  value: string;
  label: string;
};

export type Topic = {
  value: string;
  label: string;
};

export type GoogleAPIFormValues = Omit<
  GoogleAPI,
  "id" | "organizationId" | "createdAt" | "updatedAt" | "version"
>;

/** Base Trenova Interface
 *
 * @note This interface is used for all Trenova models that have the following fields:
 * - organization
 * - created
 * - modified
 *
 * Please do not put businessUnit in this interface. Add it directly to the interface that
 * extends this interface.
 * */
export type BaseModel = {
  organizationId: string;
  version: number;
  createdAt: string;
  updatedAt: string;
};
