-- +goose Up
CREATE TABLE "teams" (
  "id" serial PRIMARY KEY,
  "name" text NOT NULL UNIQUE,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

-- +goose Down
DROP TABLE teams;