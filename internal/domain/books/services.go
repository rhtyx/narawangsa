package books

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type booksStorage interface {
	CreateBook(ctx context.Context, arg postgres.CreateBookParams) error
	DeleteBook(ctx context.Context, id int64) error
	GetBook(ctx context.Context, title string) (postgres.Book, error)
	ListBooks(ctx context.Context, limit int32) ([]postgres.Book, error)
	UpdateBook(ctx context.Context, arg postgres.UpdateBookParams) error
}

type service struct {
	repository booksStorage
	tx         postgres.TxInContext
}

func NewBooksService(repository booksStorage, tx postgres.TxInContext) IBooks {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateBook(ctx context.Context, arg postgres.CreateBookParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateBook(ctx, arg)
	})
	return err
}
func (s *service) DeleteBook(ctx context.Context, id int64) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteBook(ctx, id)
	})
	return err
}
func (s *service) GetBook(ctx context.Context, title string) (postgres.Book, error) {
	var book postgres.Book
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		b, err := s.repository.GetBook(ctx, title)
		if err != nil {
			return err
		}
		book = b
		return nil
	})
	return book, err
}
func (s *service) ListBooks(ctx context.Context, limit int32) ([]postgres.Book, error) {
	var books []postgres.Book
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		bs, err := s.repository.ListBooks(ctx, 10)
		if err != nil {
			return err
		}
		books = bs
		return nil
	})
	return books, err
}
func (s *service) UpdateBook(ctx context.Context, arg postgres.UpdateBookParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.UpdateBook(ctx, arg)
	})
	return err
}
