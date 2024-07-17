-- Copyright (c) 2024 Trenova Technologies, LLC
--
-- Licensed under the Business Source License 1.1 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     https://trenova.app/pricing/
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
--
-- Key Terms:
-- - Non-production use only
-- - Change Date: 2026-11-16
-- - Change License: GNU General Public License v2 or later
--
-- For full license text, see the LICENSE file in the root directory.

DROP TABLE IF EXISTS "locations" CASCADE;

-- bun:split

DROP TABLE IF EXISTS "location_comments" CASCADE;

-- bun:split

DROP TABLE IF EXISTS "location_contacts" CASCADE;

-- bun:split

DROP INDEX IF EXISTS "locations_code_organization_id_unq" CASCADE;
DROP INDEX IF EXISTS "idx_location_name" CASCADE;
DROP INDEX IF EXISTS "idx_location_org_bu" CASCADE;
DROP INDEX IF EXISTS "idx_location_description" CASCADE;
DROP INDEX IF EXISTS "idx_location_created_at" CASCADE;
DROP INDEX IF EXISTS "idx_location_comment_comment_type_id" CASCADE;
DROP INDEX IF EXISTS "idx_location_comment_created_at" CASCADE;
DROP INDEX IF EXISTS "idx_location_comment_created_at" CASCADE;
DROP INDEX IF EXISTS "idx_location_contact_name" CASCADE;
DROP INDEX IF EXISTS "idx_location_contact_created_at" CASCADE;
