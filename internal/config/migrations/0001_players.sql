-- +goose Up

CREATE TABLE players (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  virtual_id text NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime
);

-- +goose Down
DROP TABLE players;