-- The goal of this migration is to go from "from" and "to" column
-- to storing the "from" values as a TIME column
-- and the "to" column as a int duration column that stores the difference

ALTER TABLE "location_schedule"
RENAME COLUMN "to" TO "to_duration";

UPDATE "location_schedule"
SET "to_duration" = "to_duration" - "from"
WHERE "from" <= "to_duration";

UPDATE "location_schedule"
SET "to_duration" = ("to_duration" + 24) - "from"
WHERE "from" > "to_duration";

ALTER TABLE "location_schedule"
ALTER COLUMN "from" TYPE TIME USING CONCAT("from", ':00:00')::TIME;
