-- name: GetAllTournaments :many
SELECT * FROM tournaments
ORDER BY id DESC;

-- name: GetTournamentByID :one
SELECT * FROM tournaments
WHERE id = ?;

-- name: GetTournamentByName :one
SELECT * FROM tournaments
WHERE name = ?;

-- name: CreateTournament :one
INSERT INTO tournaments (
  name
) VALUES (?) RETURNING *;

-- name: UpdateTournamentByID :one
UPDATE tournaments
SET name = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteTournamentByID :exec
DELETE FROM tournaments WHERE id = ?;