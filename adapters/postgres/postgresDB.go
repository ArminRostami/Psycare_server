package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PDB struct {
	Con *sqlx.DB
}

func Connect(connStr string) (*PDB, error) {
	const pgDriver = "pgx"

	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "db connection error")
	}
	return &PDB{Con: db}, nil
}

func (pdb *PDB) Execute(query string, args ...interface{}) error {
	tx, err := pdb.Con.Beginx()
	if err != nil {
		return errors.Wrap(err, "could not initialize db transaction")
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "could not execute query")

	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "could not commit db transaction")
	}
	return nil
}

func (pdb *PDB) NamedExecute(query string, arg interface{}) error {
	tx, err := pdb.Con.Beginx()
	if err != nil {
		return errors.Wrap(err, "could not initialize db transaction")
	}
	_, err = tx.NamedExec(query, arg)
	if err != nil {
		return errors.Wrap(err, "could not execute query")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "could not commit db transaction")

	}
	return nil
}
