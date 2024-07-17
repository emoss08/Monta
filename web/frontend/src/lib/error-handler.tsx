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

import { Control, FieldValues, Path } from "react-hook-form";
import { toast } from "sonner";
import { API_BASE_URL } from "./constants";

interface ErrorResponse {
  type: string;
  title: string;
  status: number;
  detail: string;
  instance?: string;
  "invalid-params"?: { name: string; reason: string }[];
  message?: string;
}

export async function handleError<T extends FieldValues>(
  error: any,
  control: Control<T>,
) {
  if (!error.response) {
    console.error("[Trenova] Network or other error", error);
    showErrorNotification("A network or system error occurred.");
    return;
  }

  const { data } = error.response as { data?: ErrorResponse };

  if (!data) {
    console.error("[Trenova] Error without data", error);
    showErrorNotification("An unknown error occurred.");
    return;
  }

  const invalidParams = data["invalid-params"] || [];

  switch (data.instance) {
    case `${API_BASE_URL}/probs/validation-errors`:
      console.info("[Trenova] Validation error", invalidParams);
      handleValidationErrors(invalidParams, control);
      break;
    case `${API_BASE_URL}/probs/database-errors`:
    case `${API_BASE_URL}/probs/invalid-request`: // Combined case for both types of errors
      handleInvalidRequest(invalidParams, control);
      break;
    default:
      console.error("[Trenova] Unhandled error type", data.message);
      showErrorNotification(
        data.message || "An unhandled error type occurred.",
      );
      break;
  }
}

function showErrorNotification(errorMessage?: string) {
  toast.error(
    <div className="flex flex-col space-y-1">
      <span className="font-semibold">Uh Oh! Something went wrong.</span>
      <span className="text-xs">{errorMessage}</span>
    </div>,
  );
}

function handleValidationErrors<T extends FieldValues>(
  invalidParams: { name: string; reason: string }[],
  control: Control<T>,
) {
  invalidParams.forEach((param) => {
    const { name, reason } = param;

    // Set error on the control
    control.setError(name as Path<T>, { type: "manual", message: reason });

    // Show appropriate notification based on the error attribute
    if (name === "nonFieldErrors" || name === "databaseError") {
      showErrorNotification(reason);
    } else {
      showErrorNotification("Please fix the errors and try again.");
    }
  });
}

function handleInvalidRequest<T extends FieldValues>(
  invalidParams: { name: string; reason: string }[],
  control: Control<T>,
) {
  invalidParams.forEach((param) => {
    const { name, reason } = param;

    // Set error on the control
    control.setError(name as Path<T>, { type: "manual", message: reason });

    // Show appropriate notification based on the error attribute
    if (name === "nonFieldErrors" || name === "databaseError") {
      console.log(reason);
      showErrorNotification(reason);
    } else {
      showErrorNotification("Please fix the errors and try again.");
    }
  });
}
