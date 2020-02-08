package postgres

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type PDB struct {
	Con *sqlx.DB
}

// Connect _
func Connect(connStr string) (*PDB, error) {
	const pgDriver = "pgx"

	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	pdb := &PDB{Con: db}
	return pdb, nil
}

func (pdb *PDB) exec(query string, args ...interface{}) error {
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

func (pdb *PDB) namedExec(query string, arg interface{}) error {
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
