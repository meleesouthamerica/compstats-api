-- +goose Up
CREATE TABLE halfs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  map_name text NOT NULL,
  team_1 text NOT NULL,
  team_2 text NOT NULL,
  tournament_id integer NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

-- +goose Down
DROP TABLE halfs;