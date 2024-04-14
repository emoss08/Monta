-- Create "equipment_manufactuers" table
CREATE TABLE "equipment_manufactuers" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'A', "name" character varying NOT NULL, "description" text NULL, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "equipment_manufactuers_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "equipment_manufactuers_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "equipmentmanufactuer_name_organization_id" to table: "equipment_manufactuers"
CREATE UNIQUE INDEX "equipmentmanufactuer_name_organization_id" ON "equipment_manufactuers" ("name", "organization_id");
