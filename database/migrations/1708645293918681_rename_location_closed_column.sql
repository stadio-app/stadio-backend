ALTER TABLE "location" 
RENAME COLUMN "closed" TO "status";

-- add created by and updated by columns
ALTER TABLE "event"
ADD COLUMN "created_by_id" BIGINT REFERENCES "user"("id") ON DELETE SET NULL,
ADD COLUMN "updated_by_id" BIGINT REFERENCES "user"("id") ON DELETE SET NULL;
