-- +goose Up

CREATE TABLE players (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  virtual_id text NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE players;