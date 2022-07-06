package postgres

import (
	"context"
)

const createRefreshToken = `-- name: CreateRefreshToken :exec
INSERT INTO "authentications" (
  "refresh_token"
) VALUES (
  $1
)
`

func (q *Queries) CreateRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := q.db.ExecContext(ctx, createRefreshToken, refreshToken)
	return err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
DELETE FROM "authentications"
WHERE "refresh_token" = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshToken, refreshToken)
	return err
}

const getRefreshToken = `-- name: GetRefreshToken :one
SELECT "refresh_token" FROM "authentications"
WHERE "refresh_token" = $1
`

func (q *Queries) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	row := q.db.QueryRowContext(ctx, getRefreshToken, refreshToken)
	var refresh_token string
	err := row.Scan(&refresh_token)
	return refresh_token, err
}
