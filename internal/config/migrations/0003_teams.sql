-- +goose Up
CREATE TABLE teams (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL UNIQUE,
  created_at datetime NOT NULL,
  updated_at datetime
);

-- +goose Down
DROP TABLE teams;