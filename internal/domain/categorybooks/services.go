package categorybooks

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type categoryBooksStorage interface {
	CreateBookCategory(ctx context.Context, arg postgres.CreateBookCategoryParams) error
	DeleteBookCategory(ctx context.Context, arg postgres.DeleteBookCategoryParams) error
}

type service struct {
	repository categoryBooksStorage
	tx         storage.ExecTx
}

func NewCategoryBooksService(repository categoryBooksStorage, tx storage.ExecTx) ICategoryBooks {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateBookCategory(ctx context.Context, arg postgres.CreateBookCategoryParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateBookCategory(ctx, arg)
	})
	return err
}
func (s *service) DeleteBookCategory(ctx context.Context, arg postgres.DeleteBookCategoryParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteBookCategory(ctx, arg)
	})
	return err
}
