-- +goose Up
CREATE TABLE stats (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  player_id integer NOT NULL,
  half_id integer NOT NULL,
  kills integer NOT NULL,
  deaths integer NOT NULL,
  assists integer NOT NULL,
  score integer NOT NULL,
  winner boolean NOT NULL DEFAULT FALSE,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (player_id) REFERENCES players(id),
  FOREIGN KEY (half_id) REFERENCES halfs(id)
);

-- +goose Down
DROP TABLE stats;