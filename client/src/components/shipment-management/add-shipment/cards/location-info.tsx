/*
 * COPYRIGHT(c) 2024 Trenova
 *
 * This file is part of Trenova.
 *
 * The Trenova software is licensed under the Business Source License 1.1. You are granted the right
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

import { TimeField } from "@/components/common/fields/input";
import { AsyncSelectInput } from "@/components/common/fields/select-input";
import { LocationAutoComplete } from "@/components/ui/autocomplete";
import { Skeleton } from "@/components/ui/skeleton";
import { TitleWithTooltip } from "@/components/ui/title-with-tooltip";
import { useLocations } from "@/hooks/useQueries";
import { Location } from "@/types/location";
import { ShipmentControl, ShipmentFormValues } from "@/types/order";
import { useEffect } from "react";
import { Control, useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";

function LocationSection({
  section,
  control,
}: {
  section: "origin" | "destination";
  control: Control<ShipmentFormValues>;
}) {
  const { t } = useTranslation("shipment.addshipment");

  return (
    <div className="flex-1">
      <div className="flex flex-col">
        <div className="border-border rounded-md border">
          <div className="border-border bg-background flex justify-center rounded-t-md border-b p-2">
            <TitleWithTooltip
              title={t(`card.${section}.label`)}
              tooltip={t(`card.${section}.description`)}
            />
          </div>
          <div className="bg-card grid grid-cols-1 gap-y-4 p-4">
            <div className="col-span-3">
              <AsyncSelectInput
                name={`${section}Location`}
                link="/locations/"
                control={control}
                label={t(`card.${section}.fields.${section}Location.label`)}
                placeholder={t(
                  `card.${section}.fields.${section}Location.placeholder`,
                )}
                description={t(
                  `card.${section}.fields.${section}Location.description`,
                )}
                hasPopoutWindow
                popoutLink="/dispatch/locations/"
                isClearable
                popoutLinkLabel="Location"
              />
            </div>
            <div className="col-span-3">
              <LocationAutoComplete
                control={control}
                name={`${section}Address`}
                rules={{ required: true }}
                autoCapitalize="none"
                autoCorrect="off"
                type="text"
                label={t(`card.${section}.fields.${section}Address.label`)}
                placeholder={t(
                  `card.${section}.fields.${section}Address.placeholder`,
                )}
                description={t(
                  `card.${section}.fields.${section}Address.description`,
                )}
              />
            </div>
            <div className="grid grid-cols-2 gap-x-4">
              <div className="col-span-1">
                <TimeField
                  control={control}
                  rules={{ required: true }}
                  name={`${section}AppointmentWindowStart`}
                  label={t(
                    `card.${section}.fields.${section}AppointmentWindowStart.label`,
                  )}
                  placeholder={t(
                    `card.${section}.fields.${section}AppointmentWindowStart.placeholder`,
                  )}
                  description={t(
                    `card.${section}.fields.${section}AppointmentWindowStart.description`,
                  )}
                />
              </div>
              <div className="col-span-1">
                <TimeField
                  control={control}
                  rules={{ required: true }}
                  name={`${section}AppointmentWindowEnd`}
                  label={t(
                    `card.${section}.fields.${section}AppointmentWindowEnd.label`,
                  )}
                  placeholder={t(
                    `card.${section}.fields.${section}AppointmentWindowEnd.placeholder`,
                  )}
                  description={t(
                    `card.${section}.fields.${section}AppointmentWindowEnd.description`,
                  )}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export function LocationInformation({
  shipmentControlData,
  isShipmentControlLoading,
}: {
  shipmentControlData: ShipmentControl;
  isShipmentControlLoading: boolean;
}) {
  const { locations } = useLocations();
  const { control, setValue, watch, setError } =
    useFormContext<ShipmentFormValues>();
  const originLocationValue = watch("originLocation");
  const originAddressValue = watch("originAddress");
  const destinationLocationValue = watch("destinationLocation");
  const destinationAddressValue = watch("destinationAddress");
  useEffect(() => {
    if (originLocationValue && locations) {
      const selectedOriginLocation = (locations as Location[]).find(
        (location) => location.id === originLocationValue,
      );

      if (selectedOriginLocation) {
        setValue(
          "originAddress",
          `${selectedOriginLocation.addressLine1}, ${selectedOriginLocation.city}, ${selectedOriginLocation.state} ${selectedOriginLocation.zipCode}`,
        );
      }
    }

    if (destinationLocationValue && locations) {
      const selectedDestinationLocation = (locations as Location[]).find(
        (location) => location.id === destinationLocationValue,
      );

      if (selectedDestinationLocation) {
        setValue(
          "destinationAddress",
          `${selectedDestinationLocation.addressLine1}, ${selectedDestinationLocation.city}, ${selectedDestinationLocation.state} ${selectedDestinationLocation.zipCode}`,
        );
      }
    }

    // If the origin address and destination address are the same then throw an error
    if (
      shipmentControlData &&
      shipmentControlData?.enforceOriginDestination &&
      originAddressValue &&
      destinationAddressValue &&
      originAddressValue === destinationAddressValue
    ) {
      setError("originAddress", {
        type: "manual",
        message: "Origin and Destination locations cannot be the same.",
      });
      setError("destinationAddress", {
        type: "manual",
        message: "Origin and Destination locations cannot be the same.",
      });
    }
  }, [
    originLocationValue,
    originAddressValue,
    destinationLocationValue,
    destinationAddressValue,
    locations,
    setValue,
    setError,
    shipmentControlData,
    shipmentControlData?.enforceOriginDestination,
  ]);

  if (isShipmentControlLoading) {
    return <Skeleton className="h-[40vh]" />;
  }

  return (
    <div className="flex space-x-10">
      <LocationSection section="origin" control={control} />
      <LocationSection section="destination" control={control} />
    </div>
  );
}