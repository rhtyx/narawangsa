package userlevels

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type IUserLevels interface {
	GetUserLevel(ctx context.Context, userID int64) (postgres.GetUserLevelRow, error)
	UpdateUserLevelsLevel(ctx context.Context, arg postgres.UpdateUserLevelsLevelParams) error
	CreateUserLevel(ctx context.Context, userID int64) error
}
