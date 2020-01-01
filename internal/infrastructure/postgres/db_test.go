package postgres

import (
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestSchema(t *testing.T) {
	connStr := "postgres://postgres:example@localhost/postgres?sslmode=disable"
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		t.Error("\nDB connection failed:\n", err.Error())
	}
	_, err = db.Exec(DefaultSchema.create)
	if err != nil {
		t.Error("\nDB schema error:\n", err.Error())
	}
}
func TestSchemaDrop(t *testing.T) {
	connStr := "postgres://postgres:example@localhost/postgres?sslmode=disable"
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		t.Error("\nDB connection failed:\n", err.Error())
	}
	_, err = db.Exec(DefaultSchema.drop)
	if err != nil {
		t.Error("\nDB drop error:\n", err.Error())
	}
}
