-- +goose Up
CREATE TABLE teamplayers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  player_id integer NOT NULL,
  team_id integer NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (player_id) REFERENCES players(id),
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

-- +goose Down
DROP TABLE teamplayers;