package dummy

import (
	"database/sql"
	"errors"
)

var ErrDummy = errors.New("dummy")

func init() { sql.Register("dummy", &Driver{}) }
