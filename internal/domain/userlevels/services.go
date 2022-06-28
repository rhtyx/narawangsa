package userlevels

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type userLevelsStorage interface {
	GetUserLevel(ctx context.Context, userID int64) (postgres.GetUserLevelRow, error)
	UpdateUserLevelsLevel(ctx context.Context, arg postgres.UpdateUserLevelsLevelParams) error
}

type service struct {
	repository IUserLevels
	tx         postgres.TxInContext
}

func NewUserLevelsService(repository IUserLevels, tx postgres.TxInContext) userLevelsStorage {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) GetUserLevel(ctx context.Context, userID int64) (postgres.GetUserLevelRow, error) {
	var userLevel postgres.GetUserLevelRow
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		ul, err := s.repository.GetUserLevel(ctx, userID)
		if err != nil {
			return err
		}
		userLevel = ul
		return nil
	})
	return userLevel, err
}

func (s *service) UpdateUserLevelsLevel(ctx context.Context, arg postgres.UpdateUserLevelsLevelParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.UpdateUserLevelsLevel(ctx, arg)
	})
	return err
}
