package categorybooks

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type ICategoryBooks interface {
	CreateBookCategory(ctx context.Context, arg postgres.CreateBookCategoryParams) error
	DeleteBookCategory(ctx context.Context, arg postgres.DeleteBookCategoryParams) error
}
