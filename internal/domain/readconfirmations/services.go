package readconfirmations

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type readConfirmationsStorage interface {
	CreateReadConfirmation(ctx context.Context, arg postgres.CreateReadConfirmationParams) error
	ListReadConfirmations(ctx context.Context, arg postgres.ListReadConfirmationsParams) ([]postgres.ReadConfirmation, error)
}

type service struct {
	repository readConfirmationsStorage
	tx         storage.ExecTx
}

func NewReadConfirmationsService(repository readConfirmationsStorage, tx storage.ExecTx) IReadConfirmations {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateReadConfirmation(ctx context.Context, arg postgres.CreateReadConfirmationParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateReadConfirmation(ctx, arg)
	})
	return err
}
func (s *service) ListReadConfirmations(ctx context.Context, arg postgres.ListReadConfirmationsParams) ([]postgres.ReadConfirmation, error) {
	var readConfirmations []postgres.ReadConfirmation
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		rc, err := s.repository.ListReadConfirmations(ctx, arg)
		if err != nil {
			return err
		}
		readConfirmations = rc
		return nil
	})
	return readConfirmations, err
}
