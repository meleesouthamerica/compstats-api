-- name: CreatePlayer :one
INSERT INTO players (
  name,
  virtual_id
) VALUES (?, ?) RETURNING *;

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

-- name: GetAllPlayers :many
SELECT * FROM players
ORDER BY id DESC;

-- name: UpdatePlayerByID :one
UPDATE players
SET
  name = COALESCE(sqlc.narg('name'), name),
  virtual_id = COALESCE(sqlc.narg('virtual_id'), virtual_id),
  updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePlayerByID :exec
DELETE FROM players WHERE id = ?;