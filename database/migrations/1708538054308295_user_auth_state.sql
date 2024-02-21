CREATE TABLE "auth_state" (
    "id" BIGSERIAL UNIQUE PRIMARY KEY,
    "logged_in_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "user_id" BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL,
    "ip_address" VARCHAR(39)
);

CREATE TABLE "email_verification" (
    "id" BIGSERIAL UNIQUE PRIMARY KEY,
    "code" VARCHAR(20) NOT NULL,
    "user_id" BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL, 
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
