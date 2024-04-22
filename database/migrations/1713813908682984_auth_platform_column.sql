ALTER TABLE "user"
DROP COLUMN "auth_platform";

ALTER TABLE "auth_state"
ADD COLUMN "platform" "user_auth_platform_type" NOT NULL DEFAULT 'INTERNAL';
