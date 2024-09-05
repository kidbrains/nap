package dummy

import (
	"context"
	"database/sql/driver"
	"fmt"
)

// Driver of dummy db.
type Driver struct{}

var (
	// Compile time validation that our types implement the expected interfaces
	_ driver.Driver = Driver{}

	ErrInvalidDSN = fmt.Errorf("%w invalid dsn", ErrDummy)
)

// Open returns a new connection to the database.
// The name is a string in a driver-specific format.
//
// Open may return a cached connection (one previously
// closed), but doing so is unnecessary; the sql package
// maintains a pool of idle connections for efficient re-use.
//
// The returned connection is only used by one goroutine at a
// time.
func (d Driver) Open(dsn string) (driver.Conn, error) {
	c, err := d.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}

	return c.Connect(context.Background())
}

// OpenConnector must parse the name in the same format that Driver.Open
// parses the name parameter.
func (d Driver) OpenConnector(dsn string) (driver.Connector, error) {
	if dsn == "" {
		return nil, ErrInvalidDSN
	}

	c := &Connector{
		DSN: dsn,
	}

	return c, nil
}
