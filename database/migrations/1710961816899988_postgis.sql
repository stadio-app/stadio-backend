CREATE EXTENSION IF NOT EXISTS postgis;

ALTER TABLE "address"
ADD COLUMN "coordinates" geography(POINT, 4326);

UPDATE "address"
SET "coordinates" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326);

ALTER TABLE "address"
ALTER COLUMN "coordinates" SET NOT NULL;
