ALTER TABLE "location_schedule"
ALTER COLUMN "location_id" SET NOT NULL;

ALTER TABLE "location_schedule"
ALTER COLUMN "available" SET NOT NULL;

ALTER TABLE "country"
ALTER COLUMN "name" SET NOT NULL;

ALTER TABLE "location"
ALTER COLUMN "deleted" SET NOT NULL;

ALTER TABLE "location"
ALTER COLUMN "status" SET NOT NULL;
