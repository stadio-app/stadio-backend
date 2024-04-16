ALTER TABLE "location"
ADD COLUMN "price" BIGINT NOT NULL DEFAULT 0;

ALTER TABLE "location"
ADD COLUMN "currency" VARCHAR(3) REFERENCES "currency"("currency_code") NOT NULL DEFAULT 'USD';
