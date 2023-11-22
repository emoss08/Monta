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

import * as yup from "yup";
import {
  ReasonCodeFormValues,
  ServiceTypeFormValues,
  ShipmentTypeFormValues,
} from "@/types/order";
import { StatusChoiceProps } from "@/types";
import { CodeTypeProps } from "@/lib/choices";

export const serviceTypeSchema: yup.ObjectSchema<ServiceTypeFormValues> = yup
  .object()
  .shape({
    status: yup.string<StatusChoiceProps>().required("Status is required"),
    code: yup
      .string()
      .max(10, "Code must be at most 10 characters")
      .required("Code is required"),
    description: yup.string().notRequired(),
  });

export const reasonCodeSchema: yup.ObjectSchema<ReasonCodeFormValues> = yup
  .object()
  .shape({
    status: yup.string<StatusChoiceProps>().required("Status is required"),
    code: yup
      .string()
      .max(10, "Code must be at most 10 characters")
      .required("Code is required"),
    codeType: yup.string<CodeTypeProps>().required("Code type is required"),
    description: yup.string().required("Description is required"),
  });

export const shipmentTypeSchema: yup.ObjectSchema<ShipmentTypeFormValues> = yup
  .object()
  .shape({
    status: yup.string<StatusChoiceProps>().required("Status is required"),
    code: yup
      .string()
      .max(10, "Name must be at most 100 characters")
      .required("Code is required"),
    description: yup.string().notRequired(),
  });