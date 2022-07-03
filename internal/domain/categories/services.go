package categories

import (
	"context"
	"strings"

	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type categoriesStorage interface {
	CreateCategory(ctx context.Context, name string) error
	DeleteCategory(ctx context.Context, id int64) error
	GetCategory(ctx context.Context, name string) (postgres.Category, error)
	ListCategories(ctx context.Context, arg postgres.ListCategoriesParams) ([]postgres.Category, error)
	UpdateCategory(ctx context.Context, arg postgres.UpdateCategoryParams) error
}

type service struct {
	repository categoriesStorage
	tx         storage.ExecTx
}

func NewCategoriesService(repository categoriesStorage, tx storage.ExecTx) ICategories {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateCategory(ctx context.Context, name string) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateCategory(ctx, strings.ToLower(name))
	})
	return err
}
func (s *service) DeleteCategory(ctx context.Context, id int64) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteCategory(ctx, id)
	})
	return err
}
func (s *service) GetCategory(ctx context.Context, name string) (postgres.Category, error) {
	var category postgres.Category
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		c, err := s.repository.GetCategory(ctx, name)
		if err != nil {
			return err
		}
		category = c
		return nil
	})
	return category, err
}
func (s *service) ListCategories(ctx context.Context, arg postgres.ListCategoriesParams) ([]postgres.Category, error) {
	var categories []postgres.Category
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		cs, err := s.repository.ListCategories(ctx, arg)
		if err != nil {
			return err
		}
		categories = cs
		return nil
	})
	return categories, err
}
func (s *service) UpdateCategory(ctx context.Context, arg postgres.UpdateCategoryParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.UpdateCategory(ctx, arg)
	})
	return err
}
