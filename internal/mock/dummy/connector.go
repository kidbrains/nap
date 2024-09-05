package dummy

import (
	"context"
	"database/sql/driver"
)

type Connector struct {
	DSN  string
	Conn *Conn
}

var _ driver.Connector = (*Connector)(nil)

// Connect returns a connection to the database.
// Connect may return a cached connection (one previously
// closed), but doing so is unnecessary; the sql package
// maintains a pool of idle connections for efficient re-use.
//
// The provided context.Context is for dialing purposes only
// (see net.DialContext) and should not be stored or used for
// other purposes. A default timeout should still be used
// when dialing as a connection pool may call Connect
// asynchronously to any query.
//
// The returned connection is only used by one goroutine at a
// time.
func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	c.Conn = &Conn{}

	return c.Conn, ctx.Err()
}

// Driver returns the underlying Driver of the Connector,
// mainly to maintain compatibility with the Driver method
// on sql.DB.
func (v *Connector) Driver() driver.Driver { return Driver{} }
