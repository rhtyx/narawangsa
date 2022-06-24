package postgres

import (
	"context"
	"database/sql"
)

type IQueries interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type TrxOpener interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type PostgreSQL interface {
	IQueries
	TrxOpener
}

func NewQueries(db IQueries) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db IQueries
}
