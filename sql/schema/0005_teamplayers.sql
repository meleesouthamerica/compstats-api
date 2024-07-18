-- +goose Up
CREATE TABLE "teamplayers" (
  "id" serial PRIMARY KEY,
  "player_id" integer NOT NULL,
  "team_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

ALTER TABLE "teamplayers" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id");

ALTER TABLE "teamplayers" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

-- +goose Down
DROP TABLE teamplayers;