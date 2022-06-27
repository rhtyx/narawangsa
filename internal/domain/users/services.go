package users

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type userStorage interface {
	CreateUser(ctx context.Context, arg postgres.CreateUserParams) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (postgres.User, error)
	UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error
	UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.UpdateUserRow, error)
}

type service struct {
	repository IUsers
	tx         postgres.TxInContext
}

func NewUserService(repository IUsers, tx postgres.TxInContext) userStorage {
	return &service{
		repository: repository,
		tx:         tx,
	}
}

func (s *service) CreateUser(ctx context.Context, arg postgres.CreateUserParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		userId, err := s.repository.CreateUser(ctx, arg)
		if err != nil {
			return err
		}

		err = s.repository.CreateUserLevel(ctx, userId)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(ctx context.Context, username string) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		err := s.repository.DeleteUser(ctx, username)
		if err != nil {
			return err
		}
		return nil
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

func (s *service) UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error {
	err := s.tx.Run(ctx, func(ctx context.Context) error {
		err := s.repository.UpdatePasswordUser(ctx, arg)
		if err != nil {
			return err
		}
		return nil
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
