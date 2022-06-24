package storage

import "context"

type ExecTx interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}
