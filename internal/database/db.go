package database

import (
	"context"
	"database/sql"
)

// DBTX defines what we need from a database
// Why? Because both *sql.DB and *sql.Tx support these operations
// This lets us use either direct DB calls or transactions
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New creates our database wrapper
// Why a wrapper? To hide SQL complexity and prevent SQL injection
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// Queries holds our database connection
// Why a struct? To keep database access organized and testable
type Queries struct {
	db DBTX
}

// WithTx lets us use transactions when needed
// Why? Some operations need to be atomic (all-or-nothing)
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
