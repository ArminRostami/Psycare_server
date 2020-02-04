package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// import postgres driver package for side-effects
	_ "github.com/jackc/pgx/stdlib"
)

const pgDriver = "pgx"
const connStr = "postgres://postgres:example@localhost/postgres?sslmode=disable"

func GetDB(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	return db, nil
}
