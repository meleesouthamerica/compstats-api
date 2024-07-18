// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: players.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
  name,
  virtual_id,
  created_at,
  updated_at
) VALUES (?, ?, ?, ?) RETURNING id, name, virtual_id, created_at, updated_at
`

type CreatePlayerParams struct {
	Name      string       `json:"name"`
	VirtualID string       `json:"virtualId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, createPlayer,
		arg.Name,
		arg.VirtualID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.VirtualID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPlayerByID = `-- name: FindPlayerByID :one
SELECT id, name, virtual_id, created_at, updated_at
FROM players
WHERE id = ?
LIMIT 1
`

func (q *Queries) FindPlayerByID(ctx context.Context, id int64) (Player, error) {
	row := q.db.QueryRowContext(ctx, findPlayerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.VirtualID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPlayerByVirtualID = `-- name: FindPlayerByVirtualID :one
SELECT id, name, virtual_id, created_at, updated_at
FROM players
WHERE virtual_id = ?
LIMIT 1
`

func (q *Queries) FindPlayerByVirtualID(ctx context.Context, virtualID string) (Player, error) {
	row := q.db.QueryRowContext(ctx, findPlayerByVirtualID, virtualID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.VirtualID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
