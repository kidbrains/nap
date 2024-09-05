package dummy

import (
	"context"
	"database/sql/driver"
	"fmt"
)

type Stmt struct {
	Closed bool
	Result *Result
	Rows   *Rows
}

var (
	// Compile time validation that our types implement the expected interfaces
	_ driver.Stmt = (*Stmt)(nil)

	ErrStmtClosed = fmt.Errorf("%w statement is closed", ErrDummy)
)

// Close closes the statement.
//
// Drivers must ensure all network calls made by Close
// do not block indefinitely (e.g. apply a timeout).
func (s *Stmt) Close() error {
	if s.Closed {
		return ErrStmtClosed
	}

	s.Closed = true

	return nil
}

// NumInput returns the number of placeholder parameters.
//
// If NumInput returns >= 0, the sql package will sanity check
// argument counts from callers and return errors to the caller
// before the statement's Exec or Query methods are called.
//
// NumInput may also return -1, if the driver doesn't know
// its number of placeholders. In that case, the sql package
// will not sanity check Exec or Query argument counts.
func (s *Stmt) NumInput() int { return -1 }

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
//
// Deprecated: Drivers should implement ExecContext instead (or additionally).
func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), args)
}

// ExecContext executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
func (s *Stmt) ExecContext(ctx context.Context, args []driver.Value) (driver.Result, error) {
	s.Result = &Result{}

	return s.Result, nil
}

// Query executes a query that may return rows, such as a SELECT.
//
// Deprecated: Drivers should implement QueryContext instead (or additionally).
func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), args)
}

// QueryContext executes a query that may return rows, such as a SELECT.
func (s *Stmt) QueryContext(ctx context.Context, args []driver.Value) (driver.Rows, error) {
	s.Rows = &Rows{}

	return s.Rows, nil
}
