package postgres

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (
  "name", "username", "email", "password"
) VALUES (
  $1, $2, $3, $4
) RETURNING "id"
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "username" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT "id", "name", "username", "email", "password", "created_at", "updated_at" FROM "users" 
WHERE "username" = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
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
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
}

func (q *Queries) UpdatePasswordUser(ctx context.Context, arg UpdatePasswordUserParams) error {
	_, err := q.db.ExecContext(ctx, updatePasswordUser, arg.Password, arg.UpdatedAt, arg.Username)
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
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.UpdatedAt,
		arg.Username,
	)
	var i UpdateUserRow
	err := row.Scan(&i.Name, &i.Email)
	return i, err
}