package postgres

import (
	"context"
	"time"
)

const createUserLevel = `-- name: CreateUserLevel :exec
INSERT INTO "user_levels" (
  "user_id"
) VALUES (
  $1
)
`

func (q *Queries) CreateUserLevel(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, createUserLevel, userID)
	return err
}

const getUserLevel = `-- name: GetUserLevel :one
SELECT "user_id", "level" FROM "user_levels"
WHERE "user_id" = $1
`

type GetUserLevelRow struct {
	UserID int64 `json:"user_id"`
	Level  int32 `json:"level"`
}

func (q *Queries) GetUserLevel(ctx context.Context, userID int64) (GetUserLevelRow, error) {
	row := q.db.QueryRowContext(ctx, getUserLevel, userID)
	var i GetUserLevelRow
	err := row.Scan(&i.UserID, &i.Level)
	return i, err
}

const updateUserLevelsLevel = `-- name: UpdateUserLevelsLevel :exec
UPDATE "user_levels"
SET "level" = $2,
    "updated_at" = $3
WHERE "user_id" = $1
`

type UpdateUserLevelsLevelParams struct {
	UserID    int64     `json:"user_id"`
	Level     int32     `json:"level"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) UpdateUserLevelsLevel(ctx context.Context, arg UpdateUserLevelsLevelParams) error {
	_, err := q.db.ExecContext(ctx, updateUserLevelsLevel, arg.UserID, arg.Level, arg.UpdatedAt)
	return err
}
