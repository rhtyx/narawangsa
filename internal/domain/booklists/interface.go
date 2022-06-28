package booklists

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type IBooklists interface {
	CreateBookList(ctx context.Context, arg postgres.CreateBookListParams) error
	DeleteBookList(ctx context.Context, bookID int64) error
	ListBookList(ctx context.Context, userID int64) ([]postgres.BookList, error)
	UpdateBookList(ctx context.Context, arg postgres.UpdateBookListParams) error
}
