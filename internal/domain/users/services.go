package users

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type Service struct {
	repository IUsers
}

func NewUserService(repository IUsers) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUser(ctx context.Context, arg postgres.CreateUserParams) error {
	userId, err := s.repository.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	err = s.repository.CreateUserLevel(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, username string) error {
	return s.repository.DeleteUser(ctx, username)
}

func (s *Service) GetUser(ctx context.Context, username string) (postgres.User, error) {
	return s.repository.GetUser(ctx, username)
}

func (s *Service) UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error {
	return s.repository.UpdatePasswordUser(ctx, arg)
}

func (s *Service) UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.UpdateUserRow, error) {
	return s.repository.UpdateUser(ctx, arg)
}
