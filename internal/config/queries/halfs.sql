-- name: FindHalfByID :one
SELECT *
FROM halfs
WHERE id = ?;

-- name: CreateHalf :one
INSERT INTO halfs (
  map_name,
  team_1,
  team_2,
  tournament_id
) VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetAllHalfs :many
SELECT * FROM halfs
ORDER BY id DESC;

-- name: UpdateHalfByID :one
UPDATE halfs
SET
  map_name = COALESCE(sqlc.narg(map_name), map_name),
  team_1 = COALESCE(sqlc.narg(team_1), team_1),
  team_2 = COALESCE(sqlc.narg(team_2), team_2),
  updated_at = CURRENT_TIMESTAMP
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteHalfByID :exec
DELETE FROM halfs WHERE id = ?;