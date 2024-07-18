-- +goose Up
CREATE TABLE "stats" (
  "id" serial PRIMARY KEY,
  "player_id" integer NOT NULL,
  "half_id" integer NOT NULL,
  "kills" integer NOT NULL,
  "deaths" integer NOT NULL,
  "assists" integer NOT NULL,
  "score" integer NOT NULL,
  "winner" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

ALTER TABLE "stats" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id");

ALTER TABLE "stats" ADD FOREIGN KEY ("half_id") REFERENCES "halfs" ("id");

-- +goose Down
DROP TABLE stats;