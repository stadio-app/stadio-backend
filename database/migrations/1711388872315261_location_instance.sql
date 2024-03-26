CREATE TABLE "location_instance" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "location_id" BIGINT REFERENCES "location"("id") ON DELETE CASCADE, 
  "name" VARCHAR(255)
);

INSERT INTO "location_instance" (
  "location_id"
) SELECT "id" FROM "location";

ALTER TABLE "event"
ADD COLUMN "location_instance_id" BIGINT REFERENCES "location_instance"("id") ON DELETE SET NULL;
