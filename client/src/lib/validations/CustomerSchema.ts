/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import {
  StatusChoiceProps,
  TDayOfWeekChoiceProps,
  YesNoChoiceProps,
} from "@/types";
import {
  CreateCustomerFormValues,
  CustomerContactFormValues,
  CustomerEmailProfileFormValues,
  CustomerFormValues,
  CustomerRuleProfileFormValues,
  DeliverySlotFormValues,
} from "@/types/customer";
import * as Yup from "yup";
import { ObjectSchema } from "yup";

/** Customer Schema */
export const customerSchema: ObjectSchema<CustomerFormValues> =
  Yup.object().shape({
    status: Yup.string<StatusChoiceProps>().required("Status is required"),
    code: Yup.string().required("Code is required"),
    name: Yup.string().required("Name is required"),
    addressLine1: Yup.string().notRequired(),
    addressLine2: Yup.string().notRequired(),
    city: Yup.string().notRequired(),
    state: Yup.string().notRequired(),
    zipCode: Yup.string().notRequired(),
    hasCustomerPortal: Yup.string<YesNoChoiceProps>().required(
      "Has Customer Portal is required",
    ),
    autoMarkReadyToBill: Yup.string<YesNoChoiceProps>().required(
      "Auto Mark Ready to Bill is required",
    ),
  });

/** Customer Email Profile Schema */
export const CustomerEmailProfileSchema: ObjectSchema<CustomerEmailProfileFormValues> =
  Yup.object().shape({
    subject: Yup.string().notRequired().max(100),
    comment: Yup.string().notRequired().max(100),
    fromAddress: Yup.string().notRequired(),
    blindCopy: Yup.string().notRequired(),
    readReceipt: Yup.boolean().required(),
    readReceiptTo: Yup.string().when("read_receipt", {
      is: true,
      then: (schema) => schema.required("Read Receipt To is required"),
      otherwise: (schema) => schema.notRequired(),
    }),
    attachmentName: Yup.string().notRequired(),
  });

export const CustomerRuleProfileSchema: ObjectSchema<CustomerRuleProfileFormValues> =
  Yup.object().shape({
    name: Yup.string().required("Name is required"),
    customer: Yup.string().required("Customer is required"),
    documentClass: Yup.array()
      .of(Yup.string().required())
      .min(1, "At Least one document class is required.")
      .required("Document Class is required"),
  });

const DeliverySlotSchema: ObjectSchema<DeliverySlotFormValues> =
  Yup.object().shape({
    dayOfWeek: Yup.string<TDayOfWeekChoiceProps>().required(
      "Day of Week is required",
    ),
    startTime: Yup.string().required("Start Time is required"),
    endTime: Yup.string().required("End Time is required"),
    location: Yup.string().required("Location is required"),
  });

const CustomerContactSchema: ObjectSchema<CustomerContactFormValues> =
  Yup.object().shape({
    isActive: Yup.boolean().required(),
    name: Yup.string().required("Name is required"),
    email: Yup.string().required("Email is required"),
    title: Yup.string().required("Title is required"),
    phone: Yup.string().notRequired(),
    isPayableContact: Yup.boolean().required(),
  });

export const CreateCustomerSchema: ObjectSchema<CreateCustomerFormValues> =
  Yup.object().shape({
    status: Yup.string<StatusChoiceProps>().required("Status is required"),
    name: Yup.string().required("Name is required"),
    addressLine1: Yup.string().required("Address Line 1 is required"),
    addressLine2: Yup.string().notRequired(),
    city: Yup.string().required("City is required"),
    state: Yup.string().required("State is required"),
    zipCode: Yup.string().required("Zip Code is required"),
    hasCustomerPortal: Yup.boolean(),
    autoMarkReadyToBill: Yup.boolean(),
    advocate: Yup.string().notRequired(),
    deliverySlots: Yup.array().of(DeliverySlotSchema).notRequired(),
    emailProfile: Yup.object().shape({
      subject: Yup.string().notRequired().max(100),
      comment: Yup.string().notRequired().max(100),
      fromAddress: Yup.string().email("Invalid email address").notRequired(),
      blindCopy: Yup.string()
        .notRequired()
        .test("is-valid-emails", "Invalid email or list of emails", (value) => {
          if (!value) return true; // Skip validation if empty
          return value
            .split(",")
            .every((email) => Yup.string().email().isValidSync(email.trim()));
        }),
      readReceipt: Yup.boolean().required(),
      readReceiptTo: Yup.string().when("readReceipt", {
        is: true,
        then: (schema) =>
          schema
            .required("Read Receipt To is required")
            .test(
              "is-valid-emails",
              "Invalid email or list of emails",
              (value) => {
                if (!value) return false; // Mandatory if readReceipt is true
                return value
                  .split(",")
                  .every((email) =>
                    Yup.string().email().isValidSync(email.trim()),
                  );
              },
            ),
        // make sure it is a comma email or list of emails regardless
        otherwise: (schema) =>
          schema.test(
            "is-valid-emails",
            "Invalid email or list of emails",
            (value) => {
              if (!value) return true; // Skip validation if empty
              return value
                .split(",")
                .every((email) =>
                  Yup.string().email().isValidSync(email.trim()),
                );
            },
          ),
      }),
      attachmentName: Yup.string().notRequired(),
    }),
    ruleProfile: Yup.object().shape({
      name: Yup.string().required("Name is required"),
      documentClass: Yup.array()
        .of(Yup.string().required())
        .min(1, "At Least one document class is required.")
        .required(),
    }),
    customerContacts: Yup.array().of(CustomerContactSchema).notRequired(),
  });
