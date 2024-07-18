-- name: CreatePlayer :one
INSERT INTO players (
  id,
  name,
  virtual_id,
  created_at,
  updated_at
) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: FindPlayerByVirtualID :one
SELECT *
FROM players
WHERE virtual_id = $1
LIMIT 1;

-- name: FindPlayerByID :one
SELECT *
FROM players
WHERE id = $1
LIMIT 1;
