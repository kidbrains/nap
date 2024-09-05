package dummy

import (
	"database/sql/driver"
	"fmt"
)

// Tx implements transaction.
type Tx struct {
	Closed bool
}

var (
	_ driver.Tx = (*Tx)(nil)

	ErrTxDone = fmt.Errorf("%w transaction is done", ErrDummy)
)

// Begin implements the DB interface.
// Commit implements the Tx interface.
func (tx *Tx) Commit() error {
	if tx.Closed {
		return ErrTxDone
	}

	tx.Closed = true

	return nil
}

// Rollback implements the Tx interface.
func (tx *Tx) Rollback() error {
	if tx.Closed {
		return ErrTxDone
	}

	tx.Closed = true

	return nil
}
