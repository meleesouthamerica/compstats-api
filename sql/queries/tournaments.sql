-- name: GetAllTournaments :many
SELECT * FROM tournaments
ORDER BY id DESC;

-- name: GetTournamentByID :one
SELECT * FROM tournaments
WHERE id = $1;

-- name: CreateTournament :one
INSERT INTO tournaments (
  name,
  created_at
) VALUES ($1, $2) RETURNING *;

-- name: UpdateTournamentByID :one
UPDATE tournaments
SET name = $1, updated_at = $2
WHERE id = $3
RETURNING *;

-- name: DeleteTournamentByID :exec
DELETE FROM tournaments WHERE id = $1;