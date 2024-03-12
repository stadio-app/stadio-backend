ALTER TABLE "location"
RENAME COLUMN created_by TO created_by_id;
ALTER TABLE "location"
RENAME COLUMN updated_by TO updated_by_id;

ALTER TABLE "address"
RENAME COLUMN created_by TO created_by_id;
ALTER TABLE "address"
RENAME COLUMN updated_by TO updated_by_id;
