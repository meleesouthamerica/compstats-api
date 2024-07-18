-- +goose Up
CREATE TABLE "tournamentteams" (
  "id" serial PRIMARY KEY,
  "team_id" integer NOT NULL,
  "tournament_id" integer NOT NULL,
  "is_winner" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

ALTER TABLE "tournamentteams" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

ALTER TABLE "tournamentteams" ADD FOREIGN KEY ("tournament_id") REFERENCES "tournaments" ("id");

-- +goose Down
DROP TABLE tournamentteams;