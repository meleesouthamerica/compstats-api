-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES (?, ?, ?) RETURNING *;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;
