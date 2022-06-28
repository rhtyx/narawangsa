package categories

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type ICategories interface {
	CreateCategory(ctx context.Context, name string) error
	DeleteCategory(ctx context.Context, name string) error
	GetCategory(ctx context.Context, name string) (postgres.Category, error)
	ListCategories(ctx context.Context, arg postgres.ListCategoriesParams) ([]postgres.Category, error)
	UpdateCategory(ctx context.Context, arg postgres.UpdateCategoryParams) error
}
