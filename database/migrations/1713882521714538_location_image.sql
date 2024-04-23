CREATE TABLE "location_image" (
    "id" BIGSERIAL UNIQUE PRIMARY KEY, 
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "upload_id" TEXT NOT NULL UNIQUE,
    "original_filename" VARCHAR(255) NOT NULL,
    "location_id" BIGINT REFERENCES "location"("id") ON DELETE CASCADE,
    "caption" TEXT,
    "created_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL, 
    "updated_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL
);
