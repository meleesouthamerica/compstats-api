-- name: CreatePlayer :one
INSERT INTO players (
  name,
  virtual_id,
  created_at,
  updated_at
) VALUES (?, ?, ?, ?) RETURNING *;

-- name: FindPlayerByVirtualID :one
SELECT *
FROM players
WHERE virtual_id = ?
LIMIT 1;

-- name: FindPlayerByID :one
SELECT *
FROM players
WHERE id = ?
LIMIT 1;
