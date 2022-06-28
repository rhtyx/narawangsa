package users

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type IUsers interface {
	CreateUser(ctx context.Context, arg postgres.CreateUserParams) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (postgres.User, error)
	UpdatePasswordUser(ctx context.Context, arg postgres.UpdatePasswordUserParams) error
	UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.UpdateUserRow, error)
}
