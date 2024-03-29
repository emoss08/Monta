-- Modify "document_classifications" table
ALTER TABLE "document_classifications" ADD COLUMN "status" character varying NOT NULL DEFAULT 'A';
-- Create "hazardous_material_segregations" table
CREATE TABLE "hazardous_material_segregations" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "class_a" character varying NOT NULL DEFAULT 'HazardClass1And1', "class_b" character varying NOT NULL DEFAULT 'HazardClass1And1', "segregation_type" character varying NOT NULL DEFAULT 'NotAllowed', "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "hazardous_material_segregations_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "hazardous_material_segregations_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "hazardousmaterialsegregation_class_a_class_b_organization_id" to table: "hazardous_material_segregations"
CREATE UNIQUE INDEX "hazardousmaterialsegregation_class_a_class_b_organization_id" ON "hazardous_material_segregations" ("class_a", "class_b", "organization_id");
