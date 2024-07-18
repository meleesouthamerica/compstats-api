-- +goose Up
CREATE TABLE tournaments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL UNIQUE,
  created_at datetime NOT NULL,
  updated_at datetime
);

-- +goose Down
DROP TABLE tournaments;