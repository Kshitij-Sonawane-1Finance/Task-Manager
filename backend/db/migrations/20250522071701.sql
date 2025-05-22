-- Create "admins" table
CREATE TABLE "public"."admins" (
  "id" bigserial NOT NULL,
  "email" text NULL,
  "password" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_admins_email" UNIQUE ("email")
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "password" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email")
);
-- Create "tasks" table
CREATE TABLE "public"."tasks" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "title" text NULL,
  "description" text NULL,
  "start_time" timestamptz NULL,
  "end_time" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tasks_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_tasks_user_id" to table: "tasks"
CREATE INDEX "idx_tasks_user_id" ON "public"."tasks" ("user_id");
