package postgres

import (
	"context"
	"time"
)

const createBookList = `-- name: CreateBookList :exec
INSERT INTO "book_lists" (
  "user_id", "book_id", "is_read", "pages_read", "end_date"
) VALUES (
  $1, $2, $3, $4, $5
)
`

type CreateBookListParams struct {
	UserID    int64     `json:"user_id"`
	BookID    int64     `json:"book_id"`
	IsRead    bool      `json:"is_read"`
	PagesRead int32     `json:"pages_read"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) CreateBookList(ctx context.Context, arg CreateBookListParams) error {
	_, err := q.db.ExecContext(ctx, createBookList,
		arg.UserID,
		arg.BookID,
		arg.IsRead,
		arg.PagesRead,
		arg.EndDate,
	)
	return err
}

const deleteBookList = `-- name: DeleteBookList :exec
DELETE FROM "book_lists"
WHERE "book_id" = $1
`

func (q *Queries) DeleteBookList(ctx context.Context, bookID int64) error {
	_, err := q.db.ExecContext(ctx, deleteBookList, bookID)
	return err
}

const listBookList = `-- name: ListBookList :many
SELECT id, user_id, book_id, is_read, pages_read, end_date, created_at, updated_at from "book_lists"
WHERE "user_id" = $1
ORDER BY "book_id"
`

func (q *Queries) ListBookList(ctx context.Context, userID int64) ([]BookList, error) {
	rows, err := q.db.QueryContext(ctx, listBookList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookList
	for rows.Next() {
		var i BookList
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.BookID,
			&i.IsRead,
			&i.PagesRead,
			&i.EndDate,
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

const updateBookList = `-- name: UpdateBookList :exec
UPDATE "book_lists"
SET "is_read" = $1,
    "pages_read" = $2,
    "end_date" = $3,
    "updated_at" = $4
WHERE "user_id" = $5 AND "book_id" = $6
`

type UpdateBookListParams struct {
	IsRead    bool      `json:"is_read"`
	PagesRead int32     `json:"pages_read"`
	EndDate   time.Time `json:"end_date"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int64     `json:"user_id"`
	BookID    int64     `json:"book_id"`
}

func (q *Queries) UpdateBookList(ctx context.Context, arg UpdateBookListParams) error {
	_, err := q.db.ExecContext(ctx, updateBookList,
		arg.IsRead,
		arg.PagesRead,
		arg.EndDate,
		arg.UpdatedAt,
		arg.UserID,
		arg.BookID,
	)
	return err
}
