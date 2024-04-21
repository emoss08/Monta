-- Create "user_notifications" table
CREATE TABLE "user_notifications" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "version" bigint NOT NULL DEFAULT 1, "is_read" boolean NOT NULL DEFAULT false, "title" character varying NOT NULL, "description" text NOT NULL, "action_url" character varying NULL, "user_id" uuid NOT NULL, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "user_notifications_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_notifications_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_notifications_users_user_notifications" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
