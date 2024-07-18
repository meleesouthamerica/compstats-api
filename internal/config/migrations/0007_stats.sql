-- +goose Up
CREATE TABLE stats (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  player_id integer NOT NULL,
  half_id integer NOT NULL,
  kills integer NOT NULL,
  deaths integer NOT NULL,
  assists integer NOT NULL,
  score integer NOT NULL,
  winner boolean,
  created_at datetime NOT NULL,
  updated_at datetime,
  FOREIGN KEY (player_id) REFERENCES players(id),
  FOREIGN KEY (half_id) REFERENCES halfs(id)
);

-- +goose Down
DROP TABLE stats;