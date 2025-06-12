-- Rename a column from "start_time" to "start_date"
ALTER TABLE "public"."tasks" RENAME COLUMN "start_time" TO "start_date";
-- Rename a column from "end_time" to "end_date"
ALTER TABLE "public"."tasks" RENAME COLUMN "end_time" TO "end_date";
-- Modify "tasks" table
ALTER TABLE "public"."tasks" ADD COLUMN "status" text NULL DEFAULT 'Pending', ADD COLUMN "priority" text NULL DEFAULT 'Medium';
