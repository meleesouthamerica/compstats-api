-- +goose Up
CREATE TABLE "halfs" (
  "id" serial PRIMARY KEY,
  "map_name" text NOT NULL,
  "attacker_team" text NOT NULL,
  "defender_team" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

-- +goose Down
DROP TABLE halfs;