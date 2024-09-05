package nap

import (
	"testing"
	"testing/quick"

	_ "github.com/kidbrains/nap/internal/mock/dummy"
)

func TestOpen(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Error(err)
	}

	if want, got := 3, len(db.pdbs); want != got {
		t.Errorf("Unexpected number of physical dbs. Got: %d, Want: %d", got, want)
	}
}

func TestPing(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Error(err)
	}
}

func TestBegin(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	if err = tx.Commit(); err != nil {
		t.Error(err)
	}
}

func TestExec(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if _, err = db.Exec("CREATE TABLE foo (id INTEGER);"); err != nil {
		t.Error(err)
	}
}

func TestPrepare(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT 1;")
	if err != nil {
		t.Error(err)
	}

	t.Run("exec", func(t *testing.T) {
		if _, err := stmt.Exec(); err != nil {
			t.Error(err)
		}
	})

	t.Run("query", func(t *testing.T) {
		if _, err := stmt.Query(); err != nil {
			t.Error(err)
		}
	})

	t.Run("query-row", func(t *testing.T) {
		if err := stmt.QueryRow().Scan(); err != nil {
			t.Error(err)
		}
	})

	t.Run("close", func(t *testing.T) {
		if err := stmt.Close(); err != nil {
			t.Error(err)
		}
	})
}

func TestQuery(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT 1;")
	if err != nil {
		t.Error(err)
	}

	if err = rows.Err(); err != nil {
		t.Error(err)
	}

	if err = rows.Close(); err != nil {
		t.Error(err)
	}
}

func TestQueryRow(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT 1;")
	if err = row.Err(); err != nil {
		t.Error(err)
	}

	if err = row.Scan(); err != nil {
		t.Error(err)
	}
}

func TestDriver(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if db.Driver() == nil {
		t.Error("Expected driver to be non-nil")
	}
}

func TestSet(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("IdleConns", func(t *testing.T) { db.SetMaxIdleConns(0) })
	t.Run("OpenConns", func(t *testing.T) { db.SetMaxOpenConns(0) })
	t.Run("Lifetime", func(t *testing.T) { db.SetConnMaxLifetime(0) })
}

func TestClose(t *testing.T) {
	db, err := Open("dummy", ":memory:;:memory:;:memory:")
	if err != nil {
		t.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err == nil || err.Error() != "sql: database is closed" {
		t.Errorf("Physical dbs were not closed correctly. Got: %s", err)
	}
}

func TestSlave(t *testing.T) {
	db := &DB{}
	last := -1

	err := quick.Check(
		func(n int) bool {
			index := db.slave(n)
			if n <= 1 {
				return index == 0
			}

			result := index > 0 && index < n && index != last
			last = index

			return result
		},
		nil,
	)

	if err != nil {
		t.Error(err)
	}
}
