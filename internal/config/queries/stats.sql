-- name: CreateStats :one
INSERT INTO stats (
  player_id,
  half_id,
  kills,
  deaths,
  assists,
  score,
  winner
) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING *;