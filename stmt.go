package nap

import (
	"context"
	"database/sql"
)

// Stmt is an aggregate prepared statement.
// It holds a prepared statement for each underlying physical db.
type Stmt interface {
	Close() error
	Exec(...interface{}) (sql.Result, error)
	ExecContext(context.Context, ...interface{}) (sql.Result, error)
	Query(...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, ...interface{}) (*sql.Rows, error)
	QueryRow(...interface{}) *sql.Row
	QueryRowContext(context.Context, ...interface{}) *sql.Row
}

type stmt struct {
	db    *DB
	stmts []*sql.Stmt
}

// Close closes the statement by concurrently closing all underlying
// statements concurrently, returning the first non nil error.
func (s *stmt) Close() error {
	return scatter(len(s.stmts), func(i int) error {
		return s.stmts[i].Close()
	})
}

// Exec executes a prepared statement with the given arguments
// and returns a Result summarizing the effect of the statement.
// Exec uses the master as the underlying physical db.
func (s *stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.ExecContext(context.Background(), args...)
}

// ExecContext executes a prepared statement with the given arguments.
func (s *stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.stmts[0].ExecContext(ctx, args...)
}

// Query executes a prepared query statement with the given
// arguments and returns the query results as a *sql.Rows.
// Query uses a slave as the underlying physical db.
func (s *stmt) Query(args ...interface{}) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), args...)
}

func (s *stmt) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	return s.stmts[s.db.slave(len(s.db.pdbs))].QueryContext(ctx, args...)
}

// QueryRow executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error
// will be returned by a call to Scan on the returned *Row, which is always non-nil.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *sql.Row's Scan scans the first selected row and discards the rest.
// QueryRow uses a slave as the underlying physical db.
func (s *stmt) QueryRow(args ...interface{}) *sql.Row {
	return s.QueryRowContext(context.Background(), args...)
}

// QueryRowContext executes a prepared query statement with the given arguments.
func (s *stmt) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	return s.stmts[s.db.slave(len(s.db.pdbs))].QueryRowContext(ctx, args...)
}
