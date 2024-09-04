package dummy

import (
	"database/sql/driver"
	"fmt"
)

// Driver of dummy db.
type Driver struct {
	isOpen bool
	Conn   *Conn
}

var (
	// Compile time validation that our types implement the expected interfaces
	_ driver.Driver = (*Driver)(nil)

	ErrDriverClosed = fmt.Errorf("%w database is closed", ErrDummy)
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
func (d *Driver) Open(dsn string) (driver.Conn, error) {
	d.isOpen = true
	d.Conn = &Conn{
		isOpen: true,
	}

	return d.Conn, nil
}

// Close database connection.
func (d *Driver) Close() error {
	if !d.isOpen {
		return ErrDriverClosed
	}

	d.isOpen = false

	return nil
}
