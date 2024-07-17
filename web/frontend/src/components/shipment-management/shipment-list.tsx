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

import { Shipment, ShipmentStatus } from "@/types/shipment";
import { Droppable } from "react-beautiful-dnd";
import { ShipmentInfo } from "./shipment-info";

export function ShipmentList({
  shipments,
  finalStatuses,
  progressStatuses,
}: {
  shipments: Shipment[];
  finalStatuses: ShipmentStatus[];
  progressStatuses: ShipmentStatus[];
}) {
  return (
    <Droppable droppableId="shipmentList">
      {(provided) => (
        <div ref={provided.innerRef} {...provided.droppableProps}>
          {shipments.length > 0 ? (
            shipments.map((shipment) => (
              <Droppable key={shipment.id} droppableId={shipment.id}>
                {(provided, snapshot) => (
                  <div
                    ref={provided.innerRef}
                    {...provided.droppableProps}
                    className={`mb-2 transition-colors duration-200 ${
                      snapshot.isDraggingOver ? "bg-green-500/50" : ""
                    }`}
                  >
                    <ShipmentInfo
                      shipment={shipment}
                      finalStatuses={finalStatuses}
                      progressStatuses={progressStatuses}
                    />
                    <div style={{ display: "none" }}>
                      {provided.placeholder}
                    </div>
                  </div>
                )}
              </Droppable>
            ))
          ) : (
            <div className="text-muted-foreground py-8 text-center">
              No shipments found for the given criteria.
            </div>
          )}
          {provided.placeholder}
        </div>
      )}
    </Droppable>
  );
}
