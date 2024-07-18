-- +goose Up
CREATE TABLE halfs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  map_name text NOT NULL,
  attacker_team text NOT NULL,
  defender_team text NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime
);

-- +goose Down
DROP TABLE halfs;