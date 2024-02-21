-- Create "user" table
CREATE TABLE "user" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "email" VARCHAR(255) UNIQUE NOT NULL, 
  "phone_number" VARCHAR(15) UNIQUE NOT NULL, 
  "name" VARCHAR(255) NOT NULL, 
  "avatar" TEXT NULL, 
  "birth_date" DATE NULL, 
  "bio" TEXT NULL, 
  "active" BOOLEAN NOT NULL DEFAULT FALSE 
);

-- Create "owner" table
CREATE TABLE "owner" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "first_name" VARCHAR(255) NOT NULL, 
  "middle_name" VARCHAR(255) NULL, 
  "last_name" VARCHAR(255) NOT NULL, 
  "full_name" VARCHAR(255) NOT NULL, 
  "verified" BOOLEAN NOT NULL DEFAULT FALSE, 
  "user_id" BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL
);

-- Create "address" table
CREATE TYPE "country_code_alpha_2" AS ENUM (
    'AF','AX','AL','DZ','AS','AD','AO','AI','AQ','AG','AR','AM','AW','AU','AT','AZ','BS','BH','BD','BB','BY','BE','BZ','BJ','BM','BT','BO','BA','BW','BV','BR','IO','BN','BG','BF','BI','KH','CM','CA','CV','KY','CF','TD','CL','CN','CX','CC','CO','KM','CG','CD','CK','CR','CI','HR','CU','CY','CZ','DK','DJ','DM','DO','EC','EG','SV','GQ','ER','EE','ET','FK','FO','FJ','FI','FR','GF','PF','TF','GA','GM','GE','DE','GH','GI','GR','GL','GD','GP','GU','GT','GG','GN','GW','GY','HT','HM','VA','HN','HK','HU','IS','IN','ID','IR','IQ','IE','IM','IL','IT','JM','JP','JE','JO','KZ','KE','KI','KR','KP','KW','KG','LA','LV','LB','LS','LR','LY','LI','LT','LU','MO','MK','MG','MW','MY','MV','ML','MT','MH','MQ','MR','MU','YT','MX','FM','MD','MC','MN','ME','MS','MA','MZ','MM','NA','NR','NP','NL','AN','NC','NZ','NI','NE','NG','NU','NF','MP','NO','OM','PK','PW','PS','PA','PG','PY','PE','PH','PN','PL','PT','PR','QA','RE','RO','RU','RW','BL','SH','KN','LC','MF','PM','VC','WS','SM','ST','SA','SN','RS','SC','SL','SG','SK','SI','SB','SO','ZA','GS','ES','LK','SD','SR','SJ','SZ','SE','CH','SY','TW','TJ','TZ','TH','TL','TG','TK','TO','TT','TN','TR','TM','TC','TV','UG','UA','AE','GB','US','UM','UY','UZ','VU','VE','VN','VG','VI','WF','EH','YE','ZM','ZW'
);
CREATE TABLE "address" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "latitude" double precision NOT NULL, 
  "longitude" double precision NOT NULL, 
  "maps_link" TEXT NOT NULL, 
  "full_address" VARCHAR(255) NOT NULL, 
  "country_code" "country_code_alpha_2" NOT NULL DEFAULT 'US',
  "country" VARCHAR(56) NOT NULL DEFAULT 'United States',
  "created_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL,
  "updated_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL
);


-- Create "location" table
CREATE TYPE "location_status" AS ENUM (
    'CLOSED',
    'MOVED',
    'OPERATIONAL',
    'TEMPORARILY_CLOSED'
);
CREATE TABLE "location" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "name" VARCHAR(255) NOT NULL, 
  "description" TEXT, 
  "type" VARCHAR(255) NOT NULL, 
  "owner_id" BIGINT REFERENCES "owner"("id") ON DELETE SET NULL, 
  "address_id" BIGINT REFERENCES "address"("id") ON DELETE CASCADE NOT NULL, 
  "deleted" BOOLEAN DEFAULT FALSE, 
  "closed" "location_status" DEFAULT 'OPERATIONAL', 
  "created_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL, 
  "updated_by" BIGINT REFERENCES "user"("id") ON DELETE SET NULL
);

-- Create "event" table
CREATE TABLE "event" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "name" VARCHAR(255) NOT NULL, 
  "description" TEXT, 
  "type" VARCHAR(255) NULL, 
  "start_date" TIMESTAMPTZ NOT NULL, 
  "end_date" TIMESTAMPTZ NOT NULL, 
  "location_id" BIGINT REFERENCES "location"("id") ON DELETE SET NULL
);

-- Create "participant" table
CREATE TABLE "participant" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "nickname" VARCHAR(255) NULL, 
  "admin" BOOLEAN NOT NULL DEFAULT false, 
  "participates" BOOLEAN NOT NULL DEFAULT true, 
  "skill_level" VARCHAR(255) NULL, 
  "event_id" BIGINT REFERENCES "event"("id") ON DELETE CASCADE NOT NULL, 
  "user_id" BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL
);

-- Create "review" table
CREATE TABLE "review" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY, 
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  "rating" double precision NOT NULL, 
  "message" TEXT, 
  "location_id" BIGINT REFERENCES "location"("id") ON DELETE CASCADE, 
  "event_id" BIGINT REFERENCES "event"("id") ON DELETE CASCADE NOT NULL, 
  "participant_id" BIGINT REFERENCES "participant"("id") ON DELETE CASCADE NOT NULL
);
