CREATE TYPE "week_day" AS ENUM (
  'SUNDAY',
  'MONDAY',
  'TUESDAY',
  'WEDNESDAY',
  'THURSDAY',
  'FRIDAY',
  'SATURDAY'
);

CREATE TABLE "location_schedule" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "location_id" BIGINT REFERENCES "location"("id") ON DELETE CASCADE, 
  "day" "week_day" NOT NULL,
  "on" DATE,
  "from" INTEGER, -- 24 hour time
  "to" INTEGER, -- 24 hour time
  "available" BOOLEAN DEFAULT FALSE
);
