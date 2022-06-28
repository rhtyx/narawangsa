package booklists

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type booklistsStorage interface {
	CreateBookList(ctx context.Context, arg postgres.CreateBookListParams) error
	DeleteBookList(ctx context.Context, bookID int64) error
	ListBookList(ctx context.Context, userID int64) ([]postgres.BookList, error)
	UpdateBookList(ctx context.Context, arg postgres.UpdateBookListParams) error
}

type service struct {
	repository booklistsStorage
	tx         postgres.TxInContext
}

func NewBookListsService(repository booklistsStorage, tx postgres.TxInContext) IBooklists {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateBookList(ctx context.Context, arg postgres.CreateBookListParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateBookList(ctx, arg)
	})
	return err
}
func (s *service) DeleteBookList(ctx context.Context, bookID int64) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteBookList(ctx, bookID)
	})
	return err
}
func (s *service) ListBookList(ctx context.Context, userID int64) ([]postgres.BookList, error) {
	var booklists []postgres.BookList
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		bl, err := s.repository.ListBookList(ctx, userID)
		if err != nil {
			return err
		}
		booklists = bl
		return nil
	})
	return booklists, err
}
func (s *service) UpdateBookList(ctx context.Context, arg postgres.UpdateBookListParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.UpdateBookList(ctx, arg)
	})
	return err
}
