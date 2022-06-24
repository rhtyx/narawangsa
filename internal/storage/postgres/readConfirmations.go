package postgres

import (
	"context"
)

const createReadConfirmation = `-- name: CreateReadConfirmation :exec
INSERT INTO "read_confirmations" (
  "book_list_id", "pages_read"
) VALUES (
  $1, $2
)
`

type CreateReadConfirmationParams struct {
	BookListID int64 `json:"book_list_id"`
	PagesRead  int32 `json:"pages_read"`
}

func (q *Queries) CreateReadConfirmation(ctx context.Context, arg CreateReadConfirmationParams) error {
	_, err := q.db.ExecContext(ctx, createReadConfirmation, arg.BookListID, arg.PagesRead)
	return err
}

const listReadConfirmations = `-- name: ListReadConfirmations :many
SELECT "book_list_id", "pages_read"
FROM "read_confirmations"
WHERE "book_list_id" = $1
LIMIT $2
`

type ListReadConfirmationsParams struct {
	BookListID int64 `json:"book_list_id"`
	Limit      int32 `json:"limit"`
}

type ListReadConfirmationsRow struct {
	BookListID int64 `json:"book_list_id"`
	PagesRead  int32 `json:"pages_read"`
}

func (q *Queries) ListReadConfirmations(ctx context.Context, arg ListReadConfirmationsParams) ([]ListReadConfirmationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listReadConfirmations, arg.BookListID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListReadConfirmationsRow
	for rows.Next() {
		var i ListReadConfirmationsRow
		if err := rows.Scan(&i.BookListID, &i.PagesRead); err != nil {
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
