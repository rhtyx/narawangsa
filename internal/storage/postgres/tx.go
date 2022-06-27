package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

var key int = 0

type TxInContext struct {
	db PostgreSQL
}

func NewTxInContext(db *sql.DB) *TxInContext {
	return &TxInContext{
		db: db,
	}
}

func (t *TxInContext) Run(ctx context.Context, fn func(ctx context.Context) error) error {
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

func newQueryCtx(ctx context.Context, q IQueries) context.Context {
	return context.WithValue(ctx, key, q)
}

func GetQueryCtx(ctx context.Context) (IQueries, bool) {
	q, ok := ctx.Value(key).(IQueries)

	return q, ok
}
