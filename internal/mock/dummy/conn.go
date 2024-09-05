package dummy

import (
	"context"
	"database/sql/driver"
	"fmt"
)

// Conn - connection to dummydb.
type Conn struct {
	Closed bool
	Result *Result
	Rows   *Rows
	Stmt   *Stmt
	Tx     *Tx
}

var (
	// Compile time validation that our types implement the expected interfaces
	_ driver.Conn = &Conn{}

	ErrConnClosed = fmt.Errorf("%w connection is closed", ErrDummy)
)

// Prepare returns a prepared statement, bound to this connection.
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

// PrepareContext returns a prepared statement, bound to this connection.
// context is for the preparation of the statement,
// it must not store the context within the statement itself.
func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	c.Stmt = &Stmt{}

	return c.Stmt, nil
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use.
//
// Because the sql package maintains a free pool of
// connections and only calls Close when there's a surplus of
// idle connections, it shouldn't be necessary for drivers to
// do their own connection caching.
//
// Drivers must ensure all network calls made by Close
// do not block indefinitely (e.g. apply a timeout).
func (c *Conn) Close() error {
	if c.Closed {
		return ErrConnClosed
	}

	c.Closed = true

	return nil
}

// Begin starts and returns a new transaction.
//
// Deprecated: Drivers should implement BeginTx instead (or additionally).
func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

// BeginTx starts and returns a new transaction.
// If the context is canceled by the user the sql package will
// call Tx.Rollback before discarding and closing the connection.
//
// This must check opts.Isolation to determine if there is a set
// isolation level. If the driver does not support a non-default
// level and one is set or if there is a non-default isolation level
// that is not supported, an error must be returned.
//
// This must also check opts.ReadOnly to determine if the read-only
// value is true to either set the read-only transaction property if supported
// or return an error if it is not supported.
func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	c.Tx = &Tx{}

	return c.Tx, nil
}

// Ping - check database connection alive.
func (c *Conn) Ping(ctx context.Context) error {
	if c.Closed {
		return ErrConnClosed
	}

	return nil
}

// ExecContext executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	c.Result = &Result{}

	return c.Result, ctx.Err()
}

// NamedValue convert []driver.Value into []driver.NamedValue.
func (c *Conn) NamedValue(args []driver.Value) []driver.NamedValue {
	named := make([]driver.NamedValue, len(args))
	for i, arg := range args {
		named[i] = driver.NamedValue{Value: arg}
	}

	return named
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
func (c *Conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return c.ExecContext(context.Background(), query, c.NamedValue(args))
}

// QueryContext executes a query that may return rows, such as a SELECT.
func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	c.Rows = &Rows{}

	return c.Rows, ctx.Err()
}

// Query executes a query that may return rows, such as a SELECT.
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return c.QueryContext(context.Background(), query, c.NamedValue(args))
}
