CREATE EXTENSION IF NOT EXISTS postgis;

ALTER TABLE "address"
ADD COLUMN "coordinates" geography(POINT, 4326) NOT NULL
    GENERATED ALWAYS AS (
        ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)
    ) STORED;

CREATE INDEX "address_country_code_idx" ON "address"("country_code");
CREATE INDEX "address_coordinates_idx" ON "address" USING GIST("coordinates");
