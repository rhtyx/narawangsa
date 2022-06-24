package postgres

import (
	"context"
)

const createBookCategory = `-- name: CreateBookCategory :exec
INSERT INTO "category_books" (
  "book_id", "category_id"
) VALUES (
  $1, $2
)
`

type CreateBookCategoryParams struct {
	BookID     int64 `json:"book_id"`
	CategoryID int64 `json:"category_id"`
}

func (q *Queries) CreateBookCategory(ctx context.Context, arg CreateBookCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createBookCategory, arg.BookID, arg.CategoryID)
	return err
}

const deleteBookCategory = `-- name: DeleteBookCategory :exec
DELETE FROM "category_books"
WHERE "book_id" = $1 AND "category_id" = $2
`

type DeleteBookCategoryParams struct {
	BookID     int64 `json:"book_id"`
	CategoryID int64 `json:"category_id"`
}

func (q *Queries) DeleteBookCategory(ctx context.Context, arg DeleteBookCategoryParams) error {
	_, err := q.db.ExecContext(ctx, deleteBookCategory, arg.BookID, arg.CategoryID)
	return err
}
