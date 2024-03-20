CREATE EXTENSION postgis;

ALTER TABLE "address"
ADD COLUMN "coordinates" geography(POINT, 4326);

CREATE INDEX ON "location" USING GIST("coordinates");

UPDATE "address"
SET "coordinates" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326);
