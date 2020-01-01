package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// import postgres driver package for side-effects
	_ "github.com/jackc/pgx/stdlib"
)

const pgDriver = "pgx"
const connStr = "postgres://postgres:example@localhost/postgres?sslmode=disable"

func getDB(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	_, err = db.Exec(DefaultSchema.create)
	if err != nil {
		return nil, fmt.Errorf("schema failed to create: %w", err)
	}
	return db, nil
}
