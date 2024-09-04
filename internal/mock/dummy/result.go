package dummy

import "database/sql/driver"

type Result struct{}

var (
	// Compile time validation that our types implement the expected interfaces
	_ driver.Result = (*Result)(nil)
)

// LastInsertId returns the database's internal ID of the last inserted row.
// This is typically not useful to applications.
// A dummy implementation is returned.
func (r *Result) LastInsertId() (int64, error) { return 0, nil }

// RowsAffected returns the number of rows affected by the query.
// A dummy implementation is returned.
func (r *Result) RowsAffected() (int64, error) { return 0, nil }
