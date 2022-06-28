package books

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type IBooks interface {
	CreateBook(ctx context.Context, arg postgres.CreateBookParams) (int64, error)
	DeleteBook(ctx context.Context, id int64) error
	GetBook(ctx context.Context, title string) (postgres.Book, error)
	ListBooks(ctx context.Context, limit int32) ([]postgres.Book, error)
	UpdateBook(ctx context.Context, arg postgres.UpdateBookParams) error
}
