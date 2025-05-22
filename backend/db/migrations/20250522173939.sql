-- Modify "tasks" table
ALTER TABLE "public"."tasks" ADD COLUMN "active" boolean NULL DEFAULT true;
