package dummy

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
)

func TestDriver(t *testing.T) {
	t.Run("sql", func(t *testing.T) {
		db, err := sql.Open("dummy", "dsn")
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
	})

	t.Run("open", func(t *testing.T) {
		d := Driver{}

		c, err := d.Open("dsn")
		if err != nil {
			t.Fatal(err)
		}

		if c == nil {
			t.Fatal("conn is nil")
		}
	})

	t.Run("invalid-dsn", func(t *testing.T) {
		d := Driver{}

		if _, err := d.Open(""); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestConnector(t *testing.T) {
	t.Run("connect", func(t *testing.T) {
		c := &Connector{}

		conn, err := c.Connect(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if conn == nil {
			t.Fatal("conn is nil")
		}
	})

	t.Run("driver", func(t *testing.T) {
		c := &Connector{}

		if c.Driver() == nil {
			t.Fatal("driver is nil")
		}
	})
}

func TestConn(t *testing.T) {
	t.Run("ping", func(t *testing.T) {
		ctx := context.Background()
		c := &Conn{}

		if err := c.Ping(ctx); err != nil {
			t.Fatal(err)
		}

		if err := c.Close(); err != nil {
			t.Fatal(err)
		}

		if err := c.Ping(ctx); err == nil {
			t.Fail()
		}
	})
	t.Run("close", func(t *testing.T) {
		c := &Conn{}

		if err := c.Close(); err != nil {
			t.Fatal(err)
		}

		if err := c.Close(); err == nil {
			t.Fail()
		}
	})

	t.Run("prepare", func(t *testing.T) {
		c := &Conn{}

		s, err := c.Prepare("SELECT *")

		if err != nil {
			t.Fatal(err)
		}

		if s == nil {
			t.Fail()
		}
	})

	t.Run("query", func(t *testing.T) {
		c := &Conn{}
		r, err := c.Query("SELECT * FROM table", nil)

		if err != nil {
			t.Fatal(err)
		}

		if r == nil {
			t.Fail()
		}
	})

	t.Run("exec", func(t *testing.T) {
		c := &Conn{}
		r, err := c.Exec("INSERT INTO table (?, ?)", []driver.Value{1, "name"})

		if err != nil {
			t.Fatal(err)
		}

		if r == nil {
			t.Fail()
		}
	})

	t.Run("begin", func(t *testing.T) {
		c := &Conn{}
		tx, err := c.Begin()

		if err != nil {
			t.Fatal(err)
		}

		if tx == nil {
			t.Fail()
		}
	})
}

func TestResult(t *testing.T) {
	t.Run("insert_id", func(t *testing.T) {
		r := &Result{}

		_, err := r.LastInsertId()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("affected", func(t *testing.T) {
		r := &Result{}
		_, err := r.RowsAffected()

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestRows(t *testing.T) {
	t.Run("close", func(t *testing.T) {
		r := &Rows{}

		if err := r.Close(); err != nil {
			t.Fatal(err)
		}

		if err := r.Close(); err == nil {
			t.Fail()
		}
	})

	t.Run("next", func(t *testing.T) {
		r := &Rows{}

		if err := r.Next(nil); err != nil {
			t.Fatal(err)
		}

		if err := r.Close(); err != nil {
			t.Fatal(err)
		}

		if err := r.Next(nil); err == nil {
			t.Fail()
		}
	})
	t.Run("columns", func(t *testing.T) {
		r := &Rows{}

		if r.Columns() != nil {
			t.Fail()
		}
	})
}

func TestStmt(t *testing.T) {
	t.Run("close", func(t *testing.T) {
		s := &Stmt{}

		if err := s.Close(); err != nil {
			t.Fatal(err)
		}

		if err := s.Close(); err == nil {
			t.Fail()
		}
	})

	t.Run("exec", func(t *testing.T) {
		s := &Stmt{}

		if _, err := s.Exec(nil); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("query", func(t *testing.T) {
		s := &Stmt{}

		if _, err := s.Query(nil); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("num_input", func(t *testing.T) {
		s := &Stmt{}

		if n := s.NumInput(); n != -1 {
			t.Fatal(n)
		}
	})
}

func TestTx(t *testing.T) {
	t.Run("commit", func(t *testing.T) {
		tx := &Tx{}

		if err := tx.Commit(); err != nil {
			t.Fatal(err)
		}

		if err := tx.Commit(); err == nil {
			t.Fail()
		}
	})

	t.Run("rollback", func(t *testing.T) {
		tx := &Tx{}

		if err := tx.Rollback(); err != nil {
			t.Fatal(err)
		}

		if err := tx.Rollback(); err == nil {
			t.Fail()
		}
	})
}
