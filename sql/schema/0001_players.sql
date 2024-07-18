-- +goose Up

CREATE TABLE "players" (
  "id" serial PRIMARY KEY,
  "name" text NOT NULL,
  "virtual_id" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

-- +goose Down
DROP TABLE players;