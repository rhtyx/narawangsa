package postgres

import (
	"context"
	"time"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO "categories" (
  "name"
) VALUES (
  $1
)
`

func (q *Queries) CreateCategory(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createCategory, name)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM "categories"
WHERE "id" = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, created_at, updated_at FROM "categories"
WHERE "name" = $1
`

func (q *Queries) GetCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, created_at, updated_at FROM "categories"
ORDER BY "name"
LIMIT $1
OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE "categories"
SET "name" = $1,
    "updated_at" = $2
WHERE "id" = $3
`

type UpdateCategoryParams struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        int64     `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.Name, arg.UpdatedAt, arg.ID)
	return err
}
