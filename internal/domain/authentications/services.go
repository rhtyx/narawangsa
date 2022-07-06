package authentications

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage"
)

type authenticationsStorage interface {
	CreateRefreshToken(ctx context.Context, refreshToken string) error
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type service struct {
	repository authenticationsStorage
	tx         storage.ExecTx
}

func NewAuthenticationsService(repository authenticationsStorage, tx storage.ExecTx) IAuthentications {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateRefreshToken(ctx context.Context, refreshToken string) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateRefreshToken(ctx, refreshToken)
	})
	return err
}

func (s *service) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteRefreshToken(ctx, refreshToken)
	})
	return err
}
func (s *service) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	var token string
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		t, err := s.repository.GetRefreshToken(ctx, refreshToken)
		if err != nil {
			return err
		}
		token = t
		return nil
	})
	return token, err
}
