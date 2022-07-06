package users

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type userStorage interface {
	CreateUser(ctx context.Context, arg postgres.CreateUserParams) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (postgres.User, error)
	GetUserById(ctx context.Context, id int64) (postgres.User, error)
	UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error
	UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.UpdateUserRow, error)
}

type service struct {
	repository userStorage
	tx         storage.ExecTx
}

func NewUserService(repository userStorage, tx storage.ExecTx) IUsers {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateUser(ctx context.Context, arg postgres.CreateUserParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.CreateUser(ctx, arg)
	})
	return err
}

func (s *service) DeleteUser(ctx context.Context, username string) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.DeleteUser(ctx, username)
	})
	return err
}

func (s *service) GetUser(ctx context.Context, username string) (postgres.User, error) {
	var user postgres.User
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		u, err := s.repository.GetUser(ctx, username)
		if err != nil {
			return err
		}
		user = u
		return nil
	})
	return user, err
}

func (s *service) GetUserById(ctx context.Context, id int64) (postgres.User, error) {
	var user postgres.User
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		u, err := s.repository.GetUserById(ctx, id)
		if err != nil {
			return err
		}
		user = u
		return nil
	})
	return user, err
}

func (s *service) UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		return s.repository.UpdatePasswordUser(ctx, arg)
	})
	return err
}

func (s *service) UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.UpdateUserRow, error) {
	var updatedUser postgres.UpdateUserRow
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		uu, err := s.repository.UpdateUser(ctx, arg)
		if err != nil {
			return err
		}
		updatedUser = uu
		return nil
	})
	return updatedUser, err
}
