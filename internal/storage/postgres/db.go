package postgres

import (
	"context"
	"database/sql"
)

type Queries interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type TrxOpener interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type PostgreSQL interface {
	Queries
	TrxOpener
}
