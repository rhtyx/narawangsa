package postgres

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO "users" (
  "name", "username", "email", "password"
) VALUES (
  $1, $2, $3, $4
)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	query := q.getQueries(ctx)
	_, err := query.ExecContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "username" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	query := q.getQueries(ctx)
	_, err := query.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT "id", "name", "username", "email", "password", "created_at", "updated_at" FROM "users" 
WHERE "username" = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	query := q.getQueries(ctx)
	row := query.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePasswordUser = `-- name: UpdatePasswordUser :exec
UPDATE "users"
SET "password" = $1,
    "updated_at" = $2
WHERE "username" = $3
`

type UpdatePasswordUserParams struct {
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
	Username    string    `json:"username"`
}

func (q *Queries) UpdatePasswordUser(ctx context.Context, arg UpdatePasswordUserParams) error {
	query := q.getQueries(ctx)
	_, err := query.ExecContext(ctx, updatePasswordUser, arg.NewPassword, arg.UpdatedAt, arg.Username)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "users"
SET "name" = $1,
    "email" = $2,
    "updated_at" = $3
WHERE "username" = $4
RETURNING "name", "email"
`

type UpdateUserParams struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
}

type UpdateUserRow struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	query := q.getQueries(ctx)
	row := query.QueryRowContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.UpdatedAt,
		arg.Username,
	)
	var i UpdateUserRow
	err := row.Scan(&i.Name, &i.Email)
	return i, err
}

func (q *Queries) getQueries(ctx context.Context) IQueries {
	queries, ok := GetQueryCtx(ctx)
	if !ok {
		queries = q.db
	}
	return queries
}
