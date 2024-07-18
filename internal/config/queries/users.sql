-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password,
  created_at,
  updated_at
) VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

