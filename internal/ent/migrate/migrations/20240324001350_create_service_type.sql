-- Create "service_types" table
CREATE TABLE "service_types" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'A', "code" character varying NOT NULL, "description" text NULL, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "service_types_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "service_types_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "servicetype_code_organization_id" to table: "service_types"
CREATE UNIQUE INDEX "servicetype_code_organization_id" ON "service_types" ("code", "organization_id");
