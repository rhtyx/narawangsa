package postgres

import (
	"context"
	"time"
)

const createBook = `-- name: CreateBook :exec
INSERT INTO "books" (
  "title", "author", "year", "pages", "synopsis"
) VALUES (
  $1, $2, $3, $4, $5
)
`

type CreateBookParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     string `json:"year"`
	Pages    int32  `json:"pages"`
	Synopsis string `json:"synopsis"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) error {
	_, err := q.db.ExecContext(ctx, createBook,
		arg.Title,
		arg.Author,
		arg.Year,
		arg.Pages,
		arg.Synopsis,
	)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM "books"
WHERE "id" = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author, year, pages, synopsis, created_at, updated_at FROM "books"
WHERE "title" = $1
`

func (q *Queries) GetBook(ctx context.Context, title string) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, title)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Year,
		&i.Pages,
		&i.Synopsis,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, author, year, pages, synopsis, created_at, updated_at FROM "books"
ORDER BY "title"
LIMIT $1
`

func (q *Queries) ListBooks(ctx context.Context, limit int32) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Pages,
			&i.Synopsis,
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

const updateBook = `-- name: UpdateBook :exec
UPDATE "books"
SET "title" = $1,
    "author" = $2,
    "year" = $3,
    "pages" = $4,
    "synopsis" = $5,
    "updated_at" = $6
WHERE "id" = $7
`

type UpdateBookParams struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      string    `json:"year"`
	Pages     int32     `json:"pages"`
	Synopsis  string    `json:"synopsis"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        int64     `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook,
		arg.Title,
		arg.Author,
		arg.Year,
		arg.Pages,
		arg.Synopsis,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
