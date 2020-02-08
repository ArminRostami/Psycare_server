package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// import postgres driver package for side-effects
	_ "github.com/jackc/pgx/stdlib"
)

type DB struct {
	Con *sqlx.DB
}

// Connect _
func Connect(connStr string) (*DB, error) {
	const pgDriver = "pgx"

	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	pdb := &DB{Con: db}
	return pdb, nil
}

func (pdb *DB) exec(query string, args ...interface{}) error {
	tx, err := pdb.Con.Beginx()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not initialize db transaction")
	}
	_, err = tx.Exec(query, args)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not execute query")
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not commit db transaction")
	}
	return nil
}

func (pdb *DB) namedExec(query string, arg interface{}) error {
	tx, err := pdb.Con.Beginx()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not initialize db transaction")
	}
	_, err = tx.NamedExec(query, arg)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not execute query")
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not commit db transaction")
	}
	return nil
}
