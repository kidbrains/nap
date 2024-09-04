package dummy

import (
	"database/sql/driver"
	"fmt"
)

// Tx implements transaction.
type Tx struct {
	isOpen bool
}

var (
	_ driver.Tx = (*Tx)(nil)

	ErrTxDone = fmt.Errorf("%w transaction is done", ErrDummy)
)

// Begin implements the DB interface.
// Commit implements the Tx interface.
func (tx *Tx) Commit() error {
	if !tx.isOpen {
		return ErrTxDone
	}

	tx.isOpen = false

	return nil
}

// Rollback implements the Tx interface.
func (tx *Tx) Rollback() error {
	if !tx.isOpen {
		return ErrTxDone
	}

	tx.isOpen = false

	return nil
}
