package authentications

import "context"

type IAuthentications interface {
	CreateRefreshToken(ctx context.Context, refreshToken string) error
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
}
