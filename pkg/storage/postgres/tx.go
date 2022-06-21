package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

var key int = 0

type txInContext struct {
	db PostgreSQL
}

func NewTxInContext(db *sql.DB) *txInContext {
	return &txInContext{
		db: db,
	}
}

func (t *txInContext) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error when creating transaction: %v", err)
	}

	ctx = newQueryCtx(ctx, tx)
	err = fn(ctx)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}

func newQueryCtx(ctx context.Context, q Queries) context.Context {
	return context.WithValue(ctx, key, q)
}

func GetQueryCtx(ctx context.Context) (Queries, bool) {
	q, ok := ctx.Value(key).(Queries)

	return q, ok
}
