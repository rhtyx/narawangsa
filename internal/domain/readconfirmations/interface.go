package readconfirmations

import (
	"context"

	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type IReadConfirmations interface {
	CreateReadConfirmation(ctx context.Context, arg postgres.CreateReadConfirmationParams) error
	ListReadConfirmations(ctx context.Context, arg postgres.ListReadConfirmationsParams) ([]postgres.ListReadConfirmationsRow, error)
}
