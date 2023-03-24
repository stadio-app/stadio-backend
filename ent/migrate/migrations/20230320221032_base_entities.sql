-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "email" character varying NOT NULL, "phone_number" character varying NULL, "name" character varying NOT NULL, "avatar" text NULL, "birth_date" timestamptz NULL, "bio" text NULL, "active" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_number_key" ON "users" ("phone_number");
-- create "owners" table
CREATE TABLE "owners" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "first_name" character varying NOT NULL, "middle_name" character varying NULL, "last_name" character varying NOT NULL, "full_name" character varying NOT NULL, "id_url" character varying NOT NULL, "verified" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_owner" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "owners_users_owner" FOREIGN KEY ("user_owner") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "owners_id_url_key" to table: "owners"
CREATE UNIQUE INDEX "owners_id_url_key" ON "owners" ("id_url");
-- create index "owners_user_owner_key" to table: "owners"
CREATE UNIQUE INDEX "owners_user_owner_key" ON "owners" ("user_owner");
-- create "locations" table
CREATE TABLE "locations" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "type" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "owner_locations" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "locations_owners_locations" FOREIGN KEY ("owner_locations") REFERENCES "owners" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "addresses" table
CREATE TABLE "addresses" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "latitude" double precision NOT NULL, "longitude" double precision NOT NULL, "maps_link" text NOT NULL, "full_address" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "location_address" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "addresses_locations_address" FOREIGN KEY ("location_address") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "addresses_location_address_key" to table: "addresses"
CREATE UNIQUE INDEX "addresses_location_address_key" ON "addresses" ("location_address");
-- create "events" table
CREATE TABLE "events" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "type" character varying NULL, "start_date" timestamptz NOT NULL, "end_date" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "location_events" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "events_locations_events" FOREIGN KEY ("location_events") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "participants" table
CREATE TABLE "participants" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "nickname" character varying NULL, "admin" boolean NOT NULL DEFAULT false, "participates" boolean NOT NULL DEFAULT true, "skill_level" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "event_participants" bigint NULL, "user_participants" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "participants_events_participants" FOREIGN KEY ("event_participants") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "participants_users_participants" FOREIGN KEY ("user_participants") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "reviews" table
CREATE TABLE "reviews" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "rating" double precision NOT NULL, "message" text NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "location_reviews" bigint NULL, "participant_reviews" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "reviews_locations_reviews" FOREIGN KEY ("location_reviews") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "reviews_participants_reviews" FOREIGN KEY ("participant_reviews") REFERENCES "participants" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
