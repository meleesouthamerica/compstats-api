-- name: GetAllTournaments :many
SELECT * FROM tournaments
ORDER BY id DESC;

-- name: GetTournamentByID :one
SELECT * FROM tournaments
WHERE id = ?;

-- name: CreateTournament :one
INSERT INTO tournaments (
  name,
  created_at
) VALUES (?, ?) RETURNING *;

-- name: UpdateTournamentByID :one
UPDATE tournaments
SET name = ?, updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTournamentByID :exec
DELETE FROM tournaments WHERE id = ?;