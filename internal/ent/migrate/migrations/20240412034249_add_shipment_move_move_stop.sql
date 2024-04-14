-- Modify "shipments" table
ALTER TABLE "shipments" ALTER COLUMN "status" SET DEFAULT 'New';
-- Create "shipment_moves" table
CREATE TABLE "shipment_moves" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "version" bigint NOT NULL DEFAULT 1, "reference_number" character varying(10) NOT NULL, "status" character varying NOT NULL DEFAULT 'New', "is_loaded" boolean NOT NULL DEFAULT false, "shipment_id" uuid NOT NULL, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, "tractor_id" uuid NULL, "trailer_id" uuid NULL, "primary_worker_id" uuid NULL, "secondary_worker_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "shipment_moves_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "shipment_moves_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "shipment_moves_shipments_shipment_moves" FOREIGN KEY ("shipment_id") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "shipment_moves_tractors_tractor" FOREIGN KEY ("tractor_id") REFERENCES "tractors" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "shipment_moves_tractors_trailer" FOREIGN KEY ("trailer_id") REFERENCES "tractors" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "shipment_moves_workers_primary_worker" FOREIGN KEY ("primary_worker_id") REFERENCES "workers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "shipment_moves_workers_secondary_worker" FOREIGN KEY ("secondary_worker_id") REFERENCES "workers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "shipment_moves_reference_number_key" to table: "shipment_moves"
CREATE UNIQUE INDEX "shipment_moves_reference_number_key" ON "shipment_moves" ("reference_number");
-- Create "stops" table
CREATE TABLE "stops" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "version" bigint NOT NULL DEFAULT 1, "status" character varying NOT NULL DEFAULT 'New', "stop_type" character varying NOT NULL, "sequence" bigint NOT NULL DEFAULT 1, "location_id" uuid NULL, "pieces" numeric(10,2) NULL, "weight" numeric(10,2) NULL, "address_line" character varying NULL, "appointment_start" timestamptz NULL, "appointment_end" timestamptz NULL, "arrival_time" timestamptz NULL, "departure_time" timestamptz NULL, "shipment_move_id" uuid NOT NULL, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "stops_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "stops_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "stops_shipment_moves_move_stops" FOREIGN KEY ("shipment_move_id") REFERENCES "shipment_moves" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
