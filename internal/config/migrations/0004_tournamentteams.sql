-- +goose Up
CREATE TABLE tournamentteams (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  team_id integer NOT NULL,
  tournament_id integer NOT NULL,
  is_winner boolean,
  created_at datetime NOT NULL,
  updated_at datetime,
  FOREIGN KEY (team_id) REFERENCES teams(id),
  FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

-- +goose Down
DROP TABLE tournamentteams;